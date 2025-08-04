package beginner

import (
	"bufio"
	"fmt"
	"os"
)

func SalaryCalculator() {

	reader := bufio.NewReader(os.Stdin)

	var gaji, tunjangan float64
	pajak := 0.10

	fmt.Print("Masukkan Gaji Pokok: ")
	fmt.Fscan(reader, &gaji)
	fmt.Print("Masukkan Tunjangan: ")
	fmt.Fscan(reader, &tunjangan)

	gaji_bersih := (gaji + tunjangan) - ((gaji + tunjangan) * pajak)

	fmt.Printf("gaji bersih: %2.f", gaji_bersih)

}
