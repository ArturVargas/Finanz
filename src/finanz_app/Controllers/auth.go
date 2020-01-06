package controllers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	config "../Config"
	models "../Models"
	utils "../Utils"
	"github.com/gin-gonic/gin"
)

// Register endpoint for users
func Register(c *gin.Context) {
	var user models.User
	var db = config.GetDB()
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		db.Create(&user)
		user.Password = ""
		c.JSON(http.StatusOK, &user)
	}
}

// Login endpoin for users
func Login(c *gin.Context) {
	var user models.User
	var db = config.GetDB()
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	password := user.Password
	err = db.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = models.VerifyPwd(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, error := utils.CreateToken(user.ID, user.Name, user.Email)
	if error != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, token)
}
