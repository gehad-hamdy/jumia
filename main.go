package main

import (
	"Flash/pkg/db"
	"Flash/routes"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	_ "os"
)

var (
	router = gin.Default()
)

func main() {
	gin.SetMode(gin.DebugMode)
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}
	_, connectionErr := db.Connect()
	if connectionErr != nil {
		log.Fatal(connectionErr.Error())
	}
	router := gin.Default()
	routes.CustomerRoutes(router)
	log.Fatal(router.Run(":3000"))
}