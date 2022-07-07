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
		fmt.Println("===============================\t")
		fmt.Println("* Selamat Datang Perpustakaan *\t")
		fmt.Println("===============================\t")
		fmt.Println("Silahkan Pilih Menu:")
		fmt.Println("1. Register")
		fmt.Println("2. Log In")
		fmt.Println("3. Lihat Daftar Buku")
		fmt.Println("4. Keluar")
		fmt.Print("Masukkan Pilihan Menu: ")
		fmt.Scanln(&input)

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
			fmt.Print("Masukkan Nomor HP: ")
			fmt.Scanln(&newUsers.Nomer_HP)
			fmt.Print("Masukkan Email: ")
			fmt.Scanln(&newUsers.Email)
			fmt.Print("Masukkan User Name: ")
			fmt.Scanln(&newUsers.User_Name)
			fmt.Print("Masukkan Password: ")
			fmt.Scanln(&newUsers.Password)
			fmt.Print("Masukkan Address: ")
			en := bufio.NewReader(os.Stdin)
			newUsers.Address, _ = en.ReadString('\n')

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
		fmt.Println("6. Perbarui Buku Anda")
		fmt.Println("7. Hapus Buku Anda")
		fmt.Println("8. Pinjam Buku")
		fmt.Println("9. Kembalikan Buku")
		fmt.Println("10. Lihat Daftar Buku Yang Tersedia")
		fmt.Println("11. Exit")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)
		fmt.Print("\n")
		switch pilih {
		case 1: //liat Akun
			// fmt.Println("----- Info Akun Saya -----")
			// for _, val := range AksesUsers.ReadUserInfo() {
			// 	fmt.Println("ID: ", val.Id_user)
			// 	fmt.Println("Nama: ", val.Name)
			// 	fmt.Println("Nomor HP: ", val.Nomer_HP)
			// 	fmt.Println("User Name: ", val.User_Name)
			// 	fmt.Println("Address: ", val.Address)
			// 	fmt.Println("Email: ", val.Email)
			// }

		case 2: //update user

			var Id_user string //ambil user yang Aktif

			UserAktif := AksesUsers.GetSpecificUser(Id_user)

			var nama string
			fmt.Print("Masukkan Nama :")
			fmt.Scan(&nama)
			fmt.Println(nama)
			AksesUsers.EditUser(UserAktif.Id_user, nama)
			// //input data
			// var UserEdit entity.Users
			// UserEdit.Id_user = UserAktif.Id_user
			// fmt.Println("--- Silahkan Perbaharui Data Anda -----")
			// fmt.Print("Masukkan Nama: ")
			// in := bufio.NewReader(os.Stdin)
			// UserEdit.Name, _ = in.ReadString('\n')
			// fmt.Print("Masukkan Nomor HP: ")
			// fmt.Scanln(&UserEdit.Nomer_HP)
			// fmt.Print("Masukkan Email: ")
			// fmt.Scanln(&UserEdit.Email)
			// fmt.Print("Masukkan User Name: ")
			// fmt.Scanln(&UserEdit.User_Name)
			// fmt.Print("Masukkan Password: ")
			// fmt.Scanln(&UserEdit.Password)
			// fmt.Print("Masukkan Address: ")
			// en := bufio.NewReader(os.Stdin)
			// UserEdit.Address, _ = en.ReadString('\n')

			// aksesUser := entity.AksesUsers{DB: conn}
			// aksesUser.UpdateUser(UserAktif.Id_user, UserEdit.Name, UserEdit.Nomer_HP, UserEdit.Email, UserEdit.User_Name, UserEdit.Password, UserEdit.Address)
			// fmt.Println("Berhasil Update")

			// // fmt.Scanln(&UpdateNama)
			// res := AksesUsers.UpdateUser(Id_user, "Tom")
			// if res {
			// 	fmt.Println("Data Berhasil Diperbarui")
			// } else {
			// 	fmt.Println("Data Gagal Diperbarui")
			// }

		case 3: //Hapus Akun
			var option int
			var Id_user string //ambil user yang Aktif
			UserAktif := AksesUsers.GetSpecificUser(Id_user)
			// fmt.Println(UserAktif.Id_user)
			fmt.Println("Apakah Anda Yakin Akan Menghapus Akun Anda?")
			fmt.Println("1. Ya")
			fmt.Println("2. Tidak")
			fmt.Scan(&option)
			if option == 1 {
				fmt.Println("Hapus Akun")
				fmt.Println(AksesUsers.HapusUsers(UserAktif.Id_user))
				fmt.Println("....................")
				fmt.Println("Akun Anda Sudah dihapus")
				fmt.Println("Terimakasih Atas Kunjungannya")
				menu = false
			} else {
				menu = true
			}

		case 4: //Tambah Buku Saya
			var newBook entity.Books
			var code string
			JumlahBuku := AksesBook.HitungAllBukuAktiv()
			code = strconv.Itoa(JumlahBuku)

			newBook.Id_book = "Book-0" + code
			newBook.Rent_status = true

			var Id_user string //ambil user yang Aktif
			ID := AksesUsers.GetSpecificUser(Id_user)

			newBook.Owned_by = ID.Id_user

			fmt.Println("Masukkan Judul Buku: ")
			ba := bufio.NewReader(os.Stdin)
			newBook.Title_book, _ = ba.ReadString('\t')

			fmt.Println("Masukkan Author: ")
			bb := bufio.NewReader(os.Stdin)
			newBook.Author, _ = bb.ReadString('\t')

			fmt.Println("Masukkan ISBN: ")
			bc := bufio.NewReader(os.Stdin)
			newBook.Isbn, _ = bc.ReadString('\t')

			AksesBook := entity.AksesBook{DB: conn}
			AksesBook.TambahBukuBaru(newBook)
			fmt.Println("Berhasil Input Data Buku")

		case 5:

		case 6:

		case 7: //hapus buku saya
			var IDBook string
			fmt.Println("Masukkan ID Buku Anda : ")
			fmt.Scan(&IDBook)

			var Id_user string //ambil user yang Aktif

			UserAktif := AksesUsers.GetSpecificUser(Id_user)

			fmt.Println(AksesBook.HapusBuku(UserAktif.Id_user, IDBook))

		case 8:
			var newRent entity.Rent_Book
			var code string
			jumlahdata := AksesRent.HitungAllRentBook()
			code = strconv.Itoa(jumlahdata)

			newRent.Id_rent_book = "Pinj-0" + code
			var id string
			ID := AksesUsers.GetSpecificUser(id)
			newRent.Owned_by = ID.Id_user
			newRent.Is_Returned = false
			fmt.Print("Masukkan Id Books: ")
			fmt.Scan(&newRent.Owned_by_book)

			AksesRent.PinjamBuku(newRent)
			fmt.Println("Berhasil Pinjam")

		case 9:

		case 10:

		case 11:
			fmt.Println("Terimakasih Atas Kunjungannya")
			menu = false
		}
	}
}
