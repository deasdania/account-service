package account

import (
	"account-metalit/api/account/usecase"
	"account-metalit/utilities"
	"fmt"
	"github.com/gin-gonic/gin"

	"net/http"
	"reflect"
)

type Account struct {
	AccountUsecase usecase.IAccountUsecase
}

func (a Account) Account(r *gin.RouterGroup) {
	r.GET(utilities.GET_ACCOUNT, a.GetUser)
	r.POST(utilities.LOGIN, a.Login)

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
