package config

import (
	"account-metalit/api/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	"os"
)

func InitDb() *gorm.DB {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("ACCOUNT_DATABASE_USER"),
		os.Getenv("ACCOUNT_DATABASE_PASSWORD"),
		os.Getenv("ACCOUNT_DATABASE_HOST"),
		os.Getenv("ACCOUNT_DATABASE_PORT"),
		os.Getenv("ACCOUNT_DATABASE_NAME"),
	)
	db, err := gorm.Open("mysql", dataSource)

	if err != nil {
		panic(err)
		return nil
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.UserRole{})
	db.AutoMigrate(&models.UserPermission{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.UserRole{})
	db.AutoMigrate(&models.Permission{})
	db.AutoMigrate(&models.ContentType{})
	db.AutoMigrate(&models.RolePermission{})
	db.Model(&models.Permission{}).AddForeignKey("content_type_id", "content_types(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.RolePermission{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.RolePermission{}).AddForeignKey("permission_id", "permissions(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.UserRole{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.UserRole{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.UserPermission{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.UserPermission{}).AddForeignKey("permission_id", "permissions(id)", "RESTRICT", "RESTRICT")

	return db
}
