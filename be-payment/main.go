package main

import (
	"azrielrisywan/be-assignment-payment/router"
	"fmt"
)

func main() {
	// Initialize Gin router
	r := router.SetupRouter()

	// Start the Gin server
	err := r.Run("0.0.0.0:8989")
	if err != nil {
		fmt.Println("Failed to start Gin server:", err)
	}
}
