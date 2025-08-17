package services

import (
	"fmt"
	"golang_challenge/2_inttermediatte/Manajemen_produk_stok/models"
)

var produks []models.Produk
var indeks = 1

func AddProduk(nama string, harga, stok int) models.Produk {
	produk := models.Produk{
		ID:    indeks,
		Nama:  nama,
		Harga: harga,
		Stok:  stok,
	}
	produks = append(produks, produk)
	indeks++
	fmt.Println(produks)
	return produk
}

func GetAllProduk() []models.Produk {
	return produks
}

func EditProduk(id, harga, stok int, nama string) bool {
	for i := range produks {
		if produks[i].ID == id {
			produks[i].EditProduk(nama, harga, stok)
			return true
		}
	}
	return false
}

func HapusProduk(id int) bool {
	for i := range produks {
		if produks[i].ID == id {
			produks = append(produks[:i], produks[i+1:]...)
			return true
		}
	}
	return false
}

func CariProduk(nama string) bool {
	for i, val := range produks {
		if produks[i].Nama == nama {
			fmt.Printf("no: %d\n", val.ID)
			fmt.Printf("nama: %s\n", val.Nama)
			fmt.Printf("harga: %d\n", val.Harga)
			fmt.Printf("stok: %d\n", val.Stok)
			return true
		}
	}
	return false
}

func TotalProduk() (totalHarga, jumlahStok, result int) {
	totalHarga = 0
	jumlahStok = 0
	for _, i := range produks {
		totalHarga += i.Harga
		jumlahStok += i.Stok
	}
	result = totalHarga * jumlahStok
	fmt.Println("Total Harga semua barang", totalHarga)
	fmt.Println("Total semua stok", jumlahStok)
	fmt.Println("Total", result)

	return totalHarga, jumlahStok, result
}
