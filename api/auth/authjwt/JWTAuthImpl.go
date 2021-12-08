package authjwt

import (
	"auth-service/api/models"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type AuthJWT struct {
	client *redis.Client
}

func JWTAuthService(client *redis.Client) JWTService {
	return &AuthJWT{
		client: client,
	}
}

func (a AuthJWT) CreateToken(email string) (*models.TokenDetails, error) {
	td := &models.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 60).Unix()
	td.AccessUuid = uuid.New().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.AccessUuid + "++" + email

	log.Println(td.AccessUuid, td.RefreshUuid)

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["email"] = email
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["email"] = email
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}
func (a AuthJWT) CreateAuth(email string, td *models.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := a.client.Set(td.AccessUuid, email, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := a.client.Set(td.RefreshUuid, email, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	log.Println(at, rt)
	return nil
}
func (a AuthJWT) FetchAuth(authD *models.AccessDetails) (string, error) {
	email, err := a.client.Get(authD.AccessUuid).Result()
	if err != nil {
		return "", err
	}
	if authD.Email != email {
		return "", errors.New("unauthorized")
	}
	return email, nil
}
func (a AuthJWT) DeleteAuth(givenUuid string) (int64, error) {
	deleted, err := a.client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
func (a AuthJWT) DeleteTokens(authD *models.AccessDetails) error {
	//get the refresh uuid

	refreshUuid := fmt.Sprintf("%s++%s", authD.AccessUuid, authD.Email)
	log.Println(authD.AccessUuid, refreshUuid)
	//delete access token
	deletedAt, err := a.client.Del(authD.AccessUuid).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deletedRt, err := a.client.Del(refreshUuid).Result()
	if err != nil {
		return err
	}
	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}
