package config

import (
	"fmt"
	"log"
	"os"

	models "../Models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// DB set
var DB *gorm.DB
var err error

// Init creates a connection to mysql database and
// migrates any new models
func Init() {
	Host := "otto.db.elephantsql.com"
	Port := "5432"
	User := os.Getenv("DB_NAME")
	DBName := os.Getenv("DB_NAME")
	Password := os.Getenv("DB_PASS")

	dbInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		User,
		Password,
		Host,
		Port,
		DBName,
	)
	DB, err = gorm.Open("postgres", dbInfo)
	if err != nil {
		log.Println("Failed to Conect DB")
		panic(err)
	}
	log.Println("Database Connected")

	if !DB.HasTable(&models.User{}) {
		err := DB.CreateTable(&models.User{})
		if err != nil {
			log.Println("Table already exists")
		}
	}

	DB.AutoMigrate(&models.User{}, &models.Obligaciones{}, &models.Deducciones{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return DB
}

//CloseDB ...
func CloseDB() {
	DB.Close()
}
