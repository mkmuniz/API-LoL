package champion

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Get("/champions/free", http.HandlerFunc(GetFreeChampions))
}
