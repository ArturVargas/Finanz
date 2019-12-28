package main

import (
	"fmt"

	config "./Config"
	controllers "./Controllers"
	middlewares "./Middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Test go..!")
	config.Init()

	router := gin.Default()

	adminAuth := router.Group("/admin")
	adminAuth.Use(middlewares.SetMiddleWareToken())
	{
		adminAuth.POST("/benefit", controllers.CreateBenefit)
		adminAuth.POST("/obligation", controllers.CreateOb)
		adminAuth.PUT("/benefit/:id", controllers.UpdateBenefit)
		adminAuth.PUT("/obligation/:id", controllers.UpdateOb)
		adminAuth.DELETE("/benefit/:id", controllers.DeleteBenefit)
		adminAuth.DELETE("/obligation/:id", controllers.DeleteOb)
	}

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/benefits", controllers.Benefits)
	router.GET("/obligations", controllers.ObInfo)

	router.Run()
}
