package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/timereversal/keymaxweb/controllers"
	"github.com/timereversal/keymaxweb/models"
)

func (app *application) routes(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	carKey := controllers.CarKeyInfo{}
	carKeyInfoC := controllers.CarKeyInfo{
		CarInfoService: &models.CarKeyInfoService{DB: db},
	}
	r.Post("/carkeyinfo/post", carKey.Create)
	r.Post("/carkeyinfo2/post", carKeyInfoC.Create)
	return r
}
