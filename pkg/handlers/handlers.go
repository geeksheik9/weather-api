package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/geeksheik9/weatherAPI/pkg/api"
	"github.com/geeksheik9/weatherAPI/pkg/model"
	"github.com/geeksheik9/weatherAPI/pkg/service"
)

type WeatherService struct {
	Version  string
}

//Routes sets up the routes for the RESTful interface
func (s *WeatherService) Routes(r *mux.Router) *mux.Router {
	r.HandleFunc("/ping", s.PingCheck).Methods(http.MethodGet)
	r.HandleFunc("/weather", s.GetWeather).Methods(http.MethodGet)

	return r
}

//PingCheck checks that the app is running and returns 200, OK, version
func (s *WeatherService) PingCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK, " + s.Version))
}

//GetWeather is the handler function to return all weapons in the database
func (s *WeatherService) GetWeather(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("GetWeather invoked with url: %v", r.URL)

	lat := r.URL.Query().Get("lat")
	long := r.URL.Query().Get("lon")
	apiKey := r.URL.Query().Get("apiKey")

	if lat == "" || long == "" {
		api.RespondWithError(w, http.StatusBadRequest, "lat and long are required")
		return
	}

	weather, err := service.GetWeatherByLatLong(lat, long, apiKey)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var response model.Response

	response.SetTemp(weather.Main.Temp)
	response.SetFeelsLikeTemp(weather.Main.FeelsLike)
	response.SetWeather(weather.Weather[0].Main)

	api.RespondWithJSON(w, http.StatusOK, response)
}
