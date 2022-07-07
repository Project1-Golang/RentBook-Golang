package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id_user    string      `gorm:"primaryKey;type:varchar(36);"`
	Books      []Books     `gorm:"foreignKey:owned_by; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Rent_Book  []Rent_Book `gorm:"foreignKey:owned_by; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
	err := as.DB.Find(&daftarUsers)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarUsers
}

func (as *AksesUsers) TambahUserBaru(newUsers Users) Users {
	err := as.DB.Create(&newUsers).Error
	if err != nil {
		log.Println(err)
		return Users{}
	}

	return newUsers
}

func (as *AksesUsers) GetSpecificUser(User_Name, Password string) Users { // Edit Mas Jerry
	var daftarUsers = Users{}
	err := as.DB.Where("user_name = ? and password = ?", User_Name, Password).First(&daftarUsers)
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

func (as *AksesUsers) HapusUsers(id string) string {
	postExc := as.DB.Where("Id_user = ?", id).Delete(&Users{})
	if err := postExc.Error; err != nil {
		log.Fatal(err)
		return "Gagal Menghapus User"
	}
	if aff := postExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang dihapus")
		return "Gagal Menghapus User"
	}

	return "Hapus User Berhasil !"

}

func (as *AksesUsers) ReadUserInfo(Id_user string) Users {
	var daftarUsers = Users{}
	// err := as.DB.Where("Id_user", Id_user).First(&daftarUsers)
	// err := as.DB.Find(&daftarUsers)
	err := as.DB.Select("Id_user", "Name", "Nomer_HP", "User_Name", "Address", "Email").Where("Id_user = ?", Id_user).Limit(1).Find(&daftarUsers)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		// return nil
	}

	return daftarUsers
}

func (as *AksesUsers) HitungAllUser() int {
	var jumlah int
	as.DB.Raw("SELECT count(id_user) as 'jumlah' FROM users").Scan(&jumlah)
	return jumlah + 1
}

// func (as *AksesUsers) UpdateUser(id string, UpdateNama string) bool {
// 	UpdateExc := as.DB.Model(&Users{}).Where("Id_user = ?", id).Update("Name", UpdateNama)
// 	if err := UpdateExc.Error; err != nil {
// 		log.Fatal(err)
// 		return false
// 	}
// 	if aff := UpdateExc.RowsAffected; aff < 1 {
// 		log.Println("Tidak ada data yang dihapus")
// 		return false
// 	}

// 	return true

// }

func (as *AksesUsers) UpdateUser(ID string, nama string, nohp string, email string, username string, pass string, address string) string {

	if err := as.DB.Where(Users{Id_user: ID}).
		Assign(Users{Name: nama, Nomer_HP: nohp, Email: email, User_Name: username, Password: pass}).FirstOrCreate(&Users{}).Error; err != nil {
		return "success"
	}

	return "success"
}

func (as *AksesUsers) EditUser(id string, NewData Users) string {

	UpdateExc := as.DB.Model(&Users{}).Where("id_user = ?", id).Updates(NewData)
	if err := UpdateExc.Error; err != nil {
		log.Fatal(err)
		return "Error"
	}
	if aff := UpdateExc.RowsAffected; aff < 1 {
		return "Error"
	}

	return "Sukses"
}

func (as *AksesUsers) GetSingleUser(id string) Users {
	var infouser = Users{}
	err := as.DB.Where("id_user = ?", id).First(&infouser)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return Users{}
	}
	return infouser
}

// func (as *AksesUsers) GetSpecificUser(User_Name, Password string) Users { // Edit Mas Jerry
// 	var daftarUsers = Users{}
// 	err := as.DB.Where("user_name = ? and password = ?", User_Name, Password).First(&daftarUsers)
// 	if err.Error != nil {
// 		log.Fatal(err.Statement.SQL.String())
// 		return Users{}
// 	}

// 	return daftarUsers
// }
