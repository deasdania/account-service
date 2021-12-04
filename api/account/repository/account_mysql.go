package repository

import "account-metalit/api/models"

type IAccountMysql interface {
	// User Mysql Manipulation
	GetAccountById(id string) (*models.User, error)
	GetAccountByEmail(email string) (*models.User, error)
	CreateAccount(user *models.User) error
	UpdateAccountPassword(email string, hash string) error
	GetAccountByUuid(uuid string) (*models.User, error)
	UpdateAccountAsVerified(email string) error

	// UserRole Mysql Manipulation
	CreateUserRole(user *models.User, role *models.Role) error

	// UserCodeVerification Mysql Manipulation
	CreatecodeVerification(codeVerification *models.UserCodeVerification) error
	GetcodeVerification(uuid string) (*models.UserCodeVerification, error)
	DeleteAccountCodeVerification(uuid, code string) error
}
