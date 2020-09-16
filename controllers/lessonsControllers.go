package controllers

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// AddLesson method
var AddLesson = func(w http.ResponseWriter, r *http.Request) {
	lesson := &models.Lesson{}
	err := json.NewDecoder(r.Body).Decode(lesson)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	resp := lesson.Create()
	utils.Respond(w, resp)
	fmt.Println("%V - /lesson/add", time.Now())
}

// GetLesson method
var GetLesson = func(w http.ResponseWriter, r *http.Request) {
	var request GetByIDReq
	lesson := &models.Lesson{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	data := lesson.GetLesson(request.ID)
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
	fmt.Println("%V - /lesson/get", time.Now())
}

// GetAllLessons method
var GetAllLessons = func(w http.ResponseWriter, r *http.Request) {
	course := &models.Lesson{}
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	data := course.GetLessons()
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
	fmt.Println("%V - /lesson/get_all", time.Now())
}

// EditLesson method
var EditLesson = func(w http.ResponseWriter, r *http.Request) {
	lesson := &models.Lesson{}
	err := json.NewDecoder(r.Body).Decode(lesson)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	resp := lesson.Update()
	utils.Respond(w, resp)
	fmt.Println("%V - /lesson/add", time.Now())
}

// RemoveLesson method
var RemoveLesson = func(w http.ResponseWriter, r *http.Request) {
	var request GetByIDReq
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	models.DeleteLesson(request.ID)
	resp := utils.Message(true, "success")
	utils.Respond(w, resp)
	fmt.Println("%V - /lesson/remove", time.Now())
}
