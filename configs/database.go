package database

import (
	"fmt"
	"golang-fiber-boilerplate/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB;

func loadCredentials() map[string] string {
	err := godotenv.Load()
	if err != nil {
	log.Fatal("Error loading .env.example file")
	}

	return map[string] string {
		"username": os.Getenv("DB_USERNAME"),
		"password" : os.Getenv("DB_PASSWORD"),
		"url" : os.Getenv("DB_URL"),
		"db_name" : os.Getenv("DB_NAME"),
	}
}

func Init() *gorm.DB {
	credentials := loadCredentials();
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", credentials["username"], credentials["password"], credentials["url"], credentials["db_name"]);
	dbInstance, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{});
	
	if dbErr != nil {
		panic("Failed to connect to database");
	}

	DB = dbInstance
	return dbInstance
 }

 func Migrate() {
	 err := DB.AutoMigrate(
		&models.User{},
	)
	 if err != nil {
		 return
	 }
 }