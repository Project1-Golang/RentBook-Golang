package Entity

import "time"

type rent_book struct {
	id_rent_book string    `gorm:"primaryKey;type:varchar(36);"`
	Users        []users   `gorm:"foreignKey:id_users"`
	Books        []books   `gorm:"foreignKey:id_books"`
	return_date  time.Time `gorm:"CreateTime"`
	created_at   time.Time `gorm:"autoCreateTime"`
	updated_at   time.Time `gorm:"autoCreateTime"`
	deleted_at   time.Time `gorm:"index"`
}
