package main

import (
	"encoding/json"
	"net/http"
	"phones-review/gadgets/smartphones/web"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(sph *web.CreateSmartphoneHandler) *chi.Mux {
	mux := chi.NewMux()

	// globals middlewares
	mux.Use(
		middleware.Logger,    //log every http request
		middleware.Recoverer, //recover if a pannic occurs
	)

	mux.Post("/smartphones", sph.SaveSmartphoneHandler)
	mux.Get("/hello", helloHandler)

	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("done-by", "mauri")

	res := map[string]interface{}{
		"message": "hello world",
	}

	json.NewEncoder(w).Encode(res)
}
