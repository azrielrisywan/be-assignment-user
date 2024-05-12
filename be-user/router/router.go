// router/router.go

package router

import (
	"azrielrisywan/be-assignment-user/controller"
	"azrielrisywan/be-assignment-user/middleware"
	"azrielrisywan/be-assignment-user/docs"

	swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	docs.SwaggerInfo.BasePath = ""
	router.Use(cors.New(corsConfig))

    // Testing routes
    router.GET("/test", controller.Test)

	router.POST("/testPost", controller.TestPostRequest)

	// Supabase routes
	router.POST("/signup", controller.SignUp)

	router.POST("/signin", controller.SignIn)

	// Account routes
	router.POST("/getAccountsByUser", middleware.AuthMiddleware(hmacSecret), controller.GetAccountsByUser)

	// Payment routes
	router.POST("/getPaymentsListByUser", middleware.AuthMiddleware(hmacSecret), controller.GetPaymentsListByUser)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))


    return router
}
