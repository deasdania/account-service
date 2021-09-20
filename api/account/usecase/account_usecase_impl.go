package usecase

import (
	"account-metalit/api/account/repository"
	"account-metalit/api/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
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

//The password strength must be letter size + number + sign, 9 digits or more
func (a accountUsecase) CheckPasswordLever(ps string) error {
	if len(ps) < 9 {
		return fmt.Errorf("password len is < 9")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return fmt.Errorf("password need num :%v", err)
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		return fmt.Errorf("password need a_z :%v", err)
	}
	if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
		return fmt.Errorf("password need A_Z :%v", err)
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		return fmt.Errorf("password need symbol :%v", err)
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
func (a accountUsecase) CreateUser(form_register models.FormRegister) string {
	fmt.Println(form_register)
	exist := a.CheckUserExist(form_register.Email)
	err := a.CheckPasswordLever(form_register.Password)
	if exist {
		return "user already exist"
	} else if err != nil {
		return err.Error()
	} else if form_register.Password != form_register.ConfirmPassword {
		return "password and confirm password not same"
	} else {
		hash, _ := a.HashPassword(form_register.Password)
		// return a.accountMysql.Begin(func(db *gorm.DB) error {
		user := models.Users{
			Name:     form_register.Name,
			Uuid:     uuid.New().String(),
			Email:    form_register.Email,
			Password: hash,
		}
		err := a.accountMysql.CreateAccount(&user)
		if err != nil {
			return "error when creating accounts"
		}
		// }
		// generate uuid, hash the password, create new
		return "success"
	}
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
