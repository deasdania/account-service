package repository

import (
	"account-metalit/api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type roleMysql struct {
	db *gorm.DB
}

func (a roleMysql) GetRoleByName(name string) (*models.Role, error) {
	var role models.Role
	err := a.db.Debug().Table("roles").First(&role, "name = ?", name)
	if err.Error != nil {
		return nil, err.Error
	}
	return &role, nil
}

func (a roleMysql) GetRoleById(id string) (*models.Role, error) {
	var role models.Role
	err := a.db.Debug().Table("roles").First(&role, "id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &role, nil
}

func (a roleMysql) GetRoleByUserId(id string) (*models.Role, error) {
	var role models.Role
	err := a.db.Debug().Table("roles").First(&role, "user_id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &role, nil
}

func (a roleMysql) CheckUserIsAdmin(user_id string) (bool, error) {
	var role models.Role
	err := a.db.Debug().Table("user_roles").First(&role, "user_id = ? AND role_id = 1", user_id)
	if err.Error != nil {
		return false, err.Error
	}
	return true, nil
}

func (a roleMysql) GetAllRole(orderby string) ([]*models.Role, error) {
	roles := make([]*models.Role, 0)
	fmt.Println("orderby")
	fmt.Println(orderby)
	if orderby != "" {
		sortBy := strings.ToUpper(orderby)
		if sortBy == "DESC" || sortBy == "ASC" {
			orderby = "created_date " + orderby
		} else {
			orderby = "created_date desc"
		}
	} else {
		orderby = "created_date desc"
	}
	err := a.db.Debug().Table("roles").Order(orderby).Find(&roles)
	if err.Error != nil {
		return nil, err.Error
	}
	return roles, nil
}

// func (a accountMysql) GetAccount(limit, offset, order, search string, accountsResp *models.AccountRespPagination) ([]*models.Account, error) {
// 	accounts := make([]*models.Account, 0)
// 	accountsCount := make([]*models.Account, 0)
// 	accountsCount2 := make([]*models.Account, 0)
// 	offsetLimit, err := utilities.OffsetLimit(offset, limit)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if order == "" {
// 		order = "created_date desc"
// 	}
// 	newOffset, _ := strconv.Atoi(offset)
// 	var pagination models.Pagination
// 	var count int
// 	var count2 int
// 	err = a.db.Debug().Where("user_id like ? or username like ?", search, search).Limit(offsetLimit["limit"]).Offset(offsetLimit["offset"]).Order(order).Find(&accounts).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = a.db.Debug().Find(&accountsCount).Count(&count).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = a.db.Debug().Where("user_id like ? or username like ?", search, search).Order(order).Find(&accountsCount2).Count(&count2).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	p, err := pagination.PageData(newOffset, offsetLimit["limit"], count, order, count2)
// 	if err != nil {
// 		return nil, err
// 	}
// 	accountsResp.Pagination = p

// 	return accounts, nil
// }

func (a roleMysql) CreateRole(name *models.FormName) error {
	return a.db.Debug().Table("roles").Create(&name).Error
}

func (a roleMysql) UpdateRoleName(id string, name string) error {
	return a.db.Debug().Table("roles").Where("id = ?", id).Update("name", name).Error
}

func (a roleMysql) DeleteRoleById(id string) error {
	return a.db.Where("id = ?", id).Delete(&models.Role{}).Error
}

func NewroleMysql(db *gorm.DB) IRoleMysql {
	return &roleMysql{db: db}
}
