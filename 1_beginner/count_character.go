package beginner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CountCharacter() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Kata/kalimat: ")

	kata_kalimat, _ := reader.ReadString('\n')
	kata_kalimat = strings.TrimSpace(kata_kalimat)

	jumlah_huruf := len(kata_kalimat)
	kata := strings.Fields(kata_kalimat)
	jumlah_kata := len(kata)

	fmt.Printf("Jumlah huruf:%d jumlah kata:%d", jumlah_huruf, jumlah_kata)

}
