package usecase

import (
	"account-metalit/api/account/repository"
	"account-metalit/api/auth/authjwt"
	"account-metalit/api/models"
	"account-metalit/response"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"net/http"
	"os"
	"strings"
)

type authUsecase struct {
	jWtService     authjwt.JWTService
	accountMysql   repository.IAccountMysql
	responseStruct response.IResponse
}

func (a authUsecase) Refresh(refreshToken string) (*models.AuthReq, int, error) {
	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		fmt.Println("the error: ", err)
		return nil, http.StatusUnauthorized, err
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, http.StatusUnauthorized, err
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			return nil, http.StatusUnprocessableEntity, err
		}

		//Delete the previous Refresh Token
		deleted, err := a.jWtService.DeleteAuth(refreshUuid)
		if err != nil || deleted == 0 { //if any goes wrong
			return nil, http.StatusUnauthorized, errors.New("unauthorized")
		}
		//Create new pairs of refresh and access tokens
		ts, err := a.jWtService.CreateToken(claims["email"].(string))
		if err != nil {
			return nil, http.StatusUnauthorized, err
		}
		//save the tokens metadata to redis
		err = a.jWtService.CreateAuth(claims["email"].(string), ts)
		if err != nil {
			return nil, http.StatusUnauthorized, err
		}
		authReq := &models.AuthReq{
			AccessToken:  ts.AccessToken,
			RefreshToken: ts.RefreshToken,
		}
		return authReq, 200, nil
	} else {
		return nil, 401, errors.New("refresh expired")
	}
}

func (a authUsecase) FetchAuth(authD *models.AccessDetails) (string, error) {
	email, err := a.jWtService.FetchAuth(authD)
	if err != nil {
		return "", err
	}
	return email, nil
}
func (a authUsecase) DeleteTokens(authD *models.AccessDetails) error {
	err := a.jWtService.DeleteTokens(authD)
	if err != nil {
		return err
	}
	return nil
}
func (a authUsecase) ExtractTokenMetadata(r *http.Request) (*models.AccessDetails, error) {
	token, err := a.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		return &models.AccessDetails{
			AccessUuid: accessUuid,
			Email:      claims["email"].(string),
		}, nil
	}
	return nil, err
}
func (a authUsecase) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
func (a authUsecase) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := a.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
func (a authUsecase) TokenValid(r *http.Request) error {
	token, err := a.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return err
	}
	return nil
}

func (a authUsecase) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a authUsecase) Login(email string, password string) (*models.AuthReq, int, error) {
	// var result models.AuthGoogle
	// status, err := a.checkModeAuth(tokenId, &result)
	// if err != nil {
	// 	return nil, status, err
	// }
	user, err := a.accountMysql.GetAccountByEmail(email)
	if err != nil {
		fmt.Println(err.Error())
	}
	match := a.CheckPasswordHash(password, user.Password)

	if match {
		ts, err := a.jWtService.CreateToken(email)
		if err != nil {
			return nil, http.StatusUnprocessableEntity, err
		}
		saveErr := a.jWtService.CreateAuth(email, ts)
		if saveErr != nil {
			return nil, http.StatusUnprocessableEntity, err
		}
		authReq := &models.AuthReq{
			AccessToken:  ts.AccessToken,
			RefreshToken: ts.RefreshToken,
		}
		return authReq, 200, nil
	} else {
		return nil, http.StatusUnprocessableEntity, err
	}
}

func (a authUsecase) checkModeAuth(token string, result *models.AuthGoogle) (int, error) {
	if len(os.Args) > 2 && os.Args[2] == "authskip" {
		result.Email = "noauth"
		return 200, nil
	}
	url := "https://oauth2.googleapis.com/tokeninfo?id_token="
	response, err := http.Get(url + token)
	if err != nil {
		return 400, err
	}
	jsonDecoder := json.NewDecoder(response.Body)
	errDecoder := jsonDecoder.Decode(&result)
	if errDecoder != nil {
		return 400, err
	}
	if result.Error == "invalid_token" {
		return 400, errors.New("invalid token id")
	}
	return 200, nil
}

func NewAuthUsecase(jwt authjwt.JWTService, repo repository.IAccountMysql) IAuthUsecase {
	return &authUsecase{jWtService: jwt, accountMysql: repo}
}
