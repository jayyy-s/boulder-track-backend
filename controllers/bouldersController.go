package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jayyy-s/boulder-track-backend/initializers"
	"github.com/jayyy-s/boulder-track-backend/models"
)

var boulderBody struct {
	Grade   string
	PicLink string
	Gym     string
	UserID  int
}

func BouldersCreate(c *gin.Context) {
	// get data from body
	c.Bind(&boulderBody)

	// create boulder
	boulder := models.Boulder{
		Grade:   boulderBody.Grade,
		PicLink: boulderBody.PicLink,
		Gym:     boulderBody.Gym,
		UserID:     boulderBody.UserID,
	}

	result := initializers.DB.Create(&boulder)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return
	c.JSON(200, gin.H{
		"boulder": boulder,
	})
}

// Gets all the boulders
func BouldersGetAll(c *gin.Context) {
	// Get the boulders
	var boulders []models.Boulder
	initializers.DB.Find(&boulders)

	// Respond with them
	c.JSON(200, gin.H{
		"boulders": boulders,
	})
}

// Gets boulder by primary key (id)
func BouldersGetById(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// Get the boulder by id
	var boulder models.Boulder
	initializers.DB.First(&boulder, id)

	// Respond with them
	c.JSON(200, gin.H{
		"boulder": boulder,
	})
}

func BouldersUpdate(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get data from req body
	c.Bind(&boulderBody)

	// find the boulder being updated
	var boulder models.Boulder
	initializers.DB.First(&boulder, id)

	// update
	initializers.DB.Model(&boulder).Updates(models.Boulder{
		Grade:   boulderBody.Grade,
		PicLink: boulderBody.PicLink,
		Gym:     boulderBody.Gym,
		UserID:     boulderBody.UserID,
	})

	// respond
	c.JSON(200, gin.H{
		"boulder": boulder,
	})
}

func BouldersDelete(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// delete post
	initializers.DB.Delete(&models.Boulder{}, id)

	// respond
	c.Status(200)
}
