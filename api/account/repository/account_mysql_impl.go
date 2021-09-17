package repository

import (
	"account-metalit/api/models"
	"github.com/jinzhu/gorm"
)

type accountMysql struct {
	db *gorm.DB
}

func (a accountMysql) GetAccountById(id string) (*models.Users, error) {
	var user models.Users
	err := a.db.Debug().Table("users").First(&user, "id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

func (a accountMysql) GetAccountByEmail(email string) (*models.Users, error) {
	var user models.Users
	err := a.db.Debug().Table("users").First(&user, "email = ?", email)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

func NewAccountMysql(db *gorm.DB) IAccountMysql {
	return &accountMysql{db: db}
}
