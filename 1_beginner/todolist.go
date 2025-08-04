package beginner

import (
	"bufio"
	"fmt"
	"os"
)

// Challenge 3: To-Do List di Terminal
// Materi yang dilatih: Slice, loop, switch

// Deskripsi:
// Tambah task
// Lihat daftar task
// Tandai task selesai / hapus
// 💡 Tantangan ekstra: Simpan ke file .txt (nanti di level intermediate)

// func Todolist() {

// 	reader := bufio.NewReader(os.Stdin)
// 	kegiatan := []string{"makan", "minum", "olahraga"}
// 	var pilih int
// 	var newKegiatan string
// 	var hapusKegiatan int

// 	for {

// 		fmt.Println("=== 1.Daftar Kegiatan, 2.Tambah Kegiatan, 3.Hapus Kegiatan 4.Exit === ")
// 		fmt.Print("Pilih: ")
// 		fmt.Fscan(reader, &pilih)
// 		if pilih == 1 {
// 			fmt.Println("=== Daftar Kegiatan ===")
// 			for index, n := range kegiatan {
// 				fmt.Println(index+1, n)
// 			}
// 		} else if pilih == 2 {
// 			fmt.Println("=== Tambah Kegiatan ===")
// 			fmt.Print("Nama Kegiatan: ")
// 			fmt.Fscan(reader, &newKegiatan)
// 			kegiatan = append(kegiatan, newKegiatan)
// 		} else if pilih == 3 {
// 			fmt.Println("=== Masukkan nomer kegiatan yang ingin di hapus ===")
// 			fmt.Print("Nomer: ")
// 			fmt.Fscan(reader, &hapusKegiatan)
// 			kegiatan = append(kegiatan[:hapusKegiatan-1], kegiatan[hapusKegiatan:]...)
// 		} else if pilih == 4 {
// 			break
// 		}

// 	}

// }

var tasks []string
var completed []bool

func addTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan task baru: ")
	newTask, _ := reader.ReadString('\n')
	tasks = append(tasks, newTask)
	completed = append(completed, false)
	fmt.Println("Task berhasil ditambahkan")
}

func showTask() {
	if len(tasks) == 0 {
		fmt.Println("Belum ada task")
		return
	}

	fmt.Println("Daftar task:")
	for i, task := range tasks {
		status := "[ ]"
		if completed[i] {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, task)
	}
}

func completedTask() {
	if len(tasks) == 0 {
		fmt.Println("Tidak ada task untuk diselesaikan")
		return
	}

	showTask()
	var index int
	fmt.Print("Pilih task yang selesai: ")
	fmt.Scanln(&index)

	if index < 1 || index > len(tasks) {
		fmt.Println("Nomer tidak ditemukan")
		return
	}

	completed[index-1] = true
	fmt.Printf("✅ Task \"%s\" ditandai selesai!\n", tasks[index-1])
}

func deleteTask() {
	if len(tasks) == 0 {
		fmt.Println("Tidak taks untuk dihapus")
		return
	}

	showTask()
	var index int
	fmt.Print("Masukkan task yang ingin dihapus: ")
	fmt.Scanln(&index)

	if index < 1 || index > len(tasks) {
		fmt.Println("Tidak ada task untuk dihapus")
		return
	}

	fmt.Printf("🗑️  Task \"%s\" dihapus!\n", tasks[index-1])

	tasks = append(tasks[:index-1], tasks[index:]...)
	completed = append(completed[:index-1], completed[index:]...)

}

func TodoList() {
	reader := bufio.NewReader(os.Stdin)

	var pilih_menu int

	for {
		fmt.Println("=== TO-DO LIST ===")
		fmt.Printf("1. Tambah Task \n2. Lihat Task \n3. Tadai Selesai \n4. Hapus Task \n5. Keluar \n")
		fmt.Print("Pilih Menu: ")
		fmt.Fscanln(reader, &pilih_menu)

		switch pilih_menu {
		case 1:
			addTask()
		case 2:
			showTask()
		case 3:
			completedTask()
		case 4:
			deleteTask()
		case 5:
			return
		default:
			fmt.Println("Angka yang anda masukkan tidak valid")
		}
	}
}

// UPDATE
// Tandai task selesai
// Validasi input
// Tampilkan task dengan nomor urut
// Gunakan fungsi modular
// Loop menu sampai user pilih keluar.
