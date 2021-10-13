package repository

import "account-metalit/api/models"

type IAccountMysql interface {
	GetAccountById(id string) (*models.User, error)
	GetAccountByEmail(email string) (*models.User, error)
	CreateAccount(user *models.User) error
	UpdateAccountPassword(email string, hash string) error
}
