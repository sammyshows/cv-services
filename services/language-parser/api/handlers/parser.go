package handlers

import (
	"language-parser/internal/parser"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParseContent(c *gin.Context) {
	var requestBody struct {
		Content  string `json:"content"`
		Language string `json:"language"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Step 1: Parse the content
	ast, err := parser.ParseContent(requestBody.Content, requestBody.Language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Step 2: Normalise the AST
	normalisedData, err := parser.NormaliseAST(ast, requestBody.Language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error normalizing AST: " + err.Error()})
		return
	}

	log.Println("Normalised data: ", normalisedData)

	// // Step 3: Generate Cypher Queries
	// queries := parser.GetCypherQueries(normalizedData)

	// // Step 4: Insert the data into Neo4j
	// if err := neo4j.InsertDataToNeo4j(queries); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data into Neo4j: " + err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "Code parsed and inserted into Neo4j successfully"})
}

func GetFunctions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetFunctions"})
}
