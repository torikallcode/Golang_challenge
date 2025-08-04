package beginner

import (
	"bufio"
	"fmt"
	"os"
)

// func convertToCelcius(a int) int {
// 	result := (a - 32) * 5 / 9

// 	return result
// }

// func ConvertToFahrenheit(a int) int {
// 	result := (a * 9 / 5) + 32

// 	return result
// }

// func Convert() {

// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		var fahrenheit, celcius int
// 		var pilihanSuhu, pilihanLanjut int

// 		fmt.Println("Pilih Convert Suhu")
// 		fmt.Println("1.Fahrenheit 2.Celcius")
// 		fmt.Fscan(reader, &pilihanSuhu)

// 		if pilihanSuhu == 1 {
// 			fmt.Print("Masukkan suhu Celcius: ")
// 			fmt.Fscan(reader, &celcius)
// 			fmt.Println(ConvertToFahrenheit(celcius))
// 		} else if pilihanSuhu == 2 {
// 			fmt.Print("Masukkan suhu Fahrenheit: ")
// 			fmt.Fscan(reader, &fahrenheit)
// 			fmt.Println(convertToCelcius(fahrenheit))
// 		} else {
// 			fmt.Println("Angka yang anda masukkan tidak valid")
// 		}
// 		fmt.Println("Lanjut hitung? 1.Lanjut, 2.Keluar")
// 		fmt.Fscan(reader, &pilihanLanjut)
// 		if pilihanLanjut != 1 {
// 			fmt.Println("Terimakasih")
// 			break
// 		}
// 	}

// }

// Tambah Konversi ke Kelvin
// Looping Hingga User Keluar
// Validasi Angka
// Format Output Lebih Rapi

func CelciusToFahrenheit(a float64) float64 { return (a * 9 / 5) + 32 }
func FahrenheitToCelcius(a float64) float64 { return (a - 32) * 5 / 9 }
func CelciusToKelvin(a float64) float64     { return a + 273.15 }
func KelvinToCelcius(a float64) float64     { return a - 273.15 }

func lanjut_suhu() bool {
	var pilihan string
	fmt.Print("Lanjut Konversi? (y/n): ")
	fmt.Scanln(&pilihan)
	return pilihan == "y" || pilihan == "Y"
}

func ConverSuhu() {

	reader := bufio.NewReader(os.Stdin)

	var pilihan int
	var input_suhu float64

	fmt.Println("=== Konversi Suhu === ")

	for {
		fmt.Printf("1. Celsius → Fahrenheit \n2. Fahrenheit → Celsius \n3. Celsius → Kelvin \n4. Kelvin → Celsius \n")
		fmt.Print("Pilih menu: ")
		fmt.Fscanln(reader, &pilihan)
		switch pilihan {
		case 1:
			fmt.Print("Masukkan Suhu: ")
			fmt.Fscanln(reader, &input_suhu)
			hasil := CelciusToFahrenheit(input_suhu)
			fmt.Printf("Hasil %.2f°C = %.2f°F \n", input_suhu, hasil)
		case 2:
			fmt.Print("Masukkan Suhu: ")
			fmt.Fscanln(reader, &input_suhu)
			hasil := FahrenheitToCelcius(input_suhu)
			fmt.Printf("Hasil %.2f°F = %.2f°C \n", input_suhu, hasil)
		case 3:
			fmt.Print("Masukkan Suhu: ")
			fmt.Fscanln(reader, &input_suhu)
			hasil := CelciusToKelvin(input_suhu)
			fmt.Printf("Hasil %.2f°C = %.2f°K \n", input_suhu, hasil)
		case 4:
			fmt.Print("Masukkan Suhu: ")
			fmt.Fscanln(reader, &input_suhu)
			hasil := KelvinToCelcius(input_suhu)
			fmt.Printf("Hasil %.2f°K = %.2f°C \n", input_suhu, hasil)
		}

		if !lanjut_suhu() {
			fmt.Println("Terimakasih sudah memakai program convert suhu")
			break
		}
	}
}

