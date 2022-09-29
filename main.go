package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/lucaszunder/loja/routes"
)

func main() {
	routes.Load()
	http.ListenAndServe(":8000", nil)
}
