package controllers

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"net/http"
)

type AddFieldParams struct {
	PageID         int    `json:"page_id"`
	FieldKey       string `json:"key"`
	FieldKeyType   string `json:"key_type"`
	FieldValue     string `json:"value"`
	FieldValueType string `json:"value_type"`
}

type FieldByKeyRequest struct {
	Key string
}

type UpdateFieldRequest struct {
	FieldKey       string `json:"key"`
	FieldKeyType   string `json:"key_type"`
	FieldValue     string `json:"value"`
	FieldValueType string `json:"value_type"`
}

var AddField = func(w http.ResponseWriter, r *http.Request) {
	var query AddFieldParams
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	field := &models.Field{PageID: query.PageID, Key: query.FieldKey, Type: query.FieldKeyType}

	field.Create()
	fieldValue := &models.FieldValue{FieldID: field.ID, Type: query.FieldValueType, Value: query.FieldValue}
	resp := fieldValue.Create()
	utils.Respond(w, resp)
}

var EditField = func(w http.ResponseWriter, r *http.Request) {
	var query UpdateFieldRequest
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	field := &models.Field{Key: query.FieldKey, Type: query.FieldKeyType}
	field.GetIDByKey()
	fieldValue := &models.FieldValue{FieldID: field.ID}
	fieldValue.GetValueIDByKeyID()
	fieldValue.Update(query.FieldValue, query.FieldValueType)
	field.Update(query.FieldKey, query.FieldKeyType)
	resp := utils.Message(true, "success")
	utils.Respond(w, resp)
}

var RemoveField = func(w http.ResponseWriter, r *http.Request) {
	var request FieldByKeyRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	models.DeleteField(request.Key)
	resp := utils.Message(true, "success")
	utils.Respond(w, resp)
}

//
//func GetIp(r *http.Request) *[]byte {
//	var userIP []byte
//	ip, _, err := net.SplitHostPort(r.RemoteAddr)
//	if err != nil {
//		return &userIP
//	}
//	userIP = net.ParseIP(ip)
//	if userIP == nil {
//		return &userIP
//
//	}
//	fmt.Printf("%v", string(*userIP))
//	return &userIP
//}
