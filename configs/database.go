package database

import (
	"bnw-backend/models"
	"fmt"
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
    log.Fatal("Error loading .env file");
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
	db_instance, db_err := gorm.Open(mysql.Open(dsn), &gorm.Config{});
	
	if (db_err != nil) {
		panic("Failed to connect to database");
	}

	DB = db_instance;
	return db_instance;
 }

 func Migrate() {
	DB.AutoMigrate(
		&models.User{},
	);
 }