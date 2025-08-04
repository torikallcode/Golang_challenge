package beginner

import (
	"bufio"
	"fmt"
	"os"
)

// func StudentGrades() {

// 	var nilai int
// 	fmt.Println("===Cek Nilai Siswa===")
// 	fmt.Print("Masukkan nilai: ")
// 	fmt.Scanln(&nilai)
// 	if nilai >= 85 {
// 		fmt.Println("Grade A")
// 	} else if nilai >= 75 {
// 		fmt.Println("Grade B")
// 	} else if nilai >= 65 {
// 		fmt.Println("Grade C")
// 	} else if nilai >= 40 {
// 		fmt.Println("Grade D")
// 	} else if nilai >= 0 {
// 		fmt.Println("Grade E")
// 	} else {
// 		fmt.Println("Nilai tidak Valid")
// 	}

// }

type Student_grades struct {
	Nama     string
	Nilai    int
	Predikat string
	Status   string
}

// type AllPredikat struct {
// 	Predikat_A int
// 	Predikat_B int
// 	Predikat_C int
// 	Predikat_D int
// 	Predikat_E int
// }

var students []Student_grades

func addStudents() {
	reader := bufio.NewReader(os.Stdin)
	var nilai int
	fmt.Print("Masukkan nama siswa: ")
	nama_siswa, _ := reader.ReadString('\n')
	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&nilai)
	grade := predikat(nilai)
	status := status(nilai)
	students = append(students, Student_grades{Nama: nama_siswa, Nilai: nilai, Predikat: grade, Status: status})

	// allPredikat[grade]++
}

func predikat(a int) string {
	if a >= 85 {
		return "Grade A"
	} else if a >= 75 {
		return "Grade B"
	} else if a >= 65 {
		return "Grade C"
	} else if a >= 50 {
		return "Grade D"
	} else if a >= 30 {
		return "Grade E"
	} else {
		return "Nilai Tidak valid"
	}
}

func status(a int) string {
	if a >= 70 {
		return "LULUS"
	} else {
		return "TIDAK LULUS"
	}
}

func daftar_siswa() {
	for index, i := range students {
		fmt.Printf("NO: %v \nNama: %sNilai: %d \nPredikat: %s \nStatus: %s \n", index+1, i.Nama, i.Nilai, i.Predikat, i.Status)
		fmt.Println("===============")
	}
}

func statistik_siswa() {
	nilai := 0
	nilai_tertinggi := 0
	nilai_terendah := 100
	for _, i := range students {
		nilai += i.Nilai
		if i.Nilai > nilai_tertinggi {
			nilai_tertinggi = i.Nilai
		} else if i.Nilai <= nilai_terendah {
			nilai_terendah = i.Nilai
		}
	}

	rata_rata := nilai / len(students)

	fmt.Println("Rata-rata nilai: ", rata_rata)
	fmt.Printf("Tertinggi: %d \n", nilai_tertinggi)
	fmt.Printf("Terendah: %d \n", nilai_terendah)
	fmt.Println("Jumlah siswa: ", len(students))
	// for key, value := range allPredikat {
	// 	fmt.Printf("%s : %d \n", key, value)
	// }

}

func StudentGrades() {

	// var allPredikat map[string]int
	// allPredikat["Predikat A"] = 0
	// allPredikat["Predikat B"] = 0
	// allPredikat["Predikat C"] = 0
	// allPredikat["Predikat D"] = 0
	// allPredikat["Predikat E"] = 0
	reader := bufio.NewReader(os.Stdin)

	var pilih_menu int

	for {
		fmt.Printf("1. Tambahkan Daftar Siswa \n2. Lihat Daftar Siswa \n3. Statistik Nilai Siswa \n4. Keluar \n")
		fmt.Print("Pilih menu: ")
		fmt.Fscanln(reader, &pilih_menu)
		switch pilih_menu {
		case 1:
			addStudents()
		case 2:
			daftar_siswa()
		case 3:
			statistik_siswa()
		}

	}
}
