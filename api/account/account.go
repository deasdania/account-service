package account

import (
	"account-metalit/api/account/usecase"
	authusecase "account-metalit/api/auth/usecase"

	"account-metalit/api/models"
	"account-metalit/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	// "github.com/satori/go.uuid"
	"github.com/google/uuid"
	"net/http"
	"reflect"
)

type Account struct {
	AccountUsecase usecase.IAccountUsecase
	AuthUsecase    authusecase.IAuthUsecase
}

func (a Account) Account(r *gin.RouterGroup) {
	r.GET(utilities.GET_ACCOUNT, a.GetUser)
	r.POST(utilities.CREATE_ACCOUNT, a.CreateAccount)
	r.POST(utilities.CHANGE_PASSWORD, a.ChangePassword)

	r.GET(utilities.GENERATE_UUID, a.GenerateUuid)
	r.POST("/test", func(c *gin.Context) { return })
}

func (a Account) GetUser(c *gin.Context) {
	user := a.AccountUsecase.GetUser("1")
	fmt.Println(reflect.TypeOf(user))
	c.JSON(user.Status, user)
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
	// fmt.Println(form_register)
	response := a.AccountUsecase.CreateUser(form_register)

	c.JSON(response.Status, response)

}

func (a Account) GenerateUuid(c *gin.Context) {
	myuuid := uuid.New().String()
	c.JSON(http.StatusOK, gin.H{
		"uuid": myuuid,
	})
}

func (a Account) ChangePassword(c *gin.Context) {
	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if errA != nil {
		fmt.Println(errA.Error())
	}
	email := metadata.Email
	old_password := c.PostForm("old_password")
	new_password := c.PostForm("new_password")
	confirm_password := c.PostForm("confirm_password")

	form_change_pass := models.FormChangePassword{
		Email:           email,
		OldPassword:     old_password,
		NewPassword:     new_password,
		ConfirmPassword: confirm_password,
	}
	response := a.AccountUsecase.ChangePassword(form_change_pass)

	c.JSON(response.Status, response)
}
