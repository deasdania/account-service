package models

import "time"

// Users
type Users struct {
	Id          int64     `json:"id"`
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
