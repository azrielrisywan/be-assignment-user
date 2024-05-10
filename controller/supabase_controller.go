package controller

import (
	"azrielrisywan/be-assignment-user/config"
	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
	"net/http"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(ctx *gin.Context) {
	var requestBody SignUpRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supabase := config.SetupSupabase()
	user, err := supabase.Auth.SignUp(ctx, supa.UserCredentials{
		Email: requestBody.Email,
		Password: requestBody.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func SignIn(ctx *gin.Context) {
	var requestBody SignUpRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supabase := config.SetupSupabase()
	user, err := supabase.Auth.SignIn(ctx, supa.UserCredentials{
		Email: requestBody.Email,
		Password: requestBody.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

