package manajemenuser

import (
	"fmt"
	"golang_challenge/2_inttermediatte/Manajemen_user/services"
	"golang_challenge/2_inttermediatte/Manajemen_user/utils"
)

func ManajemenUser() {

	for {
		fmt.Println("=== Manajemen User ===")
		fmt.Println("1. Tambahkan User")
		fmt.Println("2. Daftar User")
		fmt.Println("3. Edit Email")
		fmt.Println("4. Hapus User")
		fmt.Println("5. Keluar")
		pilihan := utils.PromptInt("Masukkan Pilihan: ")
		switch pilihan {
		case 1:
			nama := utils.PromptString("Masukkan nama:")
			email := utils.PromptString("Masukkan email: ")
			user := services.AddUser(nama, email)
			fmt.Println(user)
		case 2:
			users := services.GetAllUser()
			if len(users) == 0 {
				fmt.Println("belum ada user")
			} else {
				fmt.Println("=== Daftar User ===")
				for _, i := range users {
					fmt.Printf("No: %d\n", i.ID)
					fmt.Printf("Nama: %s\n", i.Name)
					fmt.Printf("Email: %s\n", i.Email)
				}
			}
		case 3:
			pilihUser := utils.PromptInt("Masukkan id User: ")
			newEmail := utils.PromptString("Masukkan email baru: ")
			services.EditEmail(pilihUser, newEmail)
		case 4:
			pilihUser := utils.PromptInt("Masukkan id User: ")
			services.DeleteUser(pilihUser)
		case 5:
			return
		}
	}

}
