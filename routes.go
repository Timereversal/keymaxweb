package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/timereversal/keymaxweb/controllers"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()
	carKey := controllers.CarKeyInfo{}
	r.Post("/carkeyinfo/post", carKey.Create)
	return r
}
