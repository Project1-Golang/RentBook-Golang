package entity

import "time"

type books struct {
	id_book     string  `gorm:"primaryKey;type:varchar(36);"`
	users       []users `gorm:"foreignKey:id_users"`
	title_book  string
	isbn        string
	author      string
	rent_status bool
	created_at  time.Time `gorm:"autoCreateTime"`
	updated_at  time.Time `gorm:"autoCreateTime"`
	deleted_at  time.Time `gorm:"index"`
}
