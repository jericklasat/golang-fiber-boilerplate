package models

import "gorm.io/gorm";

type User struct {
	gorm.Model
	Name string `json:"name"`;
	Email string `gorm:"unique" json:"email"`;
	Password string `json:"-"`;
	IsVerified bool;
	Level int `json:"level"`;
	Status int `json:"status"`;
}