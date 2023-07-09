package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
)

type application struct {
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
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
	PostgresCfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
	addr := flag.String("addr", ":6020", "Http Server Address:Port")

	flag.Parse()
	app := &application{}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	openDB(PostgresCfg.String())

	fmt.Printf("Starting server on %s \n", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("error: ", err)
	}

}

func openDB(cfg string) (*sql.DB, error) {

	db, err := sql.Open("pgx", cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected!")

	return db, nil
}
