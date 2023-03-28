package main

import (
    "net/http"	
    "github.com/gin-gonic/gin"
)

// sample REST Resource
type generativeAI struct {
	AIProduct string `json:name`
	Creator string `json:creator`
}

// sample data
var AIs = []generativeAI{
	{AIProduct: "GPT-4", Creator: "OpenAI"},
	{AIProduct: "AI Foundations", Creator: "Nvidia"},
	{AIProduct: "Bard", Creator: "Alphabet"},
	{AIProduct: "Firefly", Creator: "Adobe"},
	{AIProduct: "Alpaca", Creator: "Stanford"},
}

func main() {
    router := gin.Default()
    router.GET("/AIs", getAIs)
	router.POST("/AIs", postAI)
	router.GET("/AIs/:creator", getAIByname)
    router.Run("localhost:8080")
}

// GET endpoint implementation
func getAIs(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, AIs)
}

// POST endpoint implementation
func postAI(c *gin.Context) {
    var newAI generativeAI
    // Call BindJSON to bind the received JSON to AI
    if err := c.BindJSON(&newAI); err != nil {
        return
    }
    // Add the new AI to the slice.
    AIs = append(AIs, newAI)
    c.IndentedJSON(http.StatusCreated, newAI)
}


func getAIByname(c *gin.Context) {
    creator := c.Param("creator")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range AIs {
        if a.Creator == creator {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "AI not found"})
}
