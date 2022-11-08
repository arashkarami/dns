package main

import (
	"fmt"
	"housinganywhere/dns/app"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	env := LoadENV()
	r := mux.NewRouter()

	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)

	r.Handle("/docs", sh)

	http.Handle("/", r)
	app := app.New(r, env)
	r = app.Router
	s := &http.Server{Addr: fmt.Sprintf("%s:%s", "localhost", "5001")}
	s.ListenAndServe()
}

func LoadENV() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed On Loading Local Env File !")
	}

	ENV := map[string]string{}

	ENV["SECTORID"] = os.Getenv("SECTORID")

	return ENV
}
