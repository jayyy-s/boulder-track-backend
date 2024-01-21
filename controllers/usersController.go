package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jayyy-s/boulder-track-backend/initializers"
	"github.com/jayyy-s/boulder-track-backend/models"
	"golang.org/x/crypto/bcrypt"
)

var userBody struct {
	Username     string
	Password     string
	Email        string
	HeightInches int
	WeightPounds int
	Gender       string
	Birthdate    string
}

func UsersCreate(c *gin.Context) {
	// get data from body
	c.Bind(&userBody)

	// create boulder
	user := models.User{
		Username:     userBody.Username,
		Password:     userBody.Password,
		Email:        userBody.Email,
		HeightInches: userBody.HeightInches,
		WeightPounds: userBody.WeightPounds,
		Gender:       userBody.Gender,
		Birthdate:    userBody.Birthdate,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return
	c.JSON(200, gin.H{
		"user": user,
	})
}

// Gets all the boulders
func UsersGetAll(c *gin.Context) {
	// Get the boulders
	var users []models.User
	initializers.DB.Find(&users)

	// Respond with them
	c.JSON(200, gin.H{
		"users": users,
	})
}

// Gets boulder by primary key (id)
func UsersGetById(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// Get the boulder by id
	var user models.User
	initializers.DB.First(&user, id)

	// Respond with them
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersUpdate(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get data from req body
	c.Bind(&userBody)

	// find the boulder being updated
	var user models.User
	initializers.DB.First(&user, id)

	// update
	initializers.DB.Model(&user).Updates(models.User{
		Username:     userBody.Username,
		Password:     userBody.Password,
		Email:        userBody.Email,
		HeightInches: userBody.HeightInches,
		WeightPounds: userBody.WeightPounds,
		Gender:       userBody.Gender,
		Birthdate:    userBody.Birthdate,
	})

	// respond
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersDelete(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// delete post
	initializers.DB.Delete(&models.User{}, id)

	// respond
	c.Status(200)
}

// Actual signup function
func Signup(c *gin.Context) {
	// get signup credentials
	if c.Bind(&userBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(userBody.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// create the user
	user := models.User{
		Username:     userBody.Username,
		Password:     string(hash),
		Email:        userBody.Email,
		HeightInches: userBody.HeightInches,
		WeightPounds: userBody.WeightPounds,
		Gender:       userBody.Gender,
		Birthdate:    userBody.Birthdate,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H {})
}

func Login(c *gin.Context) {
	// get username and password from req body
	if c.Bind(&userBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// look up requested user
	var user models.User
	initializers.DB.First(&user, "username = ?", userBody.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})

		return
	}

	// compare given pw with saved pw hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userBody.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})

		return
	}

	// generate jwt token
	// JSON Web Token token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// send token back
	c.SetSameSite(http.SameSiteLaxMode)
	// secure = true when not localhost
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "" , "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"message": "I'm logged in",
	})
}