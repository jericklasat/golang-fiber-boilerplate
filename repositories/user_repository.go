package repositories

import (
	"fmt"
	"golang-fiber-boilerplate/data/dto"
	"golang-fiber-boilerplate/models"
)

func CreateUser(userData dto.User) string {
	err := DB.Create(userData).Error
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return "OK"
}

func GetUserByEmail(email string) models.User {
	var userData models.User
	err := DB.Where("email = ?", email).First(&userData)
	if err != nil {
		fmt.Println(err)
	}
	return userData
}