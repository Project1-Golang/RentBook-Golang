package entity

import (
	"log"

	"gorm.io/gorm"
)

type Rent_Book struct {
	Id_rent_book  string `gorm:"primaryKey;type:varchar(36);"`
	Owned_by      string
	Owned_by_book string
	Is_Returned   bool
	// Return_date   time.Time
	// Created_at time.Time `gorm:"autoCreateTime"`
	// Updated_at time.Time `gorm:"autoCreateTime"`
	// Deleted_at time.Time `gorm:"index"`
}

type AksesRentBook struct {
	DB *gorm.DB
}

func (as *AksesRentBook) PinjamBuku(newRent Rent_Book) Rent_Book {

	err := as.DB.Create(&newRent).Error
	if err != nil {
		log.Println(err)
		return Rent_Book{}
	}

	return newRent
}

func (as *AksesRentBook) HitungAllRentBook() int {
	var jumlah int
	as.DB.Raw("SELECT count(id_rent_book) as 'jumlah' FROM rent_books").Scan(&jumlah)
	return jumlah + 1
}
