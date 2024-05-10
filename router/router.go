// router/router.go

package router

import (
    "github.com/gin-gonic/gin"
    "azrielrisywan/be-assignment-user/controller"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Testing routes
    router.GET("/test", controller.Test)

	router.POST("/testPost", controller.TestPostRequest)

	// Supabase routes
	router.POST("/signup", controller.SignUp)

	router.POST("/signin", controller.SignIn)


    return router
}
