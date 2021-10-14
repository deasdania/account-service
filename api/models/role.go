package models

import "time"

// Role as Group while on django
type Role struct {
	Id          uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null`
	Name        string    `json:"name" gorm:"not null;size:20"`
	CreatedDate time.Time `json:"created_date" gorm:"not null;default:CURRENT_TIMESTAMP;"`
	UpdateDate  time.Time `json:"update_date" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
}

type RolePermission struct {
	Id           uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null`
	RoleID       uint64 `json:"role_id"`
	PermissionID uint64 `json:"permission_id"`
}

type FormName struct {
	Name string
}
