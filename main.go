package main

import (
	"database/sql"

	db2 "github.com/Carrilh0/hexagonal-architecture/adapters/db"
	"github.com/Carrilh0/hexagonal-architecture/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, err := productService.Create("Product Test", 10)
	if err != nil {
		panic(err)
	}

	productService.Enable(product)
}
