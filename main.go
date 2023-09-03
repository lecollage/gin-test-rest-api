package main

import (
	"fmt"
	"log"
	"os"
	"test_api/controller"
	"test_api/database"
	"test_api/middleware"
	"test_api/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("START")
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

func serveApplication() {
	mode := os.Getenv("MODE")

	var router *gin.Engine

	//DEV
	if mode == "DEV" {
		gin.SetMode(gin.DebugMode)
		router = gin.Default()
	}

	// PROD
	if mode == "PROD" {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}

	if router == nil {
		panic("MODE is not correct")
	}

	helloRoutes := router.Group("/app/test")
	helloRoutes.GET("/hello", controller.Hello)
	helloRoutes.GET("/ping", controller.Ping)
	helloRoutes.GET("/sleep", controller.Sleep)
	helloRoutes.GET("/cpu-load-sync", controller.CpuLoadSync)
	helloRoutes.GET("/cpu-load-async", controller.CpuLoadAsync)

	publicRoutes := router.Group("/app/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/app/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.GET("/entries/:id", controller.GetEntryById)
	protectedRoutes.GET("/entries", controller.GetAllEntries)
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.PUT("/entry/:id", controller.UpdateEntry)
	protectedRoutes.DELETE("/entry/:id", controller.DeleteEntry)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
