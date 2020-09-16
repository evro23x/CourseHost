package models

import (
	"backend/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

// Field struct
type Field struct {
	gorm.Model
	PageID      int          `json:"page_id" gorm:"unique_index:idx_pageid_key"`
	Key         string       `json:"key" gorm:"unique_index:idx_pageid_key"`
	Type        string       `json:"type"`
	FieldValues []FieldValue `json:"fieldValues" gorm:"foreignkey:FieldID"`
}

// Validate method
func (field *Field) Validate() (map[string]interface{}, bool) {
	if field.Key == "" || field.Type == "" {
		return utils.Message(false, "Fields key or type should be on the payload"), false
	}
	return utils.Message(true, "success"), true
}

// Create method
func (field *Field) Create() map[string]interface{} {
	if resp, ok := field.Validate(); !ok {
		return resp
	}

	d := GetDB().Create(field)
	if d.Error != nil {
		return utils.Message(false, "Cant add new field")
	}

	resp := utils.Message(true, "success")
	resp["field"] = field
	return resp
}

// GetIDByKey method
func (field *Field) GetIDByKey() *Field {
	err := GetDB().Where("key = ?", field.Key).First(&field).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return field
}

// Update method
func (field *Field) Update(Key string, ValueType string) map[string]interface{} {
	field.Key = Key
	field.Type = ValueType
	d := GetDB().Save(&field)
	if d.Error != nil {
		return utils.Message(false, "Cant update key or type for field")
	}
	resp := utils.Message(true, "success")
	resp["fieldValue"] = field
	return resp
}

// DeleteField method
func DeleteField(key string) *Field {
	fieldValue := new(FieldValue)
	err := GetDB().Table("fields").
		Select("*").
		Joins("left join field_values on fields.id = field_values.field_id").
		Where("fields.key = ?", key).Scan(&fieldValue).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = GetDB().Delete(&fieldValue).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	field := new(Field)
	err = GetDB().Table("fields").Where("key = ?", key).Delete(&field).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return field
}
