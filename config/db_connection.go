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
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Roles{})
	db.AutoMigrate(&models.UserRole{})
	db.AutoMigrate(&models.Permissions{})
	db.AutoMigrate(&models.ContentTypes{})
	db.AutoMigrate(&models.RolePermission{})
	db.Model(&models.Permissions{}).AddForeignKey("content_type_id", "content_types(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.RolePermission{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.RolePermission{}).AddForeignKey("permission_id", "permissions(id)", "RESTRICT", "RESTRICT")
	// db.Model(&models.UserRole{}).AddForeignKey("role_id,permission_id", "roles(id), permissions(id)", "RESTRICT", "RESTRICT")
	// db.Model(&models.UserPermission{}).AddForeignKey("user_id,permission_id", "users(id), permissions(id)", "RESTRICT", "RESTRICT")

	return db
}
