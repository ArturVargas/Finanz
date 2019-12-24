package main

import (
	"fmt"

	config "./Config"
	controllers "./Controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Test go..!")
	config.Init()

	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.GET("/login/:id", controllers.Login)
	router.GET("/benefits", controllers.Benefits)
	router.GET("/obligations", controllers.ObInfo)
	router.POST("/benefit", controllers.CreateBenefit)
	router.POST("/obligation", controllers.CreateOb)
	router.PUT("/benefit/:id", controllers.UpdateBenefit)
	router.PUT("/obligation/:id", controllers.UpdateOb)
	router.DELETE("/benefit/:id", controllers.DeleteBenefit)
	router.DELETE("/obligation/:id", controllers.DeleteOb)
	router.Run()
}
