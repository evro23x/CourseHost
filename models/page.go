package models

import (
	// "backend/models"
	"backend/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Page main struct for
type Page struct {
	gorm.Model
	Pid         string  `json:"pid"`
	Alias       string  `json:"alias"`
	Title       string  `json:"title"`
	Name        string  `json:"name"`
	Keywords    string  `json:"keywords"`
	Description string  `json:"description"`
	ShortText   string  `json:"shortText"`
	FullText    string  `json:"fullText"`
	Fields      []Field `json:"fields" gorm:"foreignkey:PageID"`
}

type getPageResult struct {
	ID     string                `json:"ID" gorm:"id"`
	Fields []getPageFieldsResult `json:"fields"`
}

type getPageFieldsResult struct {
	Key       string `json:"key"`
	KeyType   string `json:"key_type"`
	Value     string `json:"value"`
	ValueType string `json:"value_type"`
}

type getAllPagesResult struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"CreatedAt"`
	UpdatedAt   string `json:"UpdatedAt"`
	Pid         string `json:"pid"`
	Alias       string `json:"alias"`
	Title       string `json:"title"`
	Name        string `json:"name"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	ShortText   string `json:"shortText"`
	FullText    string `json:"fullText"`
}

// Validate method
func (page *Page) Validate() (map[string]interface{}, bool) {

	if page.Alias == "" {
		return utils.Message(false, "Page name should be on the payload"), false
	}
	return utils.Message(true, "success"), true
}

// Create method
func (page *Page) Create() map[string]interface{} {
	if resp, ok := page.Validate(); !ok {
		return resp
	}
	GetDB().Create(page)
	resp := utils.Message(true, "success")
	resp["page"] = page
	return resp
}

// GetAllPages method
func GetAllPages() []*getAllPagesResult {
	getAllPagesResult := make([]*getAllPagesResult, 0)
	err := GetDB().
		Table("pages").
		Select("pages.id, pages.created_at, pages.updated_at, pages.pid, pages.alias, pages.title, " +
			"pages.name, pages.keywords, pages.description, pages.short_text, pages.full_text").
		Where("pages.deleted_at is null").
		Scan(&getAllPagesResult).Error
	if err != nil {
		return nil
	}
	return getAllPagesResult
}

// GetPage method
func GetPage(id int) *getPageResult {
	fields := make([]getPageFieldsResult, 0)

	err := GetDB().
		Table("pages").
		Select([]string{"fields.key as key, fields.type as key_type, field_values.value as value, field_values.type as value_type"}).
		Joins("left join fields on fields.page_id = pages.id").
		Joins("left join field_values on field_values.field_id = fields.id").
		Where("pages.id = ? and fields.deleted_at is null", id).
		Scan(&fields).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	page := new(getPageResult)
	err = GetDB().Table("pages").Where("id = ?", id).Find(&page).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	page.Fields = fields
	return page
}

// DeletePage method
func DeletePage(id int) *Page {
	page := new(Page)
	err := GetDB().Table("pages").Where("id = ?", id).Delete(&page).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return page
}

// Update method
func (page *Page) Update() map[string]interface{} {
	//if resp, ok := page.Validate(); !ok {
	//	return resp
	//}
	GetDB().Save(page)
	resp := utils.Message(true, "success")
	resp["page"] = page
	return resp
}
