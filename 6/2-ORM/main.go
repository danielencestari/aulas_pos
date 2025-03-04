package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dns := "root:root@tcp(localhost:3306)/aulas_pos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{
		Name:  "Laptop",
		Price: 5600.00,
	})

	// cria um lote de produtos
	products := []Product{
		{Name: "Mouse", Price: 50.00},
		{Name: "Teclado", Price: 150.00},
		{Name: "Monitor", Price: 800.00},
	}
	db.Create(&products)

	// select a product
	var product Product
	db.First(&product, 2)
	fmt.Println(product)
	// busca por nome:
	db.First(&product, "name = ?", "Mouse")
	fmt.Println(product)

	// select all
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}

	// where
	// var products []Product
	// db.Where("price > ?", 90).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Where("name LIKE ?", "%k%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2.Name)
	// db.Delete(&p2)

}
