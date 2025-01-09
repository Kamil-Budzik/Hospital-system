package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamil-budzik/hospital-system/auth-service/models"
	"github.com/kamil-budzik/hospital-system/auth-service/utils"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.POST("/verify", verify)
}

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not save user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not authenticate user",
		})
		return
	}

	token, err := utils.CreateToken(user.ID)
	if err != nil {
		fmt.Println("Error in utils.CreateToken", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not authenticate user. Token error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

func verify(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	_, err := utils.ValidateToken(authHeader)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unathorized",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User authorized",
	})
}
