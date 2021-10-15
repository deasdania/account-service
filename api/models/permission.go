package models

import "time"

// Users
type Permission struct {
	Id            int       `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null`
	Name          string    `json:"name" gorm:"not null;size:20"`
	ContentTypeId int       `json:"content_type_id" gorm:"not null`
	CodeName      string    `json:"code_name" gorm:"not null;size:100"`
	CreatedDate   time.Time `json:"created_date" gorm:"not null;default:CURRENT_TIMESTAMP;"`
	UpdateDate    time.Time `json:"update_date" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
}

type ContentType struct {
	Id        int `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null`
	TableName int `json:"table_name" gorm:"not null`
}
