package main

import (
	"bufio"
	"fmt"
	"os"
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
	AksesRent := entity.AksesRentBook{DB: conn}

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
		fmt.Print("Masukkan Pilihan Menu: ")
		fmt.Scanln(&input)
		// fmt.Println("4. Tambah Data Buku")

		if input == 1 {
			// var adduser entity.AksesUsers
			var newUsers entity.Users

			var code string
			JumlahUser := AksesUsers.HitungAllUser()
			code = strconv.Itoa(JumlahUser)
			newUsers.Id_user = "User-0" + code
			newUsers.Status = "1"

			fmt.Println("--- Silahkan Isi Data Anda Untuk Registrasi -----")
			fmt.Print("Masukkan Nama: ")
			in := bufio.NewReader(os.Stdin)
			newUsers.Name, _ = in.ReadString('\n')
			// fmt.Scanln(&newUsers.Name)
			fmt.Print("Masukkan Nomor HP: ")
			fmt.Scanln(&newUsers.Nomer_HP)
			fmt.Print("Masukkan Email: ")
			fmt.Scanln(&newUsers.Email)
			fmt.Print("Masukkan User Name: ")
			fmt.Scanln(&newUsers.User_Name)
			fmt.Print("Masukkan Password: ")
			fmt.Scanln(&newUsers.Password)
			fmt.Print("Masukkan Address: ")
			fmt.Scanln(&newUsers.Address)

			aksesUser := entity.AksesUsers{DB: conn}
			aksesUser.TambahUserBaru(newUsers)
			fmt.Println("Berhasil Input User")
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
		fmt.Println("7. Pinjam Buku")
		fmt.Println("8. Kembalikan Buku")
		fmt.Println("9. Lihat Daftar Yang Tersedia")
		fmt.Println("10. Exit")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)
		fmt.Print("\n")
		switch pilih {
		case 1:
			fmt.Println("----- Info Akun Saya -----")
			for _, val := range AksesUsers.ReadUserInfo() {
				fmt.Println("ID: ", val.Id_user)
				fmt.Println("Nama: ", val.Name)
				fmt.Println("User Name: ", val.User_Name)
				fmt.Println("Nomor HP: ", val.Nomer_HP)
				fmt.Println("Email: ", val.Email)
				fmt.Println("Address: ", val.Address)
			}

		case 2:

		case 3:
			var Id_user int
			fmt.Print("Masukkan ID yang akan dihapus ")
			fmt.Scanln(&Id_user)
			fmt.Println(AksesUsers.HapusUsers(Id_user))

		case 4:
			var newBook entity.Books
			var code string
			JumlahBuku := AksesBook.HitungAllBukuAktiv()
			code = strconv.Itoa(JumlahBuku)

			newBook.Id_book = "Book-0" + code

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

		case 5:

		case 6:

		case 7:
			var newRent entity.Rent_Book
			var code string
			jumlahdata := AksesRent.HitungAllRentBook()
			code = strconv.Itoa(jumlahdata)

			newRent.Id_rent_book = "Pinj-0" + code
			var Id_user string
			ID := AksesUsers.GetSpecificUser(Id_user)
			newRent.Owned_by = ID.Id_user
			// newRent.Return_date = nil
			fmt.Print("Masukkan Id Books: ")
			fmt.Scan(&newRent.Owned_by_book)

			AksesRent.PinjamBuku(newRent)
			fmt.Println("Berhasil Pinjam")
			// AksesBook := entity.AksesBook{DB: conn}
			// AksesBook.TambahBukuBaru(newBook)
			// fmt.Println("Berhasil Input Data Buku")
		case 8:

		case 9:

		case 10:
			fmt.Println("Terimakasih Atas Kunjungannya")
			menu = false

		default:
			continue
		}
	}
}
