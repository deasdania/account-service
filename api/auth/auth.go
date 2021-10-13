package auth

import (
	accountusecase "account-metalit/api/account/usecase"
	"account-metalit/api/auth/usecase"
	"account-metalit/api/models"
	"account-metalit/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	AccountUsecase accountusecase.IAccountUsecase
	AuthUsecase    usecase.IAuthUsecase
}

func (a Auth) Account(r *gin.RouterGroup) {
	r.POST(utilities.CHECK_AUTH, a.CheckAuth)                // private
	r.POST(utilities.LOGIN, a.Login)                         // public
	r.POST(utilities.LOGOUT, a.Logout)                       // private
	r.POST(utilities.REFRESH, a.Refresh)                     // public
	r.POST(utilities.CREATE_ACCOUNT_PUBLIC, a.CreateAccount) // public
}

func (a Auth) CheckAuth(context *gin.Context) {
	err := a.AuthUsecase.TokenValid(context.Request)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not Valid Token",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Valid Token",
	})
}

func (a Auth) CreateAccount(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirm_password := c.PostForm("confirm_password")

	form_register := models.FormRegister{
		Name:            name,
		Email:           email,
		Password:        password,
		ConfirmPassword: confirm_password,
	}
	response := a.AccountUsecase.CreateUser(form_register)
	c.JSON(response.Status, response)
}
func (a Auth) Login(context *gin.Context) {
	email := context.PostForm("email")
	password := context.PostForm("password")

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
