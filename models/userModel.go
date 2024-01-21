package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Password     string
	Email        string `gorm:"unique"`
	HeightInches int
	WeightPounds int
	Gender       string
	Birthdate    string
}
