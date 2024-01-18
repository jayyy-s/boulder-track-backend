package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jayyy-s/boulder-track-backend/controllers"
	"github.com/jayyy-s/boulder-track-backend/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Router
	r := gin.Default()

	r.POST("/boulders", controllers.BouldersCreate)
	r.GET("/boulders", controllers.BouldersGetAll)
	r.GET("/boulders/:id", controllers.BouldersGetById)
	r.PUT("/boulders/:id", controllers.BouldersUpdate)
	r.DELETE("/boulders/:id", controllers.BouldersDelete)

	r.Run()
}