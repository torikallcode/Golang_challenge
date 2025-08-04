package beginner

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// ✅ Challenge 1: Kalkulator Sederhana (CLI)
// Materi yang dilatih: Variabel, input/output, operator, if-else

// Deskripsi:
// Input dua angka dari user
// Pilih operasi: tambah, kurang, kali, bagi
// Tampilkan hasilnya
// 💡 Tingkat lanjut: Tambahkan loop agar bisa digunakan berkali-kali.

// func Calculator() {

// 	// if operator == "*" {
// 	// 	fmt.Println(angka_pertama * angka_kedua)
// 	// } else if operator == "/" {
// 	// 	fmt.Println(angka_pertama / angka_kedua)
// 	// } else if operator == "+" {
// 	// 	fmt.Println(angka_pertama + angka_kedua)
// 	// } else if operator == "-" {
// 	// 	fmt.Println(angka_pertama - angka_kedua)
// 	// }

// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		var pilihan int
// 		fmt.Println("=== SELAMAT DATANG DI CALCULATOR SEDERHANA ===")
// 		var angka_pertama, angka_kedua int
// 		var operator string

// 		fmt.Print("Masukkan angka pertama: ")
// 		fmt.Fscanln(reader, &angka_pertama)

// 		fmt.Print("Masukkan operator: ")
// 		fmt.Fscanln(reader, &operator)

// 		fmt.Print("Masukkan angka kedua: ")
// 		fmt.Fscanln(reader, &angka_kedua)

// 		switch operator {
// 		case "*":
// 			fmt.Println(angka_pertama * angka_kedua)
// 		case "/":
// 			fmt.Println(angka_pertama / angka_kedua)
// 		case "+":
// 			fmt.Println(angka_pertama + angka_kedua)
// 		case "-":
// 			fmt.Println(angka_pertama - angka_kedua)
// 		default:
// 			fmt.Println("Operator yang anda masukkan invalid")
// 		}

// 		fmt.Println("Lanjut Menghitung ? 1.Lanjut, 2.Keluar")
// 		fmt.Fscanln(reader, &pilihan)

// 		if pilihan != 1 {
// 			fmt.Println("Terimakasih")
// 			break
// 		}
// 	}
// }

// UPDATE
// Tantangan 1: Tambahkan Looping
// Tantangan 2: Validasi Operator
// Tantangan 3: Cegah Pembagian Nol
// Tantangan 4: Tambahkan Fungsi Modular
// Tantangan 5: Tambahkan Mode “Lanjutan”
// Tantangan 6: Buat Menu

func tambah(a, b float64) float64   { return a + b }
func kurang(a, b float64) float64   { return a - b }
func kali(a, b float64) float64     { return a * b }
func bagi(a, b float64) float64     { return a / b }
func pangkat(a, b float64) float64  { return math.Pow(a, b) }
func akarKuadrat(a float64) float64 { return math.Sqrt(a) }

func lanjut() bool {
	var pilihan string
	fmt.Print("Lanjut menghitung? (y/n): ")
	fmt.Scanln(&pilihan)
	return pilihan == "y" || pilihan == "Y"
}

func Calculator() {

	reader := bufio.NewReader(os.Stdin)

	var operator int
	var a, b float64

	for {
		fmt.Println("=== Selamat Datang ===")
		fmt.Printf("1. Tambah \n2. Kurang \n3. Kali \n4. Bagi \n5. Pangkat \n6. Akar Quadrat \n")
		fmt.Print("Pilih Operator: ")
		fmt.Fscanln(reader, &operator)
		switch operator {
		case 1:
			fmt.Print("Masukkan angka pertama: ")
			fmt.Fscanln(reader, &a)
			fmt.Print("Masukkan angka kedua: ")
			fmt.Fscanln(reader, &b)
			result := tambah(a, b)
			fmt.Printf("Hasil dari %v + %v = %v \n", a, b, result)
		case 2:
			fmt.Print("Masukkan angka pertama: ")
			fmt.Fscanln(reader, &a)
			fmt.Print("Masukkan angka kedua: ")
			fmt.Fscanln(reader, &b)
			result := kurang(a, b)
			fmt.Printf("Hasil dari %v - %v = %v \n", a, b, result)
		case 3:
			fmt.Print("Masukkan angka pertama: ")
			fmt.Fscanln(reader, &a)
			fmt.Print("Masukkan angka kedua: ")
			fmt.Fscanln(reader, &b)
			result := kali(a, b)
			fmt.Printf("Hasil dari %v * %v = %v \n", a, b, result)
		case 4:
			fmt.Print("Masukkan angka pertama: ")
			fmt.Fscanln(reader, &a)
			fmt.Print("Masukkan angka kedua: ")
			fmt.Fscanln(reader, &b)
			if b != 0 {
				result := bagi(a, b)
				fmt.Printf("Hasil dari %v / %v = %v \n", a, b, result)
			} else {
				fmt.Println("Tidak bisa membagi dengan 0")
			}
		case 5:
			fmt.Print("Masukkan angka pertama: ")
			fmt.Fscanln(reader, &a)
			fmt.Print("Masukkan angka kedua: ")
			fmt.Fscanln(reader, &b)
			result := pangkat(a, b)
			fmt.Printf("Hasil dari %v ^ %v = %v \n", a, b, result)
		case 6:
			fmt.Print("Masukkan angka pertama: ")
			fmt.Fscanln(reader, &a)
			result := akarKuadrat(a)
			fmt.Printf("Hasil dari √%v = %v \n", a, result)
		default:
			fmt.Println("Operator yang anda masukkan tidak valid")
		}
		if !lanjut() {
			fmt.Println("Terimakasih sudah menggunakan kalkulator")
			break
		}
	}
}
