package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/geeksheik9/weatherAPI/pkg/model"
)

func GetWeatherByLatLong(lat, long, apiKey string) (model.OverallWeather, error) {
	url := "https://api.openweathermap.org/data/2.5/weather?lat=" + lat +"&lon=" + long + "&appid=" + apiKey +  "&units=imperial"
	resp, err := http.Get(url)
	if err != nil {
		return model.OverallWeather{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.OverallWeather{}, err
	}

	var weather model.OverallWeather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return weather, err
	}

	return weather, nil
}