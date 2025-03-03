package main

import (
	"database/sql"
	"fmt"

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
	product := NewProduct("Produto 1", 150.00)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	// update de um produto - mudar o preco de 450 pra 200
	product.Price = 100.00
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// select de um produto
	// p, err := selectProduct(db, product.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Produto: %+v , possui o preço de %.2f", p.Name, p.Price)

	// select all products
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Produto: %+v , possui o preço de %.2f\n", p.Name, p.Price)
	}

	// delete a product
	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

// func que insere um produto
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

// update a product
func updateProduct(db *sql.DB, product *Product) error {
	// faz o statement e prepara pra não ter SQL Injection
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}

	// tem que ter o defer close pra não esquecer de fechar
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

// select a product
func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// select all
func selectAllProducts(db *sql.DB) ([]Product, error) {
	// não precisa de prepare pq não tem SQL Injection
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	// tem q ter o prepare pra não ter SQL Injection
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// exec pq não é uma consulta, mas sim uma deleção
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
