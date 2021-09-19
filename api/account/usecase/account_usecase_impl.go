package usecase

import (
	"account-metalit/api/account/repository"
	"account-metalit/api/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type accountUsecase struct {
	accountMysql repository.IAccountMysql
}

func (a accountUsecase) GetUser(id string) *models.Users {
	user, err := a.accountMysql.GetAccountById(id)
	if err != nil {
		panic(err.Error())
	}
	return user
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
	var secretJWT = []byte(os.Getenv("SECRET_KEY_JWT"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(secretJWT)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil

}

func (a accountUsecase) GetToken(email string, password string) (token string) {
	user, err := a.accountMysql.GetAccountByEmail(email)
	if err != nil {
		panic(err.Error())
	}
	match := a.CheckPasswordHash(password, user.Password)

	tok := ""
	if match {
		tok, err = a.GenerateJWT(user)
		if err != nil {
			panic(err.Error())
		}
	} else {
		tok = ""
	}
	return tok
}

func NewAccountUsecase(accountMysql repository.IAccountMysql) IAccountUsecase {
	return &accountUsecase{accountMysql: accountMysql}
}
