package main

import (
	"backend/app"
	"backend/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/me/contacts", controllers.GetContactsFor).Methods("GET")

	router.HandleFunc("/user/add", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/edit", controllers.EditUser).Methods("POST")
	router.HandleFunc("/user/get", controllers.GetUser).Methods("POST")
	router.HandleFunc("/user/get_all", controllers.GetAllUser).Methods("POST")
	router.HandleFunc("/user/login", controllers.Authenticate).Methods("POST")

	router.HandleFunc("/page/get_all", controllers.GetPageList).Methods("POST")
	router.HandleFunc("/page/get", controllers.GetPageById).Methods("POST")
	router.HandleFunc("/page/add", controllers.AddPageInfo).Methods("POST")
	router.HandleFunc("/page/edit", controllers.EditPageInfo).Methods("POST")
	router.HandleFunc("/page/remove", controllers.RemovePage).Methods("POST")

	router.HandleFunc("/field/add", controllers.AddField).Methods("POST")
	router.HandleFunc("/field/edit", controllers.EditField).Methods("POST")
	router.HandleFunc("/field/remove", controllers.RemoveField).Methods("POST")

	// courses
	//
	// uid	number
	// type	number
	// name	string
	// description	string
	// image	string
	// smalltext	string
	// text	string
	// weekCount	number
	// price	number

	router.HandleFunc("/course/add", controllers.AddCourse).Methods("POST")
	router.HandleFunc("/course/get", controllers.GetCourse).Methods("POST")
	router.HandleFunc("/course/get_all", controllers.GetAllCourses).Methods("POST")
	router.HandleFunc("/course/edit", controllers.EditCourse).Methods("POST")
	router.HandleFunc("/course/remove", controllers.RemoveCourse).Methods("POST")

	// lessons
	//
	// pid	number
	// weekNumber	number
	// title	string
	// text	string
	// order	number
	// imageURL	string
	// audioURL	string
	// pdfURL	string

	router.HandleFunc("/lesson/add", controllers.AddLesson).Methods("POST")
	router.HandleFunc("/lesson/get", controllers.GetLesson).Methods("POST")
	router.HandleFunc("/lesson/get_all", controllers.GetAllLessons).Methods("POST")
	router.HandleFunc("/lesson/edit", controllers.EditLesson).Methods("POST")
	router.HandleFunc("/lesson/remove", controllers.RemoveLesson).Methods("POST")

	// Quiz
	//
	// Title       string `json:"title"`
	// Description string `json:"description"`
	// CourseID    int    `json:"courseID"`
	// ImageURL    string `json:"imageURL"`
	// VideoURL    string `json:"videoURL"`
	// Main        int    `json:"main"`

	router.HandleFunc("/quiz/add", controllers.AddQuiz).Methods("POST")
	router.HandleFunc("/quiz/get", controllers.GetQuiz).Methods("POST")
	router.HandleFunc("/quiz/get_all", controllers.GetAllQuiz).Methods("POST")
	router.HandleFunc("/quiz/edit", controllers.EditQuiz).Methods("POST")
	router.HandleFunc("/quiz/remove", controllers.RemoveQuiz).Methods("POST")

	port := os.Getenv("app_port")
	if port == "" {
		port = "8989"
	}

	headersOk := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	err := http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router))
	if err != nil {
		fmt.Print(err)
	}
}
