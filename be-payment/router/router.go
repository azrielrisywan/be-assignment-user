// router/router.go

package router

import (
	"azrielrisywan/be-assignment-payment/controller"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"azrielrisywan/be-assignment-payment/middleware"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
	var hmacSecret = "pDnYuxHNGugqD6u/q20ShEFX32uIDNFTPH3CjLZjPSES/N7QvZr+v+eDOCi31F7FbQFrzCgLqngGUolnvUXzqw=="
	
	// Enable CORS for all origins. This is not recommended for production usage.
    // Use a whitelist of allowed origins instead.
    corsConfig := cors.Config{
        AllowOrigins: []string{"*"},
        AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
    }

	router := gin.Default()
	router.Use(cors.New(corsConfig))

	// Testing routes
	router.GET("/test", controller.Test)

	router.POST("/testPost", controller.TestPostRequest)

	// Payment routes
	router.POST("/payment/send", middleware.AuthMiddleware(hmacSecret), controller.SendPayment)

	router.POST("/payment/withdraw", middleware.AuthMiddleware(hmacSecret), controller.WithdrawPayment)

	// Transaction routes
	router.POST("/transaction/list-by-user", middleware.AuthMiddleware(hmacSecret), controller.TransactionListByUser)

	return router
}
