package models

import (
	"backend/utils"
	"github.com/jinzhu/gorm"
)

// FieldValue struct
type FieldValue struct {
	gorm.Model
	FieldID uint `json:"field_id"`
	//Key     string `json:"key"`
	Type   string `json:"type"`
	Value  string `json:"value"`
	Locale string `json:"Locale" gorm:"default:'DE'"`
}

// Validate method
func (fieldValue *FieldValue) Validate() (map[string]interface{}, bool) {
	if fieldValue.Value == "" || fieldValue.Type == "" {
		return utils.Message(false, "FieldValues key or type should be on the payload"), false
	}
	return utils.Message(true, "success"), true
}

// Create method
func (fieldValue *FieldValue) Create() map[string]interface{} {
	if resp, ok := fieldValue.Validate(); !ok {
		return resp
	}

	d := GetDB().Create(fieldValue)
	if d.Error != nil {
		return utils.Message(false, "Cant add new value for field")
	}

	resp := utils.Message(true, "success")
	resp["fieldValue"] = fieldValue
	return resp
}

// GetValueIDByKeyID method
func (fieldValue *FieldValue) GetValueIDByKeyID() *FieldValue {

	d := GetDB().Where("field_id = ?", fieldValue.FieldID).First(&fieldValue)
	if d.Error != nil {
		return nil
	}
	return fieldValue
}

// Update method
func (fieldValue *FieldValue) Update(Value string, ValueType string) map[string]interface{} {
	fieldValue.Value = Value
	fieldValue.Type = ValueType
	d := GetDB().Save(&fieldValue)
	if d.Error != nil {
		return utils.Message(false, "Cant update value or type for field")
	}
	resp := utils.Message(true, "success")
	resp["fieldValue"] = fieldValue
	return resp
}
