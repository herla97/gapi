package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/herla97/gapi/controllers"
	"github.com/herla97/gapi/db"
	"github.com/herla97/gapi/middlewares"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}
}

func main() {
	// Setup Gin Engine config default.
	r := gin.Default()

	// Connect to db.
	db := db.Connect()
	defer db.Close()

	// Add db to Gin Context.
	r.Use(middlewares.ContextDB(db))

	// Add router group.
	v1 := r.Group("/api/v1")
	{
		v1.GET("/books", controllers.FindBooks)
	}

	r.Run() // listen and serve on 0.0.0.0:8080

}
