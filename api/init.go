package api

import (
	"account-metalit/api/account"
	"account-metalit/api/account/repository"
	"account-metalit/api/account/usecase"
	"account-metalit/config"
	"account-metalit/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	// "os"
)

func Init(r *gin.Engine) {

	db := config.InitDb()

	r.Use(utilities.CORSMiddleware())
	v1 := r.Group("/api")

	//account
	accountMysql := repository.NewAccountMysql(db)
	accountUsecase := usecase.NewAccountUsecase(accountMysql)
	accountController := account.Account{AccountUsecase: accountUsecase}
	accountController.Account(v1)

	r.Run(fmt.Sprintf(":8089"))

}
