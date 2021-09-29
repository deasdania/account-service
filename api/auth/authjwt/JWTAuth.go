package authjwt

import (
	"account-metalit/api/models"
)

type JWTService interface {
	//GenerateToken(email string, isUser bool) (string,error)
	//ValidateToken(token string) (*jwt.Token, error)
	CreateToken(email string) (*models.TokenDetails, error)
	CreateAuth(email string, td *models.TokenDetails) error
	DeleteTokens(authD *models.AccessDetails) error
	FetchAuth(authD *models.AccessDetails) (string, error)
	DeleteAuth(givenUuid string) (int64, error)
}
