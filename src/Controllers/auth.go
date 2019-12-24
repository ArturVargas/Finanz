package controllers

import (
	"fmt"
	"net/http"

	config "../Config"
	models "../Models"
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
		fmt.Println(&user.Password)
		db.Create(&user)
		user.Password = ""
		c.JSON(http.StatusOK, &user)
	}
}

// Login endpoin for users
func Login(c *gin.Context) {
	// param := c.Params.ByName("id")
	// id, err := strconv.Atoi(param)
	// var user models.User
	// // err = models.GetUser(&user, id)
	// if err != nil {
	// 	c.AbortWithStatus(http.StatusNotFound)
	// } else {
	// 	c.JSON(http.StatusOK, user)
	// }
	return
}
