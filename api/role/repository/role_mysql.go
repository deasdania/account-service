package repository

import "account-metalit/api/models"

type IRoleMysql interface {
	GetRoleById(id string) (*models.Role, error)
	GetAllRole(orderby string) ([]*models.Role, error)
	CreateRole(name *models.FormName) error
	UpdateRoleName(id string, name string) error
	DeleteRoleById(id string) error
}
