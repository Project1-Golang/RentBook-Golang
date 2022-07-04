package entity

import "time"

type users struct {
	id_user    string `gorm:"primaryKey;type:varchar(36);"`
	name       string
	status     bool
	nomer_HP   string
	email      string
	user_Name  string
	password   string
	address    string
	created_at time.Time `gorm:"autoCreateTime"`
	updated_at time.Time `gorm:"autoCreateTime"`
	deleted_at time.Time `gorm:"index"`
}
