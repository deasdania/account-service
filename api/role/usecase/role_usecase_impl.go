package usecase

import (
	"auth-service/api/models"
	"auth-service/api/role/repository"
	"auth-service/response"
	"fmt"
	"strconv"
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
	fmt.Println("create but check it first")
	getRon, _ := a.roleMysql.GetRoleByName(form_name.Name)
	if getRon == nil {
		role := models.Role{
			Name: form_name.Name,
		}
		err := a.roleMysql.CreateRole(&role)
		if err != nil {
			return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
		}
		return a.responseStruct.ResponseError(200, []string{"Create Role"}, map[string]string{
			"name": form_name.Name,
		})
	}
	return a.responseStruct.ResponseError(400, []string{"already in db, update it instead"}, map[string]*models.Role{
		"roles": getRon,
	})
}

func (a roleUsecase) UpdateRole(role models.Role) *response.Response {
	fmt.Println("update but check it first")
	idstr := strconv.Itoa(role.Id)
	getRon, err := a.roleMysql.GetRoleById(idstr)
	if err != nil {
		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
	}
	if getRon.Name == role.Name {
		return a.responseStruct.ResponseError(400, []string{"nothing to change"}, nil)
	}
	err = a.roleMysql.UpdateRoleName(idstr, role.Name)
	if err != nil {
		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
	}
	return a.responseStruct.ResponseError(200, []string{"Update Role"}, map[string]string{
		"name": role.Name,
	})
}

func NewRoleUsecase(roleMysql repository.IRoleMysql, responseStruct response.IResponse) IRoleUsecase {
	return &roleUsecase{roleMysql: roleMysql, responseStruct: responseStruct}
}
