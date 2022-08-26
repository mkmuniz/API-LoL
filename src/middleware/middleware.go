package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lol-data/src/repository/champion"
)

func Middlewares() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Group(champion.Routes)

	return r
}
