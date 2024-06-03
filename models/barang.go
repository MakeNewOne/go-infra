package models

type Barang struct {
	ID     int 
	Nama   string
	Berat  int
	Jenis  string
	Harga  string
	UserID int
	User   User
}

type BarangUser struct {
	ID    int
	Nama  string
	Berat int
	Jenis string
	Harga string
}

func (BarangUser) TableName() string {
	return "Barang"
}
