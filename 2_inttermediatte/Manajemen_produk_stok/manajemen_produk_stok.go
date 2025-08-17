package manajemenprodukstok

import (
	"fmt"
	"golang_challenge/2_inttermediatte/Manajemen_produk_stok/services"
	"golang_challenge/2_inttermediatte/Manajemen_produk_stok/utils"
)

func ManajemenProdukStok() {
	for {
		fmt.Println("=== Manajemen Produk ===")
		fmt.Println("1. Tambah Produk")
		fmt.Println("2. Edit Produk")
		fmt.Println("3. Hapus Produk")
		fmt.Println("4. Cari Produk")
		fmt.Println("5. Hitung Stok")
		pilih := utils.PromptInt("Pilih: ")
		switch pilih {
		case 1:
			nama := utils.PromptString("Masukkan nama barang: ")
			harga := utils.PromptInt("Masukkan harga barang: ")
			stok := utils.PromptInt("Masukkan stok barang: ")
			services.AddProduk(nama, harga, stok)
		case 2:
			pilihProduk := utils.PromptInt("Masukkan id produk: ")
			nama := utils.PromptString("Masukkan nama barang: ")
			harga := utils.PromptInt("Masukkan harga barang: ")
			stok := utils.PromptInt("Masukkan stok barang: ")
			services.EditProduk(pilihProduk, harga, stok, nama)
		case 3:
			pilihProduk := utils.PromptInt("Masukkan id produk: ")
			services.HapusProduk(pilihProduk)
		case 4:
			produks := services.GetAllProduk()
			if len(produks) == 0 {
				fmt.Println("Produk tidak ada")
			} else {
				for _, i := range produks {
					fmt.Printf("no: %d\n", i.ID)
					fmt.Printf("nama: %s\n", i.Nama)
				}
			}
			cariProduk := utils.PromptString("Masukkan nama produk: ")
			services.CariProduk(cariProduk)
		case 5:
			services.TotalProduk()
		case 6:
			return
		}
	}
}
