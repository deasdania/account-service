package usecase

import "account-metalit/api/models"

type IAccountUsecase interface {
	GetUser(id string) *models.Users
	GetToken(email string, password string) (token string)
}
