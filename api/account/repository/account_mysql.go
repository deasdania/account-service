package repository

import "account-metalit/api/models"

type IAccountMysql interface {
	GetAccountById(id string) (*models.Users, error)
	GetAccountByEmail(email string) (*models.Users, error)
	CreateAccount(user *models.Users) error
	UpdateAccountPassword(email string, hash string) error
}
