package account

import (
	"account-metalit/api/account/usecase"
	authusecase "account-metalit/api/auth/usecase"

	"account-metalit/api/models"
	"account-metalit/utilities"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Account struct {
	AccountUsecase usecase.IAccountUsecase
	AuthUsecase    authusecase.IAuthUsecase
}

func (a Account) Account(r *gin.RouterGroup) {
	r.GET(utilities.GET_ACCOUNT, a.GetUser)
	r.POST(utilities.CREATE_ACCOUNT, a.CreateAccount)
	r.POST(utilities.CHANGE_PASSWORD, a.ChangePassword)
	r.POST(utilities.RESEND_VERIFICATION_CODE, a.ResendVerificationCode)
	r.POST(utilities.PATCH_ACCOUNT_VERIFIED, a.VerifiedUser)

	r.GET(utilities.GENERATE_UUID, a.GenerateUuid)
	r.POST("/test", func(c *gin.Context) { return })
}

func (a Account) GetUser(c *gin.Context) {
	user := a.AccountUsecase.GetUser("1")
	fmt.Println(reflect.TypeOf(user))
	c.JSON(user.Status, user)
}

// Create Account by Admin godoc
// @Summary Admin can create another user using this API
// @Description Register new user API
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param Body formData models.BodyCreateAccount true "the body to create a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/create/account [post]
func (a Account) CreateAccount(c *gin.Context) {
	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if errA != nil {
		fmt.Println(errA.Error())
		return
	}
	fmt.Println(metadata)
	isAdmin := a.AccountUsecase.CheckUserIsAdmin(metadata.Email)
	if isAdmin {
		name := c.PostForm("name")
		email := c.PostForm("email")
		password := c.PostForm("password")
		confirm_password := c.PostForm("confirm_password")
		// role := c.PostForm("roles")

		form_register := models.FormRegister{
			Name:            name,
			Email:           email,
			Password:        password,
			ConfirmPassword: confirm_password,
		}
		response := a.AccountUsecase.CreateUser(form_register, utilities.ADMIN)
		c.JSON(response.Status, response)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "you are not allowed",
	})
}

// Generate uuid godoc
// @Summary Generate new uuid with this API
// @Description Generate uuid
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/generate/uuid [post]
func (a Account) GenerateUuid(c *gin.Context) {
	_, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if errA != nil {
		fmt.Println(errA.Error())
		return
	}
	myuuid := uuid.New().String()
	c.JSON(http.StatusOK, gin.H{
		"uuid": myuuid,
	})
}

// Change Password godoc
// @Summary Change user password
// @Description Change account password the logged on user with this API
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param Body formData models.BodyChangePasswordAccount true "the body to create a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /change/account/password [post]
func (a Account) ChangePassword(c *gin.Context) {
	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if errA != nil {
		fmt.Println(errA.Error())
		return
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

// Get the logged user Verification Code godoc
// @Summary Get verification code
// @Description Get the logged user Verification Code
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/codeverification [get]
func (a Account) ResendVerificationCode(c *gin.Context) {
	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if errA != nil {
		fmt.Println(errA.Error())
		return
	}
	email := metadata.Email
	code, err := a.AccountUsecase.CheckUserCodeVerification(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Printf("resend codeverification %v \n", code)
	c.JSON(http.StatusOK, gin.H{
		"message": "the verification code sent to the email",
	})
}

// Verified the user
// @Summary User Verified
// @Description Match the code from the request with what it should be, and delete if it just the same and update the user as verified
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/user/verified [post]
func (a Account) VerifiedUser(c *gin.Context) {
	metadata, err := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	email := metadata.Email
	code := new(models.BodyCodeVerification)
	err = c.Bind(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := a.AccountUsecase.VerifiedUserAccount(email, *code)
	c.JSON(response.Status, response)
}
