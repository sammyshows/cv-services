package main

import (
	"language-parser/api/handlers"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// Define routes
	r.GET("/functions", handlers.GetFunctions)
	r.POST("/parse", handlers.ParseContent)

	// Start the server on port 3000
	log.Println("---------------------------------")
	log.Println("Server is running on port 3000")
	log.Println("---------------------------------")
	r.Run(":3000")
}
