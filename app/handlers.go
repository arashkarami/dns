package app

import (
	"encoding/json"
	"housinganywhere/dns/app/models"
	"io"
	"log"
	"net/http"
)

// swagger:route POST /api/v1/navigation locationRequest
// Calculate navigation of drone
//
// security:
// - apiKey: []
//
// responses:
//
//	200: LocationResponse
func (a *App) NavigationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			sendResponse(w, r, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		req := models.LocationRequest{}
		err := parse(w, r, &req)
		if err != nil {
			sendResponse(w, r, err, http.StatusBadRequest)
			return
		}
		loc, err := a.DroneService.GetLocation(&req)
		if err != nil {
			sendResponse(w, r, "Error in calcuation", http.StatusInternalServerError)
			return
		}
		sendResponse(w, r, loc, http.StatusOK)

	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"alive": true}`)
}

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func sendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}
