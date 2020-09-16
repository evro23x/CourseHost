package controllers

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// CreateUser method
var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := user.Create() //Create User
	utils.Respond(w, resp)
	fmt.Println("%V - /user/add", time.Now())
}

// EditUser method
var EditUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	bearerToken := strings.Split(r.Header.Get("Authorization"), " ")
	resp := user.Update(bearerToken[1])
	utils.Respond(w, resp)
	fmt.Println("%V - /user/edit", time.Now())
}

// GetAllUser method
var GetAllUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	data := user.GetAllUserInfo()
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
	fmt.Println("%V - /user/get_all", time.Now())
}

// GetUser method
var GetUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	bearerToken := strings.Split(r.Header.Get("Authorization"), " ")
	resp := user.GetUserInfo(bearerToken[1])
	utils.Respond(w, resp)
	fmt.Println("%V - /user/edit", time.Now())
}

// Authenticate method
var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	User := &models.User{}
	err := json.NewDecoder(r.Body).Decode(User) //decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(User.Email, User.Password)
	utils.Respond(w, resp)
	fmt.Println("%V - Authenticate", time.Now())
}
