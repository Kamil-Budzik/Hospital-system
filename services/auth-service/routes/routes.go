package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamil-budzik/hospital-system/auth-service/models"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)
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

	// token, err := utils.GenerateToken(user.Email, user.ID)
	// if err != nil {
	// 	fmt.Println("Error in utils.GenerateToken", err)
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Could not authenticate user. Token error",
	// 	})
	// 	return
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   "OUDAB",
	})
}
