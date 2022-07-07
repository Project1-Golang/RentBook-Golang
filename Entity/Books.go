package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Books struct {
	Id_book     string `gorm:"primaryKey;type:varchar(36);"`
	Owned_by    string
	Rent_Book   []Rent_Book `gorm:"foreignKey:owned_by; constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
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
	// daftarBook.Id_book = uint(UID)
	// err := as.DB.Raw("Select * from student").Scan(&daftarStudent)
	err := as.DB.First(&daftarBook)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return Books{}
	}

	return daftarBook
}

func (as *AksesBook) HapusBuku(Id_User string, IDBook string) string {
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

//// Get updated records count with `RowsAffected`
//result := db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
// UPDATE users SET name='hello', age=18 WHERE role = 'admin';

// result.RowsAffected // returns updated records count
// result.Error

// func (as *AksesBook) UpdateBook(Id_book string) Books {
// 	var daftarBook = Books{}
// 	// err := as.DB.Where("Id_user", Id_user).First(&daftarUsers)
// 	// err := as.DB.Find(&daftarUsers)
// 	err := as.DB.Select("Id_book", "Title_book", "Isbn", "Author").Where("Id_book = ?", Id_book).Limit(1).Find(&daftarBook)
// 	if err.Error != nil {
// 		log.Fatal(err.Statement.SQL.String())
// 		// return nil
// 	}
// }

func (as *AksesUsers) UpdateBook(id string, nama string) string {

	UpdateExc := as.DB.Model(&Users{}).Where("Id_book = ?", id).Update("name", nama)
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

// func (as *AksesUsers) ReadUserInfo(Id_user string) Users {
// 	var daftarUsers = Users{}
// 	// err := as.DB.Where("Id_user", Id_user).First(&daftarUsers)
// 	// err := as.DB.Find(&daftarUsers)
// 	err := as.DB.Select("Id_user", "Name", "Nomer_HP", "User_Name", "Address", "Email").Where("Id_user = ?", Id_user).Limit(1).Find(&daftarUsers)
// 	if err.Error != nil {
// 		log.Fatal(err.Statement.SQL.String())
// 		// return nil
// 	}

// 	return daftarUsers
// }
