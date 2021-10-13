package repository

import (
	"account-metalit/api/models"
	"github.com/jinzhu/gorm"
)

type accountMysql struct {
	db *gorm.DB
}

func (a accountMysql) GetAccountById(id string) (*models.User, error) {
	var user models.User
	err := a.db.Debug().Table("users").First(&user, "id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

func (a accountMysql) CreateAccount(user *models.User) error {
	result := a.db.Debug().Table("users").Create(&user)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return result.Error
}

func (a accountMysql) GetAccountByEmail(email string) (*models.User, error) {
	var user models.User
	err := a.db.Debug().Table("users").First(&user, "email = ?", email)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

func (a accountMysql) UpdateAccountPassword(email string, hash string) error {
	err := a.db.Debug().Table("users").Where("email = ?", email).Update("password", hash).Error
	if err != nil {
		return err
	}
	return err
}
func NewAccountMysql(db *gorm.DB) IAccountMysql {
	return &accountMysql{db: db}
}
