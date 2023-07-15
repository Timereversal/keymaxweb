package main

import (
	"flag"
	"fmt"
	"github.com/timereversal/keymaxweb/models"
	"log"
	"net/http"
)

type application struct {
}

func main() {
	//r := chi.NewRouter()
	//carsC := controllers.Car{}
	//r.Post("/cars/new", carsC.New)
	//err := http.ListenAndServe("8080", r)
	//if err != nil {
	//	panic(err)
	//
	//}

	cfg := models.DefaultPostgresConfig()
	db, err := models.OpenDB(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	addr := flag.String("addr", ":6020", "Http Server Address:Port")

	flag.Parse()

	app := &application{}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(db),
	}

	fmt.Printf("Starting server on %s \n", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("error: ", err)
	}

}
