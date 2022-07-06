package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id_user    string      `gorm:"primaryKey;type:varchar(36);"`
	Books      []Books     `gorm:"foreignKey:owned_by"`
	Rent_Book  []Rent_Book `gorm:"foreignKey:owned_by"`
	Name       string
	Status     string
	Nomer_HP   string
	Email      string `gorm:"unique"`
	User_Name  string `gorm:"unique"`
	Password   string
	Address    string
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoCreateTime"`
	// Deleted_at time.Time
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
	// uid := uuid.New()
	// newUsers.Id_user = uid.String()
	err := as.DB.Create(&newUsers).Error
	if err != nil {
		log.Println(err)
		return Users{}
	}

	return newUsers
}

func (as *AksesUsers) GetSpecificUser(UID string) Users {
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

func (as *AksesUsers) GetUserEmail(Email string) bool {
	getEmail := as.DB.Where("email = ?", Email).Find(&Users{})
	if err := getEmail.Error; err != nil {
		return false
	}
	if aff := getEmail.RowsAffected; aff < 1 {
		return false
	}
	return true
}

func (as *AksesUsers) GetUserName(UserName string) bool {
	getUserName := as.DB.Where("user_name = ?", UserName).Find(&Users{})
	if err := getUserName.Error; err != nil {
		return false
	}
	if aff := getUserName.RowsAffected; aff < 1 {
		return false
	}
	return true
}

func (as *AksesUsers) GetUserPassword(Password string) bool {
	getPassword := as.DB.Where("password = ?", Password).Find(&Users{})
	if err := getPassword.Error; err != nil {
		return false
	}
	if aff := getPassword.RowsAffected; aff < 1 {
		return false
	}
	return true
}

func (as *AksesUsers) HapusUsers(Id_user int) bool {
	postExc := as.DB.Where("ID = ?", Id_user).Delete(&Users{})
	// ada masalah ga(?)
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

func (as *AksesUsers) HitungAllUser() int {
	var jumlah int
	as.DB.Raw("SELECT count(id_user) as 'jumlah' FROM users").Scan(&jumlah)
	return jumlah + 1
}
