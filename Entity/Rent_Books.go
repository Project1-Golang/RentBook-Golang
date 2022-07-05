package entity

import "time"

type Rent_Book struct {
	Id_rent_book  string `gorm:"primaryKey;type:varchar(36);"`
	Owned_by      string
	Owned_by_book string
	Return_date   time.Time `gorm:"CreateTime"`
	Created_at    time.Time `gorm:"autoCreateTime"`
	Updated_at    time.Time `gorm:"autoCreateTime"`
	Deleted_at    time.Time `gorm:"index"`
}
