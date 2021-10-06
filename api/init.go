package api

import (
	"account-metalit/api/account"
	"account-metalit/api/account/repository"
	"account-metalit/api/account/usecase"
	"account-metalit/api/auth"
	"account-metalit/api/auth/authjwt"
	usecaseauth "account-metalit/api/auth/usecase"
	"account-metalit/config"
	"account-metalit/response"
	"account-metalit/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	// "os"
)

func Init(r *gin.Engine) {

	db := config.InitDb()
	redisDb := config.InitDbRedis()

	r.Use(utilities.CORSMiddleware())
	private := r.Group("/api")

	public := r.Group("/api")

	responseStruct := response.InitResponse()
	accountMysql := repository.NewAccountMysql(db)

	//AUTH
	authService := authjwt.JWTAuthService(redisDb)
	authUsecase := usecaseauth.NewAuthUsecase(authService, accountMysql)
	authController := auth.Auth{AuthUsecase: authUsecase}
	authController.Account(public)
	private.Use(utilities.CheckRestClientJWT(authUsecase))

	//account
	accountUsecase := usecase.NewAccountUsecase(accountMysql, responseStruct)
	accountController := account.Account{AccountUsecase: accountUsecase, AuthUsecase: authUsecase}
	accountController.Account(private)

	fmt.Println(utilities.ACCOUNT_PORT)
	r.Run(fmt.Sprintf(":8089"))

}
