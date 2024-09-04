package main

import (
	"language-parser/api/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define routes
	r.GET("/function", handlers.GetFunctions)

	r.POST("/parse", handlers.ParseContent)

	// Start the server on port 3000
	log.Println("Server is running on port 3000")
	r.Run(":3000")
}
