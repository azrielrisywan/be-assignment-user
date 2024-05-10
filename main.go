package main

import (
    "fmt"
    "azrielrisywan/be-assignment-user/router"
)

func main() {
    // Initialize Gin router
    r := router.SetupRouter()

    // Start the Gin server
    err := r.Run("localhost:8888")
    if err != nil {
        fmt.Println("Failed to start Gin server:", err)
    }
}
