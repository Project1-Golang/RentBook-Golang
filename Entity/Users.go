package entity

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	Id_user    string      `gorm:"primaryKey;type:varchar(36);"`
	Books      []Books     `gorm:"foreignKey:owned_by"`
	Rent_Book  []Rent_Book `gorm:"foreignKey:owned_by"`
	Name       string
	Status     bool
	Nomer_HP   string
	Email      string `gorm:"unique"`
	User_Name  string `gorm:"unique"`
	Password   string
	Address    string
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoCreateTime"`
	Deleted_at time.Time `gorm:"index"`
}

type AksesUsers struct {
	DB *gorm.DB
}

func (as *AksesUsers) GetAllData() []Users {
	var daftarUsers = []Users{}
	// err := as.DB.Raw("Select * from student").Scan(&daftarStudent)
	err := as.DB.Find(&daftarUsers)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarUsers
}

func (as *AksesUsers) TambahUserBaru(newUsers Users) Users {
	// if newUsers.name == "Jerry" {
	// 	newUsers.id_user = uint(1)
	// }
	uid := uuid.New()
	newUsers.Id_user = uid.String()
	err := as.DB.Create(&newUsers).Error
	if err != nil {
		log.Println(err)
		return Users{}
	}

	return newUsers
}

func (as *AksesUsers) GetSpecificUser(UID int) Users {
	var daftarUsers = Users{}
	// daftarUsers.Id_user = uint(UID)
	// err := as.DB.Raw("Select * from student").Scan(&daftarStudent)
	err := as.DB.First(&daftarUsers)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return Users{}
	}

	return daftarUsers
}

// func (as *AksesUsers) HapusMurid(IDUsers int) bool {
// 	postExc := as.DB.Where("ID = ?", IDUsers).Delete(&users{})
// 	// ada masalah ga(?)
// 	if err := postExc.Error; err != nil {
// 		log.Fatal(err)
// 		return false
// 	}
// 	// berapa data yang berubah (?)
// 	if aff := postExc.RowsAffected; aff < 1 {
// 		log.Println("Tidak ada data yang dihapus")
// 		return false
// 	}

// 	return true

// }
