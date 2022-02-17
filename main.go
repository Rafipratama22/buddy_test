package main

import (
	"log"

	"github.com/Rafipratama22/mnc_test.git/routes"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	routing routes.Server
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	route := routing.Start()
	route.Run(":8080")
}