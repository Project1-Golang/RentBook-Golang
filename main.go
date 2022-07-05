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
			var newUsers entity.AksesUsers
			// newUsers.Id_user = 1
			fmt.Print("Masukkan nama: ")
			fmt.Scanln(&newUsers)
			fmt.Print("Masukkan nomorhp: ")
			fmt.Scanln(&newUsers)
			fmt.Print("Masukkan email: ")
			fmt.Scanln(&newUsers)
			fmt.Print("Masukkan user_Name: ")
			fmt.Scanln(&newUsers)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&newUsers)
			fmt.Print("Masukkan address: ")
			fmt.Scanln(&newUsers)

			fmt.Println(newUsers)
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
