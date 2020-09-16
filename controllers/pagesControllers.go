package controllers

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type getPageRequest struct {
	Id int
}

type EditPageRequest struct {
	PageID int    `json:"id"`
	PId    string `json:"pid"`
	Alias  string `json:"alias"`
}

var AddPageInfo = func(w http.ResponseWriter, r *http.Request) {
	page := &models.Page{}
	err := json.NewDecoder(r.Body).Decode(page)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	resp := page.Create()
	utils.Respond(w, resp)
	fmt.Println("%V - /page/add", time.Now())
}

var EditPageInfo = func(w http.ResponseWriter, r *http.Request) {
	page := &models.Page{}
	err := json.NewDecoder(r.Body).Decode(&page)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	resp := page.Update()
	//resp := utils.Message(true, "success")
	//resp["result"] = data
	utils.Respond(w, resp)
	fmt.Println("%V - /page/edit", time.Now())
}

var GetPageById = func(w http.ResponseWriter, r *http.Request) {
	var request getPageRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	data := models.GetPage(request.Id)
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
	fmt.Println("%V - /page/get", time.Now())
}

var GetPageList = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllPages()
	resp := utils.Message(true, "success")
	resp["pages"] = data
	utils.Respond(w, resp)
	fmt.Println("%V - /page/list", time.Now())
}

var RemovePage = func(w http.ResponseWriter, r *http.Request) {
	var request getPageRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	models.DeletePage(request.Id)
	resp := utils.Message(true, "success")
	utils.Respond(w, resp)
	fmt.Println("%V - /page/remove", time.Now())
}
