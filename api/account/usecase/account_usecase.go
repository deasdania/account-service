package usecase

import (
	"auth-service/api/models"
	"auth-service/response"
)

type IAccountUsecase interface {
	GetUser(id string) *response.Response
	CheckUserExist(email string) bool
	CheckPasswordLever(ps string) []string
	CreateUser(form_register models.FormRegister, member_type string) *response.Response
	ChangePassword(form_change_pass models.FormChangePassword) *response.Response
	CheckUserIsAdmin(email string) bool
	CheckUserCodeVerification(email string) (string, error)
	VerifiedUserAccount(email string, code models.BodyCodeVerification) *response.Response
}
