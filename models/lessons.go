package models

import (
	// "backend/models"
	"backend/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Lesson main struct for
type Lesson struct {
	gorm.Model
	Pid        int    `json:"pid"`
	WeekNumber int    `json:"weekNumber"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	Order      int    `json:"order"`
	ImageURL   string `json:"imageURL"`
	AudioURL   string `json:"audioURL"`
	PdfURL     string `json:"pdfURL"`
}

// Validate method
func (lesson *Lesson) Validate() (map[string]interface{}, bool) {

	// if lesson.Alias == "" {
	// 	return utils.Message(false, "Lesson name should be on the payload"), false
	// }
	return utils.Message(true, "success"), true
}

// Create method
func (lesson *Lesson) Create() map[string]interface{} {
	if resp, ok := lesson.Validate(); !ok {
		return resp
	}
	GetDB().Create(lesson)
	resp := utils.Message(true, "success")
	resp["lesson"] = lesson
	return resp
}

// GetLesson method
func (lesson *Lesson) GetLesson(lessonID int) *Lesson {
	// acc := &Lesson{}
	err := GetDB().Table("lessons").Where("id = ?", lessonID).First(lesson).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return lesson
}

// GetLessons method
func (lesson *Lesson) GetLessons() []*Lesson {
	lessons := make([]*Lesson, 0)
	err := GetDB().Find(&lessons).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return lessons
}

// DeleteLesson method
func DeleteLesson(id int) *Lesson {
	lesson := new(Lesson)
	err := GetDB().Table("lessons").Where("id = ?", id).Delete(&lesson).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return lesson
}

// Update method
func (lesson *Lesson) Update() map[string]interface{} {
	err := GetDB().Omit("created_at", "deleted_at").Model(&lesson).Updates(&lesson).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resp := utils.Message(true, "success")
	resp["lesson"] = lesson.GetLesson(int(lesson.Model.ID))
	return resp
}
