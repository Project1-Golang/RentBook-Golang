package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Rent_Book struct {
	Id_rent_book string `gorm:"primaryKey;type:varchar(36);"`
	// Owned_by      string
	// Owned_by_book string
	Id_User     string
	Id_book     string
	Is_Returned bool
	Return_date string
	Created_at  time.Time `gorm:"autoCreateTime"`
	Updated_at  time.Time `gorm:"autoCreateTime"`
	// Deleted_at    time.Time `gorm:"index"`
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

func (as *AksesRentBook) KembalikanBuku(IDRENT string, IDBOOK string) string {

	UpdateExc := as.DB.Model(&Rent_Book{}).Where("id_rent_book = ? and id_book=? ", IDRENT, IDBOOK).Update("is_returned", 1)

	if err := UpdateExc.Error; err != nil {
		log.Fatal(err)
		return "error"
	}
	if aff := UpdateExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang dihapus")
		return "error"
	}

	return "berhasil"
}

func (as *AksesRentBook) HitungAllRentBook() int {
	var jumlah int
	as.DB.Raw("SELECT count(id_rent_book) as 'jumlah' FROM rent_books").Scan(&jumlah)
	return jumlah + 1
}

func (as *AksesRentBook) RentByUser(id string) []Rent_Book {
	daftarrent := []Rent_Book{}
	err := as.DB.Where("id_user = ? and is_returned = ?", id, false).Find(&daftarrent)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarrent
}

func (as *AksesRentBook) ValidasiPinjam(idbook string, iduser string) bool {
	validasi := Books{}
	var valid bool
	//idbook salah
	err := as.DB.Where("id_book = ?", idbook).Find(&validasi)
	// err := as.DB.Where("owned_by != ?", iduser).Find(&validasi)
	//SELECT * FROM books  WHERE `id_book`!="ID INPUTAN" AND `owned_by`!="ID USER"
	if err := err.Error; err != nil {
		return valid
	}
	if aff := err.RowsAffected; aff < 1 {
		return !valid
	}
	return valid
}
