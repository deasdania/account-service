package usecase

import (
	"account-metalit/api/models"
	"account-metalit/response"
)

type IRoleUsecase interface {
	GetRoles(id string, orderby string) *response.Response
	CheckRoleExist(id string) bool
	CreateRole(form_name models.FormName) *response.Response
}
