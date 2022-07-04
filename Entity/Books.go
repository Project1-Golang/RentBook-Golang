package entity

type books struct {
	id_book     string `gorm:"primaryKey;type:varchar(36);"`
	title       string
	isbn        string
	author      string
	rent_status bool
}
