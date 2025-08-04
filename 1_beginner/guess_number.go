package beginner

import (
	"fmt"
	"math/rand"
)

// Challenge 4: Tebak Angka
// Materi yang dilatih: Random number, loop, if-else

// Deskripsi:
// Komputer memilih angka 1-100 secara acak
// User menebak sampai benar
// Beri petunjuk: “Terlalu kecil” atau “Terlalu besar”

// func GuessNumber() {

// 	reader := bufio.NewReader(os.Stdin)

// 	var pilihan_user int
// 	angka_acak := rand.Intn(100)
// 	fmt.Println("=== Game Tebak Angka ===")

// 	for {
// 		fmt.Print("Masukkan pilihanmu: ")
// 		fmt.Fscan(reader, &pilihan_user)
// 		if pilihan_user == angka_acak {
// 			fmt.Printf("Tebakan anda benar: %d", angka_acak)
// 			break
// 		} else if pilihan_user > angka_acak {
// 			fmt.Println("Angka anda terlalu besar")
// 		} else if pilihan_user < angka_acak {
// 			fmt.Println("Angka anda terlalu kecil")
// 		}
// 	}
// }

// UPDATE
// Batasi jumlah percobaan
// Beri mode kesulitan
// Berikan skor
// Looping Main Lagi

func level(a int) {
	tebakan_user := 0
	jumlah_tebakan := 0
	if a == 1 {
		angka_acak := rand.Intn(50)
		for {
			fmt.Print("Masukkan Pilihan: ")
			if jumlah_tebakan > 10 {
				fmt.Printf("Game Over, angkanya adalah %d", angka_acak)
				return
			} else {
				jumlah_tebakan++
				Skor := 100 - (jumlah_tebakan * 10)
				fmt.Scanln(&tebakan_user)
				if tebakan_user < angka_acak {
					fmt.Println("Tebakan anda lebih kecil")
				} else if tebakan_user > angka_acak {
					fmt.Println("Tebakan anda lebih besar")
				} else {
					fmt.Printf("Tebakan anda benar: %d \n", angka_acak)
					fmt.Printf("Skor anda: %d", Skor)
					return
				}
			}
		}
	} else if a == 2 {
		angka_acak := rand.Intn(100)
		for {
			fmt.Print("Masukkan Pilihan: ")
			if jumlah_tebakan > 10 {
				fmt.Printf("Game Over, angkanya adalah %d", angka_acak)
				return
			} else {
				jumlah_tebakan++
				Skor := 100 - (jumlah_tebakan * 10)
				fmt.Scanln(&tebakan_user)
				if tebakan_user < angka_acak {
					fmt.Println("Tebakan anda lebih kecil")
				} else if tebakan_user > angka_acak {
					fmt.Println("Tebakan anda lebih besar")
				} else {
					fmt.Printf("Tebakan anda benar: %d \n", angka_acak)
					fmt.Printf("Skor anda: %d", Skor)
					return
				}
			}
		}
	} else if a == 3 {
		angka_acak := rand.Intn(500)
		for {
			fmt.Print("Masukkan Pilihan: ")
			if jumlah_tebakan > 10 {
				fmt.Printf("Game Over, angkanya adalah %d", angka_acak)
				return
			} else {
				jumlah_tebakan++
				Skor := 100 - (jumlah_tebakan * 10)
				fmt.Scanln(&tebakan_user)
				if tebakan_user < angka_acak {
					fmt.Println("Tebakan anda lebih kecil")
				} else if tebakan_user > angka_acak {
					fmt.Println("Tebakan anda lebih besar")
				} else {
					fmt.Printf("Tebakan anda benar: %d \n", angka_acak)
					fmt.Printf("Skor anda: %d", Skor)
					return
				}
			}
		}
	}
}

func GuessNumber() {

	var pilih_level int

	fmt.Println("=== Game Tebak Angka ===")
	fmt.Printf("1. Mudah \n2. Normal \n3. Sulit \n")
	fmt.Print("Pilih Level: ")
	fmt.Scanln(&pilih_level)
	level(pilih_level)
}
