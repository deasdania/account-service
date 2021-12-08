package api

import (
	"auth-service/api/account"
	"auth-service/api/account/repository"
	"auth-service/api/account/usecase"
	"auth-service/api/auth"
	"auth-service/api/auth/authjwt"
	usecaseauth "auth-service/api/auth/usecase"
	"auth-service/api/role"
	reporole "auth-service/api/role/repository"
	usecaserole "auth-service/api/role/usecase"
	"auth-service/config"
	"auth-service/middleware"
	"auth-service/response"

	// "github.com/gin-contrib/cors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(r *gin.Engine) {

	db := config.InitDb()
	redisDb := config.InitDbRedis()

	// r.Use(cors.Default())
	r.Use(middleware.CORSMiddleware())
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
	private.Use(middleware.CheckRestClientJWT(authUsecase))

	//account
	accountController := account.Account{AccountUsecase: accountUsecase, AuthUsecase: authUsecase}
	accountController.Account(private)

	//role
	roleController := role.Role{RoleUsecase: roleUsecase, AuthUsecase: authUsecase, AccountUsecase: accountUsecase}
	roleController.Role(private)

	r.Run(os.Getenv("PORT_RUN"))

}
