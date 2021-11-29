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
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

func Init(r *gin.Engine) {

	db := config.InitDb()
	redisDb := config.InitDbRedis()

	// r.Use(cors.Default())
	r.Use(utilities.CORSMiddleware())
	private := r.Group("/api")

	public := r.Group("/api")

	ra := r.Group("/")
	ra.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ra.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/swagger/index.html")
	})

	responseStruct := response.InitResponse()
	accountMysql := repository.NewAccountMysql(db)
	roleMysql := reporole.NewroleMysql(db)

	roleUsecase := usecaserole.NewRoleUsecase(roleMysql, responseStruct)
	accountUsecase := usecase.NewAccountUsecase(roleMysql, accountMysql, responseStruct)

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
	roleController := role.Role{RoleUsecase: roleUsecase, AuthUsecase: authUsecase, AccountUsecase: accountUsecase}
	roleController.Role(private)

	fmt.Println(utilities.ACCOUNT_PORT)
	r.Run(os.Getenv("PORT_RUN"))

}
