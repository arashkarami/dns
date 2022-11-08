package test

import (
	"housinganywhere/dns/app"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestNavigationHandler_ShouldReturn_StatusMethodNotAllowed_WhenRequestMethodIsNot_Post(t *testing.T) {
	req, err := http.NewRequest("GET", "", strings.NewReader("{ \"x\":\"123.12\", \"y\":\"456.56\", \"z\":\"789.89\", \"vel\":\"20\"}"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	env := LoadENV()
	r := mux.NewRouter()
	app := app.New(r, env)

	handler := http.HandlerFunc(app.NavigationHandler())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestNavigationHandler_ShouldReturn_StatusBadRequest_WhenRequestBodyIsInvalid(t *testing.T) {
	req, err := http.NewRequest("POST", "", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	env := LoadENV()
	r := mux.NewRouter()
	app := app.New(r, env)

	handler := http.HandlerFunc(app.NavigationHandler())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestNavigationHandler_ShouldReturn_CorrectValue(t *testing.T) {
	req, err := http.NewRequest("POST", "", strings.NewReader("{ \"x\":\"123.12\", \"y\":\"456.56\", \"z\":\"789.89\", \"vel\":\"20\"}"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	env := LoadENV()
	r := mux.NewRouter()
	app := app.New(r, env)

	handler := http.HandlerFunc(app.NavigationHandler())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"loc":1389.5700000000002}
	`
	got := rr.Body.String()
	if strings.Compare(got, expected) > 0 {
		t.Errorf("handler returned unexpected body: got [%v] want %v -- %d", got, expected, strings.Compare(got, expected))
	}
}

func LoadENV() map[string]string {
	ENV := map[string]string{}

	ENV["SECTORID"] = "1"

	return ENV
}
