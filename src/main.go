package main

import (
	"fmt"

	"./controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Test go..!")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/register", controllers.Register)
	r.Run()
}

// Explicación
/*
	gin.Default() devuelve una instancia de gin engine con los registros de
	gin.DefaultWriter y el middleware de recuperacion y conectado.
	donde recibe los endpoints
*/
