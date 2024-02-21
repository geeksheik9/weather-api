package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"

	"github.com/geeksheik9/weatherAPI/pkg/handlers"
)

var version string

func main() {
	logrus.Info("INITIALIZING Weather API")

	weatherService := handlers.WeatherService{
		Version:  version,
	}

	r := mux.NewRouter().StrictSlash(true)

	r = weatherService.Routes(r)
	fmt.Printf("Server listen on port %v\n", 3000)
	logrus.Info("END")
	logrus.Fatal(http.ListenAndServe(":"+"3000", cors.AllowAll().Handler(r)))
}