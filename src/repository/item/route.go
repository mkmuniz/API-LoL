package item

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Get("/items", http.HandlerFunc(GetAllItems))
	r.Get("/items/:name", http.HandlerFunc(GetItemByName))
}
