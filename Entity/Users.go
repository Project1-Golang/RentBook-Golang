package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type users struct {
	id_user    string `gorm:"primaryKey;type:varchar(36);"`
	name       string
	status     bool
	nomer_HP   string
	email      string `gorm:"unique" binding:"required,email"`
	user_Name  string `gorm:"unique" binding:"required"`
	password   string
	address    string
	created_at time.Time `gorm:"autoCreateTime"`
	updated_at time.Time `gorm:"autoCreateTime"`
	deleted_at time.Time `gorm:"index"`
}

type AksesUsers struct {
	DB *gorm.DB
}

func (as *AksesUsers) GetAllData() []users {
	var daftarUsers = []users{}
	// err := as.DB.Raw("Select * from student").Scan(&daftarStudent)
	err := as.DB.Find(&daftarUsers)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarUsers
}

// func (as *AksesUsers) TambahUserBaru(newUsers users) users {
// 	if newUsers.name == "Jerry" {
// 		newUsers.id_user = uint(1)
// 	}
// 	uid := uuid.New()
// 	newUsers.id_user = uid.String()
// 	err := as.DB.Create(&newUsers).Error
// 	if err != nil {
// 		log.Println(err)
// 		return users{}
// 	}

// 	return newUsers
// }

// func (as *AksesUsers) GetSpecificUser(UID int) users {
// 	var daftarUsers = users{}
// 	daftarUsers.id_user = uint(UID)
// 	// err := as.DB.Raw("Select * from student").Scan(&daftarStudent)
// 	err := as.DB.First(&daftarUsers)
// 	if err.Error != nil {
// 		log.Fatal(err.Statement.SQL.String())
// 		return users{}
// 	}

// 	return daftarUsers
// }

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
