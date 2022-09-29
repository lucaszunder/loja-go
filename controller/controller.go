package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/lucaszunder/loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.ListProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		priceFloat, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Erro na conversão do preco")
		}

		quantityInt, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Erro na conversão da quantidade")
		}
		models.CreateNewProduct(name, description, priceFloat, quantityInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.GetProduct(productId)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		priceFloat, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Erro na conversão do preco")
		}

		quantityInt, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Erro na conversão da quantidade")
		}
		models.UpdateProduct(id, name, description, priceFloat, quantityInt)
	}
	http.Redirect(w, r, "/", 301)
}
