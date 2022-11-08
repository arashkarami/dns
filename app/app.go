package app

import (
	"housinganywhere/dns/app/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type App struct {
	Router       *mux.Router
	DroneService services.DroneService
}

func New(router *mux.Router, envs map[string]string) *App {
	a := &App{
		Router: router,
	}

	a.initRoutes()

	sectorId := envs["SECTORID"]
	if sId, err := strconv.ParseFloat(sectorId, 32); err == nil {
		a.DroneService = services.NewDroneService(sId)
	}

	return a
}

func (a *App) initRoutes() {
	api := a.Router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/navigation", a.NavigationHandler()).Methods(http.MethodPost)
	api.HandleFunc("/health", healthcheck).Methods(http.MethodGet)
}
