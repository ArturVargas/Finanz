package main

import (
	"fmt"

	"./controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

//DB initial Struct Config
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func buildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "golmer.db.elephantsql.com",
		Port:     5462,
		User:     "gjdmtzuo",
		DBName:   "gjdmtzuo",
		Password: "",
	}
	return &dbConfig
}

func dbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func main() {
	fmt.Println("Test go..!")

	db, err := gorm.Open("postgres", dbURL(buildDBConfig()))
	if err != nil {
		fmt.Println("dbStatus: ", err)
	}

	defer db.Close()

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
