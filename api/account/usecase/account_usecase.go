package usecase

import "account-metalit/api/models"

type IAccountUsecase interface {
	GetUser(id string) *models.Users
	GetToken(email string, password string) (token string)
	CheckUserExist(email string) bool
	CheckPasswordLever(ps string) error
	CreateUser(form_register models.FormRegister) string
}
