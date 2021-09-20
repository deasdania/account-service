package models

import "time"

// Users
type Users struct {
	Id          uint64    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key;not null`
	Uuid        string    `json:"uuid" sql:"unique"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	CreatedDate time.Time `json:"-"`
	UpdateDate  time.Time `json:"-"`
}

func (u *Users) BeforeCreate() (err error) {
	u.CreatedDate = time.Now()
	u.UpdateDate = time.Now()
	return
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}

type FormRegister struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
