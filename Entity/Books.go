package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Books struct {
	Id_book  string `gorm:"primaryKey;type:varchar(36);"`
	Owned_by string
	// Rent_Book   []Rent_Book `gorm:"foreignKey:owned_by; constraint:OnUpdate:CASCADE, OnDelete:CASCADE; "`
	Title_book  string
	Isbn        string
	Author      string
	Rent_status bool
	Created_at  time.Time `gorm:"autoCreateTime"`
	Updated_at  time.Time `gorm:"autoCreateTime"`
	// Deleted_at  time.Time `gorm:"index"`
}

type AksesBook struct {
	DB *gorm.DB
}

func (as *AksesBook) GetAllData() []Books {
	var daftarBook = []Books{}
	// err := as.DB.Raw("Select * from Books").Scan(&daftarBook)
	err := as.DB.Find(&daftarBook)

	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarBook
}

func (as *AksesBook) TambahBukuBaru(newBook Books) Books {

	err := as.DB.Create(&newBook).Error
	if err != nil {
		log.Println(err)
		return Books{}
	}

	return newBook
}

func (as *AksesBook) GetSpecificBuku(UID int) Books {
	var daftarBook = Books{}
	err := as.DB.First(&daftarBook)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return Books{}
	}

	return daftarBook
}

func (as *AksesBook) HapusBukuSaya(Id_User string, IDBook string) string {
	postExc := as.DB.Where("Id_book = ?", IDBook).Where("owned_by = ?", Id_User).Delete(&Books{})
	if err := postExc.Error; err != nil {
		log.Fatal(err)
		return "Delete My Book Error"
	}
	// berapa data yang berubah (?)
	if aff := postExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang dihapus")
		return "Delete My Book Error"
	}

	return "Anda Telah Berhasil menghapus Buku"

}

func (as *AksesBook) HitungAllBukuAktiv() int {
	var jumlah int
	as.DB.Raw("SELECT count(id_book) as 'jumlah' FROM books").Scan(&jumlah)
	return jumlah + 1
}

func (as *AksesBook) UpdateStatusBook(id string, status bool) string {

	UpdateExc := as.DB.Model(&Books{}).Where("id_book = ?", id).Update("rent_status", status)
	if err := UpdateExc.Error; err != nil {
		log.Fatal(err)
		return "Error"
	}
	if aff := UpdateExc.RowsAffected; aff < 1 {
		return "Error"
	}

	return "Sukses"
}

func (as *AksesBook) Get_Book_belongto_User(ID string) []Books {
	var daftarUserBook = []Books{}
	err := as.DB.Where("owned_by = ?", ID).Find(&daftarUserBook)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarUserBook
}

func (as *AksesBook) GetBookAnotherUser_StatusRentOk(ID string) []Books {
	var daftarUserBook = []Books{}
	err := as.DB.Where("owned_by != ? and rent_status=1", ID).Find(&daftarUserBook)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarUserBook
}

func (as *AksesBook) Get_Book_notbelongto_User(ID string) []Books {
	var daftarUserBook = []Books{}
	err := as.DB.Where("owned_by != ?", ID).Find(&daftarUserBook)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarUserBook
}
func (as *AksesBook) UpdateMyBook(IDBook string, id_user string, NewData Books) string {
	UpdateExc := as.DB.Model(&Books{}).Where("id_book = ? and owned_by = ?", IDBook, id_user).Updates(NewData)
	if err := UpdateExc.Error; err != nil {
		log.Fatal(err)
		return "Error"
	}
	if aff := UpdateExc.RowsAffected; aff < 1 {
		return "Error"
	}
	return "Update Buku Berhasil"
}
