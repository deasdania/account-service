package api

import (
	"account-metalit/api/account"
	"account-metalit/api/account/repository"
	"account-metalit/api/account/usecase"
	"account-metalit/api/auth"
	"account-metalit/api/auth/authjwt"
	usecaseauth "account-metalit/api/auth/usecase"
	"account-metalit/api/role"
	reporole "account-metalit/api/role/repository"
	usecaserole "account-metalit/api/role/usecase"
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
	roleMysql := reporole.NewroleMysql(db)

	accountUsecase := usecase.NewAccountUsecase(accountMysql, responseStruct)
	roleUsecase := usecaserole.NewRoleUsecase(roleMysql, responseStruct)

	//AUTH
	authService := authjwt.JWTAuthService(redisDb)
	authUsecase := usecaseauth.NewAuthUsecase(authService, accountMysql)
	authController := auth.Auth{AccountUsecase: accountUsecase, AuthUsecase: authUsecase}
	authController.Account(public)
	private.Use(utilities.CheckRestClientJWT(authUsecase))

	//account
	accountController := account.Account{AccountUsecase: accountUsecase, AuthUsecase: authUsecase}
	accountController.Account(private)

	//role
	roleController := role.Role{RoleUsecase: roleUsecase, AuthUsecase: authUsecase}
	roleController.Role(private)

	fmt.Println(utilities.ACCOUNT_PORT)
	r.Run(fmt.Sprintf(":8089"))

}
