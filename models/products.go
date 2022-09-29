package models

import (
	_ "github.com/lib/pq"
	"github.com/lucaszunder/loja/db"
)

type Product struct {
	Name, Description string
	Price             float64
	Id, Quantity      int
}

func ListProducts() []Product {
	db := db.ConnectDatabase()

	allProducts, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for allProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()

	createDataQuery, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	createDataQuery.Exec(name, description, price, quantity)

	defer db.Close()
}
