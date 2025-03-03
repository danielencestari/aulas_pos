package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}
func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/aulas_pos")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// inserir um novo produto - depois que fez a proteção contra SQL Injection
	product := NewProduct("Produto 2", 450.00)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

}

func insertProduct(db *sql.DB, product *Product) error {
	// proteger contra SQL Injection
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}
