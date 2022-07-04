package entity

import "time"

type users struct {
	id_user    string `gorm:"primaryKey;type:varchar(36);"`
	name       string `json:"name" binding:"required"`
	status     bool
	nomer_HP   string
	email      string `gorm:"unique" json:"email" binding:"required,email"`
	user_Name  string `gorm:"unique" json:"username" binding:"required"`
	password   string `json:"password"`
	address    string
	created_at time.Time `gorm:"autoCreateTime"`
	updated_at time.Time `gorm:"autoCreateTime"`
	deleted_at time.Time `gorm:"index"`
}
