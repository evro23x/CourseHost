package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// golint comment for "_" import not in main or test package
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	//db.DropTable(&FieldValue{}, &Field{}, &Page{}, &User{})
	db.Debug().AutoMigrate(&User{}, &Page{}, &Field{}, &FieldValue{}, &Course{}, &Lesson{}, &Quiz{})
	//db.Model(&FieldValue{}).RemoveForeignKey("field_id", "fields(id)")
	//db.Model(&FieldValue{}).RemoveForeignKey("page_id", "pages(id)")

	db.Model(&Field{}).AddForeignKey("page_id", "pages(id)", "RESTRICT", "RESTRICT")
	db.Model(&FieldValue{}).AddForeignKey("field_id", "fields(id)", "RESTRICT", "RESTRICT")

}

// GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
