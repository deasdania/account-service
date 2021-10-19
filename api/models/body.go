package models

type BodyLoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type BodyCreateAccount struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Name            string `json:"name" binding:"required"`
}

type BodyChangePasswordAccount struct {
	OldPassword     string `json:"old_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type BodyGetRole struct {
	Id      string `json:"id"`
	Orderby string `json:"orderby"`
}

type BodyUpdateRole struct {
	RoleId   string `json:"role_id" binding:"required"`
	RoleName string `json:"role_name" binding:"required"`
}
