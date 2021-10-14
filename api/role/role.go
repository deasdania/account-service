package role

import (
	authusecase "account-metalit/api/auth/usecase"
	"account-metalit/api/models"
	"account-metalit/api/role/usecase"

	// "account-metalit/api/models"
	"account-metalit/utilities"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "github.com/satori/go.uuid"
	// "github.com/google/uuid"
	// "net/http"
	// "reflect"
)

type Role struct {
	RoleUsecase usecase.IRoleUsecase
	AuthUsecase authusecase.IAuthUsecase
}

func (a Role) Role(r *gin.RouterGroup) {
	r.GET(utilities.GET_ROLE, a.GetRole)
	r.POST(utilities.CREATE_ROLE, a.CreateRole)
}

func (a Role) GetRole(c *gin.Context) {
	id, _ := c.GetQuery("id")
	orderby, _ := c.GetQuery("orderby")
	response := a.RoleUsecase.GetRoles(id, orderby)
	c.JSON(response.Status, response)
}

func (a Role) CreateRole(c *gin.Context) {
	name := c.PostForm("role_name")
	form_name := models.FormName{Name: name}
	response := a.RoleUsecase.CreateRole(form_name)
	c.JSON(response.Status, response)
}
