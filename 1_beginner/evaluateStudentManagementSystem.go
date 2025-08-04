package beginner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Login struct {
	Username string
	Password string
}

func tampilanAdmin() {
	fmt.Println("=== SISTEM MANAJEMEN MAHASISWA ===")
	fmt.Println("1. Akademik")
	fmt.Println("2. Utilities")
	fmt.Println("3. Manajemen Tugas")
	fmt.Println("4. Mini Games")
	fmt.Println("5. Database Mahasiswa")
	fmt.Println("6. Hitung Gahu Asisten Dosen")
	fmt.Println("0. Exit")
}

func tampilanMhs() {
	fmt.Println("1. Lihat data siswa")
	fmt.Println("2. Lihat nilai dan ipk")
	fmt.Println("3. Lihat tugas")

}

func loginAdmin() {
	reader := bufio.NewReader(os.Stdin)

	akun := []Login{
		{Username: "Torikal", Password: "220602047"},
	}

	for {
		fmt.Print("Masukkan Username: ")
		u, _ := reader.ReadString('\n')
		u = strings.TrimSpace(u)
		fmt.Print("Masukkan Password: ")
		p, _ := reader.ReadString('\n')
		p = strings.TrimSpace(p)

		berhasil_login := false
		for _, i := range akun {
			if u == i.Username && p == i.Password {
				berhasil_login = true
				break
			}
		}

		if berhasil_login {
			tampilanAdmin()
			return
		} else {
			fmt.Println("Username dan password invalid")
		}
	}
}

func loginMhs() {

	reader := bufio.NewReader(os.Stdin)
	akun := []Login{
		{Username: "Akbar", Password: "220602047"},
	}

	fmt.Print("Masukkan username: ")
	u, _ := reader.ReadString('\n')
	u = strings.TrimSpace(u)
	fmt.Print("Masukkan password: ")
	p, _ := reader.ReadString('\n')
	p = strings.TrimSpace(p)

	berhasil_login := false
	for _, i := range akun {
		if u == i.Username && p == i.Password {
			berhasil_login = true
			break
		}
	}

	if berhasil_login {
		tampilanMhs()
		return
	} else {
		fmt.Println("Username & password invalid")
	}

}

func StudentManagementSystem() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Hai Welcome")
		fmt.Println("1. Login Sebagai Siswa")
		fmt.Println("2. Login Sebagai Admin")
		fmt.Print("Pilih jenis login: ")
		pilih_login, _ := reader.ReadString('\n')
		pilih_login = strings.TrimSpace(pilih_login)
		switch pilih_login {
		case "1":
			loginMhs()
			return
		case "2":
			loginAdmin()
			return
		default:
			fmt.Println("Pilihan Invalid")
		}
	}
}
