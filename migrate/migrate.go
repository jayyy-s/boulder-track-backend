package main

import (
	"github.com/jayyy-s/boulder-track-backend/initializers"
	"github.com/jayyy-s/boulder-track-backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Boulder{})
	initializers.DB.AutoMigrate(&models.User{})
}