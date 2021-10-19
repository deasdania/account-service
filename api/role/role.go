package role

import (
	accountusecase "account-metalit/api/account/usecase"
	authusecase "account-metalit/api/auth/usecase"
	"account-metalit/api/models"
	"account-metalit/api/role/usecase"
	"account-metalit/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Role struct {
	RoleUsecase    usecase.IRoleUsecase
	AuthUsecase    authusecase.IAuthUsecase
	AccountUsecase accountusecase.IAccountUsecase
}

func (a Role) Role(r *gin.RouterGroup) {
	r.GET(utilities.GET_ROLE, a.GetRole)        //Hanya untuk Admin
	r.POST(utilities.CREATE_ROLE, a.CreateRole) //Hanya untuk Admin
	r.PUT(utilities.UPDATE_ROLE, a.UpdateRole)  //Hanya untuk Admin
}

// GetRole godoc
// @Summary GetRole Private
// @Description get existing role
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param Body query string false "set 'id' or 'orderby' as Query Params"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/role [get]
func (a Role) GetRole(c *gin.Context) {
	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if errA != nil {
		fmt.Println(errA.Error())
		return
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

// CreateRole godoc
// @Summary CreateRole Private
// @Description create newrole could be access just by user has admin role
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param Body formData models.FormName true "set 'name' to create new role"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/create/role [post]
func (a Role) CreateRole(c *gin.Context) {
	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if errA != nil {
		fmt.Println(errA.Error())
		return
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

// UpdateRole godoc
// @Summary UpdateRole Private
// @Description update existing role name
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param Body formData models.BodyUpdateRole true "set 'role_id' and 'role_name' to update the role"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/update/role [put]
func (a Role) UpdateRole(c *gin.Context) {
	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if errA != nil {
		fmt.Println(errA.Error())
		return
	}
	fmt.Println(metadata)
	isAdmin := a.AccountUsecase.CheckUserIsAdmin(metadata.Email)
	if isAdmin {
		id := c.PostForm("role_id")
		name := c.PostForm("role_name")
		idtrim := strings.Trim(id, " ")
		nametrim := strings.Trim(name, " ")
		intid, _ := strconv.Atoi(idtrim)
		role := models.Role{Id: intid, Name: nametrim}
		response := a.RoleUsecase.UpdateRole(role)
		c.JSON(response.Status, response)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "you are not allowed",
	})
}
