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
	err := a.db.Debug().Model(&models.User{}).First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a accountMysql) GetAccountByUuid(uuid string) (*models.User, error) {
	var user models.User
	err := a.db.Debug().Model(&models.User{}).First(&user, "uuid = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a accountMysql) CreateAccount(user *models.User) error {
	return a.db.Debug().Model(&models.User{}).Create(&user).Error
}

func (a accountMysql) GetcodeVerification(uuid string) (*models.UserCodeVerification, error) {
	var userCode models.UserCodeVerification
	err := a.db.Debug().Model(&models.UserCodeVerification{}).First(&userCode, "user_uuid = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	return &userCode, nil
}

func (a accountMysql) CreatecodeVerification(codeVerification *models.UserCodeVerification) error {
	return a.db.Debug().Model(&models.UserCodeVerification{}).Create(&codeVerification).Error
}

func (a accountMysql) CreateUserRole(user *models.User, role *models.Role) error {
	formUserRole := models.UserRole{
		UserId: user.Id,
		RoleId: role.Id,
	}
	return a.db.Debug().Model(&models.UserRole{}).Create(&formUserRole).Error
}

func (a accountMysql) GetAccountByEmail(email string) (*models.User, error) {
	var user models.User
	err := a.db.Debug().Model(&models.User{}).First(&user, "email = ?", email)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

func (a accountMysql) UpdateAccountPassword(email string, hash string) error {
	return a.db.Debug().Model(&models.User{}).Where("email = ?", email).Update("password", hash).Error
}

func (a accountMysql) UpdateAccountAsVerified(email string) error {
	return a.db.Debug().Model(&models.User{}).Where("email = ?", email).Update("is_verified", true).Error
}

func (a accountMysql) DeleteAccountCodeVerification(uuid, code string) error {
	return a.db.Debug().Delete(&models.UserCodeVerification{}).Where("user_uuid = ? AND code = ?", uuid, code).Error
}

func NewAccountMysql(db *gorm.DB) IAccountMysql {
	return &accountMysql{db: db}
}
