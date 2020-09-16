package models

import (
	// "backend/models"
	"backend/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Course main struct for
type Course struct {
	gorm.Model
	UID         int    `json:"uid"`
	Type        int    `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Smalltext   string `json:"smalltext"`
	Text        string `json:"text"`
	WeekCount   int    `json:"weekCount"`
	Price       int    `json:"price"`
}

// Validate method
func (course *Course) Validate() (map[string]interface{}, bool) {

	// if course.Alias == "" {
	// 	return utils.Message(false, "Course name should be on the payload"), false
	// }
	return utils.Message(true, "success"), true
}

// Create method
func (course *Course) Create() map[string]interface{} {
	if resp, ok := course.Validate(); !ok {
		return resp
	}
	GetDB().Create(course)
	resp := utils.Message(true, "success")
	resp["course"] = course
	return resp
}

// GetCourseByID method
func (course *Course) GetCourseByID(courseID int) *Course {
	temp := &Course{}
	err := GetDB().Table("courses").Where("id = ?", courseID).First(temp).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return temp
}

// GetCourses method
func (course *Course) GetCourses() []*Course {
	courses := make([]*Course, 0)
	err := GetDB().Find(&courses).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return courses
}

// DeleteCourse method
func DeleteCourse(id int) *Course {
	course := new(Course)
	err := GetDB().Table("courses").Where("id = ?", id).Delete(&course).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return course
}

// Update method
func (course *Course) Update() map[string]interface{} {
	//if resp, ok := course.Validate(); !ok {
	//	return resp
	//}
	err := GetDB().Omit("created_at", "deleted_at").Model(&course).Updates(&course).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resp := utils.Message(true, "success")
	resp["course"] = course.GetCourseByID(int(course.Model.ID))
	return resp
}
