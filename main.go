package main

import (
	"fmt"
	"rentbook-golang/db"
	"rentbook-golang/entity"
	"strconv"
)

func Halaman_login() (string, string) {
	var username, pass string
	fmt.Println("Silahkan Log In")
	fmt.Println("Masukkan Username :")
	fmt.Scanln(&username)
	fmt.Println("Masukkan Password : ")
	fmt.Scanln(&pass)
	return username, pass
}

func main() {
	conn := db.InitDB()
	db.MigrateDB(conn)
	AksesBook := entity.AksesBook{DB: conn}
	AksesUsers := entity.AksesUsers{DB: conn}
	var menu = false
	for !menu {
		var input int = 0
		fmt.Println("===============================")
		fmt.Println("* Selamat Datang Perpustakaan *")
		fmt.Println("===============================")
		fmt.Println("Silahkan Pilih Menu:")
		fmt.Println("1. Register")
		fmt.Println("2. Log In")
		fmt.Println("3. Lihat Daftar Buku")
		fmt.Println("4. Keluar")
		fmt.Print("Masukkan Pilihan menu: ")
		fmt.Scanln(&input)
		// fmt.Println("4. Tambah Data Buku")

		if input == 1 {
			// var adduser entity.AksesUsers
			var newUsers entity.Users
			newUsers.Id_user = "User1"
			newUsers.Status = "1"
			// newUsers.Deleted_at = nil
			fmt.Println("--- Silahkan Isi Data Anda Untuk Registrasi -----")
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
			fmt.Println("Berhasil input User")
		} else if input == 2 {
			UserName, Password := Halaman_login()
			UserAuth := AksesUsers.GetUserName(UserName)
			passAuth := AksesUsers.GetUserPassword(Password)

			if !UserAuth && !passAuth {
				fmt.Println("Username dan Password Tidak tersedia \n Silahkan Register Terlebih dahulu")
			} else if !passAuth || !UserAuth {
				fmt.Println("Username atau Password anda Salah \n Silahkan Periksa Kembali")
			} else {
				fmt.Println("Anda Berhasil Login")
				menu = true
			}
		} else if input == 3 {
			fmt.Println("Daftar Buku Yang Ada")
			for _, val := range AksesBook.GetAllData() {
				fmt.Println(val.Id_book, val.Title_book, val.Author)
			}
		} else {
			fmt.Println("Terimakasih Atas Kunjungannya")
			break
		}

	}
	for menu {
		var pilih int
		fmt.Println("--------------- SELAMAT DATANG ---------------")
		fmt.Println("----- Silahkan Pilih Fitur yang Tersedia -----")
		fmt.Println("1. Lihat Akun Saya")
		fmt.Println("2. Perbarui Akun Saya")
		fmt.Println("3. Hapus Akun Saya")
		fmt.Println("4. Tambah Buku Saya")
		fmt.Println("5. Lihat Daftar Buku Anda")
		fmt.Println("6. Hapus Buku Anda")
		fmt.Println("7. Exit")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)
		fmt.Print("\n")
		switch pilih {
		case 4:
			var newBook entity.Books
			var code string
			JumlahBuku := AksesBook.HitungAllBukuAktiv()
			// fmt.Println(JumlahBuku)
			code = strconv.Itoa(JumlahBuku)

			newBook.Id_book = "Book-" + code

			newBook.Rent_status = true
			var Id_user string
			ID := AksesUsers.GetSpecificUser(Id_user)
			newBook.Owned_by = ID.Id_user

			fmt.Print("Masukkan Judul Buku: ")
			fmt.Scan(&newBook.Title_book)
			fmt.Print("Masukkan Author: ")
			fmt.Scan(&newBook.Author)
			fmt.Print("Masukkan ISBN: ")
			fmt.Scan(&newBook.Isbn)

			AksesBook := entity.AksesBook{DB: conn}
			AksesBook.TambahBukuBaru(newBook)
			fmt.Println("Berhasil Input Data Buku")
		default:
			continue

		case 7:
			fmt.Println("Terimakasih Atas Kunjungannya")
			menu = false
		}
	}
}
