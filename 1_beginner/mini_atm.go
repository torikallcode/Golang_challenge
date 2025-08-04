package beginner

import (
	"bufio"
	"fmt"
	"os"
)

func MiniAtm() {

	reader := bufio.NewReader(os.Stdin)

	saldo := 1000000
	var pilihan, tarik_tunai, nominal_transfer int
	var rekening_penerima string

	fmt.Println("==== Selamat Datang ====")
	for {
		fmt.Printf("1. Cek Saldo \n2. Tarik tunai \n3. Transfer \n4. Keluar \n")
		fmt.Print("Masukkan pilihan anda: ")
		fmt.Fscanln(reader, &pilihan)
		switch pilihan {
		case 1:
			fmt.Println("Saldo anda sebesar:", saldo)
		case 2:
			fmt.Print("Masukkan Jumlah yang ingin di tarik: ")
			fmt.Fscanln(reader, &tarik_tunai)
			fmt.Println("Tarik tunai sebesar:", tarik_tunai)
			fmt.Println("Saldo anda sekarang", saldo-tarik_tunai)
		case 3:
			fmt.Print("Masukkan nomer rekening penerima: ")
			fmt.Fscanln(reader, &rekening_penerima)
			fmt.Print("Masukkan nominal transfer: ")
			fmt.Fscanln(reader, &nominal_transfer)
			fmt.Println("== Transfer berhasil ==")
			fmt.Println("Penerima: ", rekening_penerima)
			fmt.Println("Jumlah Transfer: ", nominal_transfer)
		case 4:
			return
		}

	}

}
