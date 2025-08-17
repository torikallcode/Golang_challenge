package models

type Produk struct {
	ID    int
	Nama  string
	Harga int
	Stok  int
}

func (p *Produk) EditProduk(nama string, harga, stok int) {
	p.Nama = nama
	p.Harga = harga
	p.Stok = stok
}

type Total struct {
	TotalHarga int
	JumlahStok int
}
