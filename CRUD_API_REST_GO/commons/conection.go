package commons

import (
	"log"

	"github.com/alejandrosubero/crud-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:admin@/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("migrando.....")
	db.AutoMigrate(&models.Persona{})
	db.AutoMigrate(&models.Direccion{})
}
