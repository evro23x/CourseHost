package controllers

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GetByIDReq struct
type GetByIDReq struct {
	ID int
}

// AddCourse method
var AddCourse = func(w http.ResponseWriter, r *http.Request) {
	course := &models.Course{}
	err := json.NewDecoder(r.Body).Decode(course)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	resp := course.Create()
	utils.Respond(w, resp)
	fmt.Println("%V - /course/add", time.Now())
}

// GetCourse method
var GetCourse = func(w http.ResponseWriter, r *http.Request) {
	var request GetByIDReq
	course := &models.Course{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	data := course.GetCourseByID(request.ID)
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
	fmt.Println("%V - /course/get", time.Now())
}

// GetAllCourses method
var GetAllCourses = func(w http.ResponseWriter, r *http.Request) {
	course := &models.Course{}
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	data := course.GetCourses()
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
	fmt.Println("%V - /course/get_all", time.Now())
}

// EditCourse method
var EditCourse = func(w http.ResponseWriter, r *http.Request) {
	course := &models.Course{}
	err := json.NewDecoder(r.Body).Decode(course)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	resp := course.Update()
	utils.Respond(w, resp)
	fmt.Println("%V - /course/add", time.Now())
}

// RemoveCourse method
var RemoveCourse = func(w http.ResponseWriter, r *http.Request) {
	var request GetByIDReq
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	models.DeleteCourse(request.ID)
	resp := utils.Message(true, "success")
	utils.Respond(w, resp)
	fmt.Println("%V - /course/remove", time.Now())
}
