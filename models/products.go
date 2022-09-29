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

	allProducts, err := db.Query("select * from products order by id asc")

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

func GetProduct(productId string) Product {
	db := db.ConnectDatabase()
	getProduct, err := db.Query("select * from products where id=$1", productId)

	if err != nil {
		panic(err.Error())
	}

	p := Product{}

	for getProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = getProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
	}

	defer db.Close()
	return p
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

func DeleteProduct(id string) {
	db := db.ConnectDatabase()

	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id=$1;")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)

	defer db.Close()
}

func UpdateProduct(id, name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()

	updateDataQuery, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateDataQuery.Exec(name, description, price, quantity, id)

	defer db.Close()
}
