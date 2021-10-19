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

// Token Validator godoc
// @Summary Validate token
// @Description Register member API
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/check/authorize [post]
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

// CreateAccountMember godoc
// @Summary CreateAccountMember user
// @Description Register member API
// @Tags Public
// @Param Body formData models.BodyCreateAccount true "the body to create a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/createaccount [post]
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
	response := a.AccountUsecase.CreateUser(form_register, utilities.MEMBER)
	c.JSON(response.Status, response)
}

// Login godoc
// @Summary Login user
// @Description Logging in to get jwt token to access admin or user api by roles
// @Tags Public
// @Param Body formData models.BodyLoginInput true "the body to login user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/login [post]
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
	if token != nil {
		context.JSON(status, token)
		return
	} else {
		context.JSON(status, gin.H{
			"error": "password not match",
		})
		return
	}
}

// Logout godoc
// @Summary Logout Private
// @Description logout from system
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/logout [post]
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

// Refres Token godoc
// @Summary Refres Token Private
// @Description Refres Token
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/refresh [post]
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
