package account

import (
	"account-metalit/api/account/usecase"
	"account-metalit/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"reflect"
)

type Account struct {
	AccountUsecase usecase.IAccountUsecase
}

func (a Account) Account(r *gin.RouterGroup) {
	r.GET(utilities.GET_ACCOUNT, a.GetUser)
	r.POST(utilities.LOGIN, a.Login)
	r.POST(utilities.CREATE_ACCOUNT, a.CreateAccount)

	r.GET(utilities.GENERATE_UUID, a.GenerateUuid)
	r.POST("/test", func(c *gin.Context) { return })
}

func (a Account) GetUser(c *gin.Context) {
	user := a.AccountUsecase.GetUser("1")
	fmt.Println(reflect.TypeOf(user))
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (a Account) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	token := a.AccountUsecase.GetToken(email, password)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (a Account) CreateAccount(c *gin.Context) {
	// email := c.PostForm("email")
	password := c.PostForm("password")
	confirm_password := c.PostForm("confirm_password")

	if password != confirm_password {
		c.JSON(http.StatusOK, gin.H{
			"message": "password and confirm password not same",
		})
	} else {

	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"email":            email,
	// 	"password":         password,
	// 	"confirm_password": confirm_password,
	// })
}

func (a Account) GenerateUuid(c *gin.Context) {
	myuuid := uuid.NewV4()
	c.JSON(http.StatusOK, gin.H{
		"uuid": myuuid,
	})
}
