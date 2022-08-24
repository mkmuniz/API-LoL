package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lol-data/src/repository/item"
)

func Run() {
	r := chi.NewRouter()
	r.Get("/items", http.HandlerFunc(item.GetAllItems))
	r.Get("/items/:name", http.HandlerFunc(item.GetItemByName))
	http.ListenAndServe(":8080", r)
}
