package main

import (
	"fmt"

	config "./Config"
	controllers "./Controllers"
	models "./Models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db = config.DB

func main() {
	fmt.Println("Test go..!")

	db, err := gorm.Open("postgres", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("dbStatus: ", err)
	}

	defer db.Close()

	db.AutoMigrate(&models.User{})

	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.GET("/login/:id", controllers.Login)
	r.Run()
}

// Explicación
/*
	gin.Default() devuelve una instancia de gin engine con los registros de
	gin.DefaultWriter y el middleware de recuperacion y conectado.
	donde recibe los endpoints
*/

// Añadiendo DB PostgreSql.
/*
	En Go no existen los objetos pero en su lugar tenemos structs
	Cree un struct "DBConfig" con las que serán las propiedades de la bd
	Cree una funcion "buildDBConfig" que sera del tipo struct dentro de esta funcion
	hago una asignación especial a las propiedades del struct y se guarda en "dbConfig"
	y retorno su valor para poder usarlo fuera de la funcion.
	Creo una funcion "dbURL" que recibe un parametro del tipo struct y con "fmt.Sprintf" formateo una
	cadena de texto.
	Dentro de la funcion "main" hago la conexion a la bd con gorm.Open.
*/
