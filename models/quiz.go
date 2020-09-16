package models

import (
	// "backend/models"
	"backend/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Quiz main struct for
type Quiz struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	CourseID    int    `json:"courseID"`
	ImageURL    string `json:"imageURL"`
	VideoURL    string `json:"videoURL"`
	Main        int    `json:"main"`
}

// Validate method
func (quiz *Quiz) Validate() (map[string]interface{}, bool) {

	// if quiz.Alias == "" {
	// 	return utils.Message(false, "Quiz name should be on the payload"), false
	// }
	return utils.Message(true, "success"), true
}

// Create method
func (quiz *Quiz) Create() map[string]interface{} {
	if resp, ok := quiz.Validate(); !ok {
		return resp
	}
	GetDB().Create(quiz)
	resp := utils.Message(true, "success")
	resp["quiz"] = quiz
	return resp
}

// GetQuiz method
func (quiz *Quiz) GetQuiz(quizID int) *Quiz {
	temp := &Quiz{}
	err := GetDB().Table("quizzes").Where("id = ?", quizID).First(temp).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return temp
}

// GetAllQuiz method
func (quiz *Quiz) GetAllQuiz() []*Quiz {
	quizs := make([]*Quiz, 0)
	err := GetDB().Find(&quizs).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return quizs
}

// DeleteQuiz method
func DeleteQuiz(id int) *Quiz {
	quiz := new(Quiz)
	err := GetDB().Table("quizzes").Where("id = ?", id).Delete(&quiz).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return quiz
}

// Update method
func (quiz *Quiz) Update() map[string]interface{} {
	err := GetDB().Omit("created_at", "deleted_at").Model(&quiz).Updates(&quiz).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resp := utils.Message(true, "success")
	resp["quiz"] = quiz.GetQuiz(int(quiz.Model.ID))
	return resp
}
