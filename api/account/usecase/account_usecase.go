package usecase

import (
	"account-metalit/api/models"
	"account-metalit/response"
)

type IAccountUsecase interface {
	// GetUser(id string) *models.Users
	GetUser(id string) *response.Response
	GetToken(email string, password string) (token string)
	CheckUserExist(email string) bool
	CheckPasswordLever(ps string) []string
	// CreateUser(form_register models.FormRegister) string
	CreateUser(form_register models.FormRegister) *response.Response
}
