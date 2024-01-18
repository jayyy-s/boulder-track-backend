package models

import "gorm.io/gorm"

type Boulder struct {
	gorm.Model
	Grade string
	PicLink string
	Gym string
}