package models

type User struct {
	ID       int
	Name     string
	Username string
	Password string
	Barang   BarangUser
}
