package routes

import (
	"net/http"

	"github.com/lucaszunder/loja/controller"
)

func Load() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/update", controller.Update)
}
