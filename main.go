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
	var UserAktif entity.Users

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
				UserAktif = AksesUsers.GetSpecificUser(UserName, Password) // Edit Mas Jerry
				menu = true
			}
		} else if input == 3 {
			fmt.Println("Daftar Buku Yang Ada")
			var no int
			fmt.Println("Berikut adalah daftar Buku yang tersedia")
			for _, val := range AksesBook.GetAllData() {
				fmt.Println("****************************************")
				no++
				fmt.Println("No :", no)
				fmt.Println("ID :", val.Id_book)
				fmt.Println("Judul :", val.Title_book)
				fmt.Println("Penulis :", val.Author)
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
			var IDUSER = UserAktif.Id_user
			userinfo := AksesUsers.GetSingleUser(IDUSER)

			fmt.Println("*******************************")
			fmt.Println("----- Info Akun Saya -----")
			fmt.Println("*******************************")
			fmt.Println("Nama : ", userinfo.Name)
			fmt.Println("Nomor Hp : ", userinfo.Nomer_HP)
			fmt.Println("User Name : ", userinfo.User_Name)
			fmt.Println("Email : ", userinfo.Email)
			fmt.Println("Alamat : ", userinfo.Address)
		case 2: //update user
			var IDUSER = UserAktif.Id_user

			newUpdate := entity.Users{}
			// newUpdate.Id_user = IDUSER
			fmt.Println("--- Silahkan Perbarui Data Anda -----")
			fmt.Print("Masukkan Nama: ")
			Name := bufio.NewReader(os.Stdin)
			newUpdate.Name, _ = Name.ReadString('\t')
			fmt.Print("Masukkan Nomor HP: ")
			hp := bufio.NewReader(os.Stdin)
			newUpdate.Nomer_HP, _ = hp.ReadString('\t')
			fmt.Print("Masukkan Email: ")
			em := bufio.NewReader(os.Stdin)
			newUpdate.Email, _ = em.ReadString('\t')
			fmt.Print("Masukkan User Name: ")
			un := bufio.NewReader(os.Stdin)
			newUpdate.Email, _ = un.ReadString('\t')
			fmt.Print("Masukkan Password: ")
			pas := bufio.NewReader(os.Stdin)
			newUpdate.Email, _ = pas.ReadString('\t')
			fmt.Print("Masukkan Address: ")
			en := bufio.NewReader(os.Stdin)
			newUpdate.Address, _ = en.ReadString('\t')

			AksesUsers.EditUser(IDUSER, newUpdate)
			fmt.Println("Update Data Berhasil")
		case 3: //Hapus Akun
			var option int
			var IDUSER = UserAktif.Id_user
			fmt.Println("Apakah Anda Yakin Akan Menghapus Akun Anda?")
			fmt.Println("1. Ya")
			fmt.Println("2. Tidak")
			fmt.Scan(&option)
			if option == 1 {
				fmt.Println("Hapus Akun")
				fmt.Println(AksesUsers.HapusUsers(IDUSER))
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

			var IDUSER = UserAktif.Id_user

			newBook.Owned_by = IDUSER
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

		case 5: //daftar buku Saya
			var IDUSER = UserAktif.Id_user

			fmt.Println("*******************************")
			fmt.Println("----- BUKU SAYA -----")
			fmt.Println("*******************************")

			bukusaya := AksesBook.Get_Book_belongto_User(IDUSER)
			var no int
			for _, val := range bukusaya {
				no++
				fmt.Println(no, "Judul Buku :", val.Title_book)
			}

		case 6: //update book user
			var IDUSER = UserAktif.Id_user
			var IDBOOK string

			fmt.Println("*******************************")
			fmt.Println("----- BUKU ANDA -----")
			fmt.Println("*******************************")

			bukusaya := AksesBook.Get_Book_belongto_User(IDUSER)
			var no int
			for _, val := range bukusaya {
				no++
				fmt.Println(no, "Judul Buku :", val.Title_book, "ID BUKU :", val.Id_book)
			}
			fmt.Println("Masukkan ID BUKU ANDA: ")
			fmt.Scan(&IDBOOK)
			if IDBOOK != "" {
				var UpdateBook entity.Books
				fmt.Println("Masukkan Judul Buku: ")
				ba := bufio.NewReader(os.Stdin)
				UpdateBook.Title_book, _ = ba.ReadString('\t')

				fmt.Println("Masukkan Author: ")
				bb := bufio.NewReader(os.Stdin)
				UpdateBook.Author, _ = bb.ReadString('\t')

				fmt.Println("Masukkan ISBN: ")
				bc := bufio.NewReader(os.Stdin)
				UpdateBook.Isbn, _ = bc.ReadString('\t')

				AksesBook := entity.AksesBook{DB: conn}
				AksesBook.UpdateMyBook(IDBOOK, IDUSER, UpdateBook)

				fmt.Println("Buku Berhasil di Update")
			} else {
				fmt.Println("Masukkan ID BUKU ANDA")
			}

		case 7: //hapus buku saya
			var IDUSER = UserAktif.Id_user

			fmt.Println("*******************************")
			fmt.Println("----- BUKU ANDA -----")
			fmt.Println("*******************************")

			bukusaya := AksesBook.Get_Book_belongto_User(IDUSER)
			var no int
			for _, val := range bukusaya {
				no++
				fmt.Println(no, "Judul Buku :", val.Title_book, "ID BUKU :", val.Id_book)
			}
			var IDBook string
			fmt.Println("Masukkan ID Buku yang akan Anda Hapus : ")
			fmt.Scan(&IDBook)
			ID := UserAktif.Id_user
			fmt.Println(AksesBook.HapusBukuSaya(ID, IDBook))

		case 8: //Pinjam Buku
			//tampilkan buku bukan milik saya
			fmt.Println("*******************************")
			fmt.Println("--- DAFTAR BUKU YANG BISA ANDA PINJAM ---")
			fmt.Println("*******************************")
			ID := UserAktif.Id_user
			daftarBukuRent := AksesBook.GetBookAnotherUser_StatusRentOk(ID)
			var num int
			for _, val := range daftarBukuRent {
				if val.Rent_status {
					num++
					fmt.Println(num, "Judul Buku :", val.Title_book, "IDBUKU :", val.Id_book)
				}
			}
			//proses peminjaman ke tabel Rent
			var newRent entity.Rent_Book
			var code string
			jumlahdata := AksesRent.HitungAllRentBook()
			code = strconv.Itoa(jumlahdata)

			newRent.Id_rent_book = "Pinj-0" + code
			// var id string
			newRent.Id_User = ID

			newRent.Is_Returned = false
			newRent.Return_date = ""
			fmt.Print("Masukkan Id Books: ")
			fmt.Scan(&newRent.Id_book)

			AksesRent.PinjamBuku(newRent)

			//ubah status buku di tabel buku

			AksesBook.UpdateStatusBook(newRent.Id_book, false)

			fmt.Println("Berhasil Pinjam")

		case 9: //Kembalikan Buku

		case 10: //Lihat Daftar Buku Yang Tersedia
			fmt.Println("*******************************")
			fmt.Println("--- DAFTAR SEMUA BUKU ---")
			fmt.Println("*******************************")

			var no int
			for _, val := range AksesBook.GetAllData() {
				if val.Rent_status {
					no++
					fmt.Println(no, "Judul Buku :", val.Title_book, "Penulis :", val.Author)
				}

			}

			fmt.Println("*******************************")
			fmt.Println("--- DAFTAR BUKU YANG BISA ANDA PINJAM ---")
			fmt.Println("*******************************")
			ID := UserAktif.Id_user
			daftarBukuRent := AksesBook.Get_Book_notbelongto_User(ID)
			var num int
			for _, val := range daftarBukuRent {
				if val.Rent_status {
					num++
					fmt.Println(num, "Judul Buku :", val.Title_book)
				}
			}
		case 11: //Exit
			fmt.Println("Terimakasih Atas Kunjungannya")
			UserAktif = entity.Users{} //Session deleted
			menu = false
		}
	}
}
