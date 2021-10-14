package usecase

import (
	"account-metalit/api/models"
	"account-metalit/api/role/repository"
	"account-metalit/response"
	// "fmt"
)

type roleUsecase struct {
	roleMysql      repository.IRoleMysql
	responseStruct response.IResponse
}

func (a roleUsecase) GetRoles(id string, orderby string) *response.Response {
	if id != "" {
		role, err := a.roleMysql.GetRoleById(id)
		if err != nil {
			return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
		}
		return a.responseStruct.ResponseSuccess(200, []string{"Get Role"}, map[string]*models.Role{
			"role": role,
		})
	}
	roles, err := a.roleMysql.GetAllRole(orderby)
	if err != nil {
		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
	}
	return a.responseStruct.ResponseSuccess(200, []string{"Get Role"}, map[string][]*models.Role{
		"roles": roles,
	})
}

func (a roleUsecase) CheckRoleExist(id string) bool {
	_, err := a.roleMysql.GetRoleById(id)
	if err != nil {
		return false
	}
	return true
}

func (a roleUsecase) CreateRole(form_name models.FormName) *response.Response {
	err := a.roleMysql.CreateRole(&form_name)
	if err != nil {
		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
	}
	return a.responseStruct.ResponseError(200, []string{"Create Role"}, map[string]string{
		"name": form_name.Name,
	})

}

func NewRoleUsecase(roleMysql repository.IRoleMysql, responseStruct response.IResponse) IRoleUsecase {
	return &roleUsecase{roleMysql: roleMysql, responseStruct: responseStruct}
}
