package main

import (
	"fmt"
	"rentbook-golang/db"
	"rentbook-golang/entity"
)

// "go-db/entity/users"

func main() {
	conn := db.InitDB()
	db.MigrateDB(conn)
	AksesBook := entity.AksesBook{DB: conn}
	// AksesUsers := entity.AksesUsers{DB: conn}
	var input int = 0
	for input != 99 {
		fmt.Println("\tSistem Peminjaman Buku")
		fmt.Println("1. Register")
		fmt.Println("2. Log In")
		fmt.Println("3. Lihat Daftar Buku")
		fmt.Println("99. Keluar")
		fmt.Print("Masukkan Pilihan menu: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			// var adduser entity.AksesUsers
			var newUsers entity.Users
			newUsers.Id_user = "User1"
			newUsers.Status = "1"
			// newUsers.Deleted_at = nil

			fmt.Print("Masukkan nama: ")
			fmt.Scanln(&newUsers.Name)
			fmt.Print("Masukkan nomorhp: ")
			fmt.Scanln(&newUsers.Nomer_HP)
			fmt.Print("Masukkan email: ")
			fmt.Scanln(&newUsers.Email)
			fmt.Print("Masukkan user_Name: ")
			fmt.Scanln(&newUsers.User_Name)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&newUsers.Password)
			fmt.Print("Masukkan address: ")
			fmt.Scanln(&newUsers.Address)

			aksesUser := entity.AksesUsers{DB: conn}
			aksesUser.TambahUserBaru(newUsers)
			fmt.Println("Berhasl input User")
		case 2:
			fmt.Println("Menu Log In")
			fmt.Print("Masukkan user_Name: ")
			// fmt.Scanln(&newUsers.user_Name)
			fmt.Print("Masukkan password: ")
			// fmt.Scanln(&newUsers.password)

		case 3:
			fmt.Println("Daftar Buku Yang Ada")
			for _, val := range AksesBook.GetAllData() {
				fmt.Println(val.Id_book, val.Title_book, val.Author)
			}
		default:
			continue
		}
	}
	fmt.Println("Terima kasih sudah mencoba program saya")
}
