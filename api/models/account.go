package models

import "time"

// Users
type Users struct {
	Id          uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null`
	Uuid        string    `json:"uuid" gorm:"unique;not null"`
	Name        string    `json:"name" gorm:"not null;size:150"`
	Email       string    `json:"email" gorm:"unique;not null;size:75"`
	Password    string    `json:"password" gorm:"not null;size:70"`
	CreatedDate time.Time `json:"created_date" gorm:"not null;default:CURRENT_TIMESTAMP;"`
	UpdateDate  time.Time `json:"update_date" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
}

// func (u *Users) BeforeCreate() (err error) {
// 	u.CreatedDate = time.Now()
// 	u.UpdateDate = time.Now()
// 	return
// }

type FormRegister struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type FormChangePassword struct {
	Email           string `json:"email"`
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
