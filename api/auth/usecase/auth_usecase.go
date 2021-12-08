package usecase

import (
	"auth-service/api/models"
	"net/http"
)

type IAuthUsecase interface {
	Login(email string, password string) (*models.AuthReq, int, error)
	ExtractTokenMetadata(r *http.Request) (*models.AccessDetails, error)
	DeleteTokens(*models.AccessDetails) error
	TokenValid(r *http.Request) error
	FetchAuth(authD *models.AccessDetails) (string, error)
	Refresh(token string) (*models.AuthReq, int, error)
	CreateVerificationCode(uuid models.UserUuid) error
}
