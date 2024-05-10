//controller/testcontroller.go

package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// testcontroller handles requests for /test endpoint
func Test(c *gin.Context) {
    c.String(http.StatusOK, "i use Gin!")
}

// testPostRequest handles POST requests with request body
func TestPostRequest(c *gin.Context) {
	var requestBody struct {
		Name string `json:"name"`
		Age int `json:"age"`
	}

	// Bind the request body to the defined structure
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process the request body and return it
	c.JSON(http.StatusOK, requestBody)
}
