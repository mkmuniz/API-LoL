package server

import (
	"log"
	"net/http"

	"github.com/lol-data/configs"
	"github.com/lol-data/src/middleware"
)

func Run() {
	configs.Load()

	log.Printf("API is running at port %v", configs.GetAPIConfig())
	http.ListenAndServe(configs.GetAPIConfig(), middleware.Middlewares())
}
