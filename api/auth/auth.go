package auth

import (
	"account-metalit/api/auth/usecase"
	"account-metalit/api/models"
	"account-metalit/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	AuthUsecase usecase.IAuthUsecase
}

func (a Auth) Account(r *gin.RouterGroup) {
	r.POST(utilities.LOGIN, a.Login)
	r.POST(utilities.LOGOUT, a.Logout)
	r.POST(utilities.REFRESH, a.Refresh)
}

func (a Auth) Login(context *gin.Context) {
	email := context.PostForm("email")
	password := context.PostForm("password")

	// token := a.AccountUsecase.GetToken(email, password)

	// context.JSON(http.StatusOK, gin.H{
	// 	"token": token,
	// })
	// var authReq models.Auth
	// err := context.ShouldBind(&authReq)
	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	token, status, err := a.AuthUsecase.Login(email, password)
	if err != nil {
		context.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(status, token)
}

func (a Auth) Logout(context *gin.Context) {
	metadata, err := a.AuthUsecase.ExtractTokenMetadata(context.Request)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}
	err = a.AuthUsecase.DeleteTokens(metadata)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}

func (a Auth) Refresh(context *gin.Context) {
	var authReq models.Auth
	err := context.ShouldBind(&authReq)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "binding json " + err.Error(),
		})
		return
	}
	result, status, err := a.AuthUsecase.Refresh(authReq.Token)
	if err != nil {
		if err != nil {
			context.JSON(status, err)
			return
		}
	}
	context.JSON(status, result)
}
