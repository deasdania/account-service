package seed

import (
	"account-metalit/api/models"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"

	"time"
)

func Load(db *gorm.DB) {
	fmt.Println("open seed")
	err := db.Debug().AutoMigrate(&models.User{}, &models.UserRole{}, &models.UserPermission{}, &models.Role{}, &models.UserRole{}, &models.Permission{}, &models.ContentType{}, &models.RolePermission{}).Error
	if err != nil {
		fmt.Println("cannot migrate table: %v", err)
	} else {
		db.Model(&models.Permission{}).AddForeignKey("content_type_id", "content_types(id)", "RESTRICT", "RESTRICT")
		db.Model(&models.RolePermission{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
		db.Model(&models.RolePermission{}).AddForeignKey("permission_id", "permissions(id)", "RESTRICT", "RESTRICT")
		db.Model(&models.UserRole{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Model(&models.UserRole{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
		db.Model(&models.UserPermission{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Model(&models.UserPermission{}).AddForeignKey("permission_id", "permissions(id)", "RESTRICT", "RESTRICT")

		var roles = []models.Role{
			models.Role{
				Id:          1,
				Name:        "admin",
				CreatedDate: time.Now(),
				UpdateDate:  time.Now(),
			},
			models.Role{
				Id:          2,
				Name:        "member",
				CreatedDate: time.Now(),
				UpdateDate:  time.Now(),
			},
		}

		bytes, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_ACCOUNT_PASSWORD")), 14)

		var user = models.User{
			Id:          1,
			Uuid:        uuid.New().String(),
			Name:        os.Getenv("ADMIN_ACCOUNT_USERNAME"),
			Email:       os.Getenv("ADMIN_ACCOUNT_EMAIL"),
			Password:    string(bytes),
			CreatedDate: time.Now(),
			UpdateDate:  time.Now(),
		}

		var userRole = models.UserRole{
			Id:     1,
			UserId: 1,
			RoleId: 1,
		}
		err := db.Model(&models.User{}).Create(&user).Error
		if err != nil {
			fmt.Printf("[IGNORE THIS]cannot seed User table: %v \n", err)
		}
		for i, _ := range roles {
			err = db.Model(&models.Role{}).Create(&roles[i]).Error
			if err != nil {
				fmt.Printf("[IGNORE THIS]cannot seed Roles table: %v \n", err)
			}
		}
		err = db.Model(&models.UserRole{}).Create(&userRole).Error
		if err != nil {
			fmt.Printf("[IGNORE THIS]cannot seed UserRole table: %v \n", err)
		}
	}
}
