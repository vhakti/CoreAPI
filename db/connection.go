package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN string = "host=192.168.8.62 port=5432 user=azuan dbname=e-catalog password=M4rt1n4L4P3l1gr0s4"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		//panic(error)
		log.Fatal(error)
	} else {
		log.Println("Connected")
	}
}
