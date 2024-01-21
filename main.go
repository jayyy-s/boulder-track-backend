package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jayyy-s/boulder-track-backend/controllers"
	"github.com/jayyy-s/boulder-track-backend/initializers"
	"github.com/jayyy-s/boulder-track-backend/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Router
	r := gin.Default()

	// boulder CRUD
	r.POST("/boulders", controllers.BouldersCreate)
	r.GET("/boulders", controllers.BouldersGetAll)
	r.GET("/boulders/:id", controllers.BouldersGetById)
	r.PUT("/boulders/:id", controllers.BouldersUpdate)
	r.DELETE("/boulders/:id", controllers.BouldersDelete)

	// user CRUD
	r.POST("/users", controllers.UsersCreate)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/users", controllers.UsersGetAll)
	r.GET("/users/:id", controllers.UsersGetById)
	r.PUT("/users/:id", controllers.UsersUpdate)
	r.DELETE("/users/:id", controllers.UsersDelete)

	r.Run()
}