package role

import (
	accountusecase "account-metalit/api/account/usecase"
	authusecase "account-metalit/api/auth/usecase"
	"account-metalit/api/models"
	"account-metalit/api/role/usecase"

	// "account-metalit/api/models"
	"account-metalit/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	// "github.com/satori/go.uuid"
	// "github.com/google/uuid"
	"net/http"
	// "reflect"
)

type Role struct {
	RoleUsecase    usecase.IRoleUsecase
	AuthUsecase    authusecase.IAuthUsecase
	AccountUsecase accountusecase.IAccountUsecase
}

func (a Role) Role(r *gin.RouterGroup) {
	r.GET(utilities.GET_ROLE, a.GetRole)        //Hanya untuk Admin
	r.POST(utilities.CREATE_ROLE, a.CreateRole) //Hanya untuk Admin
}

func (a Role) GetRole(c *gin.Context) {
	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if errA != nil {
		fmt.Println(errA.Error())
	}
	fmt.Println(metadata)
	isAdmin := a.AccountUsecase.CheckUserIsAdmin(metadata.Email)
	if isAdmin {
		id, _ := c.GetQuery("id")
		orderby, _ := c.GetQuery("orderby")
		response := a.RoleUsecase.GetRoles(id, orderby)
		c.JSON(response.Status, response)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "you are not allowed",
	})
}

func (a Role) CreateRole(c *gin.Context) {
	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if errA != nil {
		fmt.Println(errA.Error())
	}
	fmt.Println(metadata)
	isAdmin := a.AccountUsecase.CheckUserIsAdmin(metadata.Email)
	if isAdmin {
		name := c.PostForm("role_name")
		form_name := models.FormName{Name: name}
		response := a.RoleUsecase.CreateRole(form_name)
		c.JSON(response.Status, response)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "you are not allowed",
	})

}
