package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Books struct {
	Id_book     string `gorm:"primaryKey;type:varchar(36);"`
	Owned_by    string
	Rent_Book   []Rent_Book `gorm:"foreignKey:Owned_by_book"`
	Title_book  string
	Isbn        string
	Author      string
	Rent_status bool
	Created_at  time.Time `gorm:"autoCreateTime"`
	Apdated_at  time.Time `gorm:"autoCreateTime"`
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
	// if newBook.name == "Harry Potter" {
	// 	newBook.id = uint(1)
	// }
	// uid := uuid.New()
	// books.id = uid.String()
	err := as.DB.Create(&newBook).Error
	if err != nil {
		log.Println(err)
		return Books{}
	}

	return newBook
}

func (as *AksesBook) GetSpecificBuku(UID int) Books {
	var daftarBook = Books{}
	// daftarBook.Id_book = uint(UID)
	// err := as.DB.Raw("Select * from student").Scan(&daftarStudent)
	err := as.DB.First(&daftarBook)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return Books{}
	}

	return daftarBook
}

func (as *AksesBook) HapusBuku(IDBook int) bool {
	postExc := as.DB.Where("ID = ?", IDBook).Delete(&Books{})
	if err := postExc.Error; err != nil {
		log.Fatal(err)
		return false
	}
	// berapa data yang berubah (?)
	if aff := postExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang dihapus")
		return false
	}

	return true

}

func (as *AksesBook) HitungAllBukuAktiv() int {
	var jumlah int
	as.DB.Raw("SELECT count(id_book) as 'jumlah' FROM books").Scan(&jumlah)
	return jumlah + 1
}

//// Get updated records count with `RowsAffected`
//result := db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
// UPDATE users SET name='hello', age=18 WHERE role = 'admin';

// result.RowsAffected // returns updated records count
// result.Error
