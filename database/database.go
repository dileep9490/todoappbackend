package database

import (
	"fmt"

	"github.com/dileep9490/todoapp/Backend/config"
	"github.com/dileep9490/todoapp/Backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	var (
		DATABASE = config.Config("PG_DATABASE")
		HOST     = config.Config("PG_HOST")
		PASSWORD = config.Config("PG_PASSWORD")
		PORT     = config.Config("PG_PORT")
		USER     = config.Config("PG_USER")
	)
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", HOST, USER, PASSWORD, DATABASE, PORT)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Println("Database connection error")
		panic(err)
	}
	DB = db
	fmt.Println("Database connected")

	DB.AutoMigrate(&models.User{}, &models.Todo{})
	fmt.Println("Database Migrated")

}
