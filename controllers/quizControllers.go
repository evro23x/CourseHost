package controllers

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// AddQuiz method
var AddQuiz = func(w http.ResponseWriter, r *http.Request) {
	quiz := &models.Quiz{}
	err := json.NewDecoder(r.Body).Decode(quiz)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	resp := quiz.Create()
	utils.Respond(w, resp)
	fmt.Println("%V - /quiz/add", time.Now())
}

// GetQuiz method
var GetQuiz = func(w http.ResponseWriter, r *http.Request) {
	var request GetByIDReq
	quiz := &models.Quiz{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	data := quiz.GetQuiz(request.ID)
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
	fmt.Println("%V - /quiz/get", time.Now())
}

// GetAllQuiz method
var GetAllQuiz = func(w http.ResponseWriter, r *http.Request) {
	quiz := &models.Quiz{}
	err := json.NewDecoder(r.Body).Decode(&quiz)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	data := quiz.GetAllQuiz()
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
	fmt.Println("%V - /quiz/get_all", time.Now())
}

// EditQuiz method
var EditQuiz = func(w http.ResponseWriter, r *http.Request) {
	quiz := &models.Quiz{}
	err := json.NewDecoder(r.Body).Decode(quiz)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	resp := quiz.Update()
	utils.Respond(w, resp)
	fmt.Println("%V - /quiz/add", time.Now())
}

// RemoveQuiz method
var RemoveQuiz = func(w http.ResponseWriter, r *http.Request) {
	var request GetByIDReq
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	models.DeleteQuiz(request.ID)
	resp := utils.Message(true, "success")
	utils.Respond(w, resp)
	fmt.Println("%V - /quiz/remove", time.Now())
}
