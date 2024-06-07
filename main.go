package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User struct untuk menyimpan data user
type User struct {
	Nama    string `json:"nama"`
	Nomor   string `json:"nomor"`
	Created string `json:"created"`
}

// Initialize database
var db *gorm.DB

func init() {
	// Konfigurasi koneksi database
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USER")
	// pass := os.Getenv("DB_PASS")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/test", user, pass, host, port)
	var err error
	db, err = gorm.Open(mysql.Open("root:cog938gb18@tcp(localhost:3306)/test"), &gorm.Config{})
	if err != nil {
		panic("Tidak dapat terhubung ke database")
	}

	// Migrasi skema database
	db.AutoMigrate(&User{}, "testing")
}

func main() {
	e := echo.New()

	// Endpoint untuk membuat user baru
	e.POST("/user", createUser)

	e.Logger.Fatal(e.Start(":8000"))
}

// Handler untuk membuat user baru
func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	if err := db.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (User) TableName() string {
	return "testing"
}
