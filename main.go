package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/timereversal/keymaxweb/controllers"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	carsC := controllers.Car{}
	r.Post("/cars/new", carsC.New)
	err := http.ListenAndServe("8080", r)
	if err != nil {
		panic(err)
	}
}
