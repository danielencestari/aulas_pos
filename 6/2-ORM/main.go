package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key;autoIncrement;unique"`
	Name  string
	Price float64
}

func main() {
	dns := "root:root@tcp(localhost:3306)/aulas_pos"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Name: "Laptop", Price: 5600.00})
}
