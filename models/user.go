package models

import (
	"backend/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

// Token JWT claims struct
type Token struct {
	UserID uint
	jwt.StandardClaims
}

//User a struct to rep user user
type User struct {
	gorm.Model
	UserType         int    `json:"userType"`
	Email            string `json:"email" gorm:"unique_index:idx_uniq_email"`
	Password         string `json:"password"`
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	BirthDate        string `json:"birthDate"`
	InsuranceCompany string `json:"insuranceCompany"`
	Company          string `json:"company"`
	ShortInfo        string `json:"shortInfo"`
	FullInfo         string `json:"fullInfo"`
	Photo            string `json:"photo"`
	Token            string `json:"token"`

	//
	// Email    string `json:"email"`
	// Password string `json:"password"`
	// Token    string `json:"token";sql:"-"`
}

// Validate incoming user details...
func (user *User) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(user.Email, "@") {
		return utils.Message(false, "Email address is required"), false
	}

	if len(user.Password) < 6 {
		return utils.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &User{}

	//check for errors and duplicate emails
	err := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return utils.Message(false, "Email address already in use by another user."), false
	}

	return utils.Message(false, "Requirement passed"), true
}

// Create method
func (user *User) Create() map[string]interface{} {

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	fmt.Println("%V", user.Email)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		return utils.Message(false, "Failed to create user, connection error.")
	}

	// Create new JWT token for the newly registered user
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString
	// fmt.Printf("user = %v \n\n", user)
	GetDB().Save(user)

	user.Password = "" //delete password

	response := utils.Message(true, "User has been created")
	response["user"] = user
	return response
}

// Update method
func (user *User) Update(token string) map[string]interface{} {
	//if resp, ok := page.Validate(); !ok {
	//	return resp
	//}
	var searchUser User
	if err := GetDB().Where("token = ?", token).First(&searchUser).Error; err != nil {

		return utils.Message(false, "Invalid token for update user")
	}

	GetDB().Model(&user).Where("ID = ?", searchUser.ID).Update(user)
	resp := utils.Message(true, "success")
	resp["user"] = searchUser
	return resp

}

// GetUserInfo method
func (user *User) GetUserInfo(token string) map[string]interface{} {
	//if resp, ok := page.Validate(); !ok {
	//	return resp
	//}
	var searchUser User
	if err := GetDB().Where("token = ?", token).First(&searchUser).Error; err != nil {

		return utils.Message(false, "Invalid token for update user")
	}
	resp := utils.Message(true, "success")
	resp["user"] = searchUser
	return resp
}

// GetAllUserInfo method
func (user *User) GetAllUserInfo() []*User {
	users := make([]*User, 0)
	err := GetDB().Find(&users).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return users
}

// Login method
func Login(email, password string) map[string]interface{} {

	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email address not found")
		}
		return utils.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return utils.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	user.Password = ""

	//Create JWT token
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString //Store the token in the response

	resp := utils.Message(true, "Logged In")
	resp["user"] = user
	return resp
}

// GetUser method
func GetUser(u uint) *User {

	acc := &User{}
	GetDB().Table("users").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
