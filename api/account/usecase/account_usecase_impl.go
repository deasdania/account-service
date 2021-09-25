package usecase

import (
	"account-metalit/api/account/repository"
	"account-metalit/api/models"
	"account-metalit/response"
	"account-metalit/utilities"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
	// "strings"
	"time"
)

type accountUsecase struct {
	accountMysql   repository.IAccountMysql
	responseStruct response.IResponse
}

func (a accountUsecase) GetUser(id string) *response.Response {
	user, err := a.accountMysql.GetAccountById(id)
	if err != nil {
		fmt.Println(err.Error())
		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
	}

	return a.responseStruct.ResponseSuccess(200, []string{"Get User"}, user)
}

func (a accountUsecase) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (a accountUsecase) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a accountUsecase) GenerateJWT(user *models.Users) (string, error) {
	var secretJWT = []byte(os.Getenv(utilities.KEY_JWT))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = map[string]string{
		"id":    fmt.Sprintf("%d", user.Id),
		"uuid":  user.Uuid,
		"name":  user.Name,
		"email": user.Email,
	}
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(secretJWT)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil

}

//The password strength must be letter size + number + sign, 9 digits or more
func (a accountUsecase) CheckPasswordLever(ps string) []string {
	errMessage := []string{}
	isErr := false
	if len(ps) < 9 {
		errMessage = append(errMessage, "password len is < 9")
		isErr = true
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		errMessage = append(errMessage, "password need num")
		isErr = true
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		errMessage = append(errMessage, "password need a_z")
		isErr = true
	}
	if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
		errMessage = append(errMessage, "password need A_Z")
		isErr = true
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		errMessage = append(errMessage, "password need symbol")
		isErr = true
	}
	// errAll := strings.Join(errMessage, ", ")
	if isErr {
		// return fmt.Errorf(errAll)
		return errMessage
	}
	return nil
}

func (a accountUsecase) CheckUserExist(email string) bool {
	_, err := a.accountMysql.GetAccountByEmail(email)
	if err != nil {
		return false
	}
	return true
}

func (a accountUsecase) CreateUser(form_register models.FormRegister) *response.Response {
	fmt.Println(form_register)
	exist := a.CheckUserExist(form_register.Email)
	if exist {
		return a.responseStruct.ResponseError(400, []string{"user already exist"}, nil)
	}
	err := a.CheckPasswordLever(form_register.Password)
	if err != nil {
		return a.responseStruct.ResponseError(400, err, nil)
	}
	if form_register.Password != form_register.ConfirmPassword {
		return a.responseStruct.ResponseError(400, []string{"password and confirm password not same"}, nil)
	} else {
		hash, _ := a.HashPassword(form_register.Password)
		user := models.Users{
			Name:     form_register.Name,
			Uuid:     uuid.New().String(),
			Email:    form_register.Email,
			Password: hash,
		}
		err := a.accountMysql.CreateAccount(&user)
		if err != nil {
			return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
		}
		return a.responseStruct.ResponseError(200, []string{"Create User"}, map[string]string{
			"id":    fmt.Sprintf("%d", user.Id),
			"uuid":  user.Uuid,
			"name":  user.Name,
			"email": user.Email,
		})
	}
}

func (a accountUsecase) ChangePassword(form_change_pass models.FormChangePassword) *response.Response {
	user, err := a.accountMysql.GetAccountByEmail(form_change_pass.Email)
	if err != nil {
		fmt.Println(err.Error())
	}
	match := a.CheckPasswordHash(form_change_pass.OldPassword, user.Password)

	if !match {
		return a.responseStruct.ResponseError(400, []string{"password is not match"}, nil)
	}
	errCheck := a.CheckPasswordLever(form_change_pass.NewPassword)
	if errCheck != nil {
		return a.responseStruct.ResponseError(400, errCheck, nil)
	}
	if form_change_pass.NewPassword == form_change_pass.OldPassword {
		return a.responseStruct.ResponseError(400, []string{"new password and couldn't be same with the previous one"}, nil)
	}
	if form_change_pass.NewPassword != form_change_pass.ConfirmPassword {
		return a.responseStruct.ResponseError(400, []string{"new password and confirm password not the same"}, nil)
	} else {
		hash, _ := a.HashPassword(form_change_pass.NewPassword)
		errUpdate := a.accountMysql.UpdateAccountPassword(user.Email, hash)
		if errUpdate != nil {
			return a.responseStruct.ResponseError(400, []string{errUpdate.Error()}, nil)
		}
		return a.responseStruct.ResponseError(200, []string{"Password changed"}, map[string]string{
			"id":    fmt.Sprintf("%d", user.Id),
			"uuid":  user.Uuid,
			"name":  user.Name,
			"email": user.Email,
		})
	}
}

func (a accountUsecase) GetToken(email string, password string) (token string) {
	user, err := a.accountMysql.GetAccountByEmail(email)
	if err != nil {
		fmt.Println(err.Error())
	}
	match := a.CheckPasswordHash(password, user.Password)

	tok := ""
	if match {
		tok, err = a.GenerateJWT(user)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		tok = ""
	}
	return tok
}

func NewAccountUsecase(accountMysql repository.IAccountMysql, responseStruct response.IResponse) IAccountUsecase {
	return &accountUsecase{accountMysql: accountMysql, responseStruct: responseStruct}
}
