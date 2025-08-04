package beginner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Name  string
	Age   int
	Major string
}

func StudentData() {

	reader := bufio.NewReader(os.Stdin)

	var mahasiswa []Student
	var age int
	var pilihan int

	for {
		fmt.Printf("=== MANAGE DATA === \n1. Input Student \n2. Show Students \n3. Exit \n================== \n")
		fmt.Print("Masukkan Pilihan: ")
		fmt.Scanln(&pilihan)
		if pilihan == 1 {
			fmt.Println("=== Masukkan Data Mahasiswa ===")
			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Print("Age: ")
			fmt.Scanln(&age)
			fmt.Print("Major: ")
			major, _ := reader.ReadString('\n')
			major = strings.TrimSpace(major)
			mahasiswa = append(mahasiswa, Student{Name: name, Age: age, Major: major})
		} else if pilihan == 2 {
			if len(mahasiswa) == 0 {
				fmt.Println("Tidak ada data mahasiswa")
			} else {
				for index, i := range mahasiswa {
					fmt.Printf("=== Mahasiswa %v ===\n", index+1)
					fmt.Printf("Name: %v\nAge: %v\nMajor: %v\n", i.Name, i.Age, i.Major)
					// fmt.Println("==================")
				}
			}
		} else if pilihan == 3 {
			break
		}
	}

}
