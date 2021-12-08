package usecase

import (
	"auth-service/api/models"
	"auth-service/response"
)

type IRoleUsecase interface {
	GetRoles(id string, orderby string) *response.Response
	CheckRoleExist(id string) bool
	CreateRole(form_name models.FormName) *response.Response
	UpdateRole(role models.Role) *response.Response
}
