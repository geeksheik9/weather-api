package model

import "errors"

type Response struct {
	Temp string `json:"temp"`
	FeelsLikeTemp string `json:"feels_like_temp"`
	Weather string `json:"weather"`
}

func (r *Response) SetTemp(temp float64) error {
	switch {
		case temp <= 42:
			r.Temp = "Cold"
			break
		case temp > 42 && temp < 75:
			r.Temp = "Comfortable"
			break
		case temp > 75 && temp < 90:
			r.Temp = "Warm"
			break
		case temp >= 90:
			r.Temp = "Hot"
			break
		default:
			return errors.New("Invalid temperature")
	}
	return nil
}

func (r *Response) SetFeelsLikeTemp(temp float64) error {
	switch {
		case temp < 42:
			r.FeelsLikeTemp = "Cold"
			break
		case temp > 42 && temp < 75:
			r.FeelsLikeTemp = "Comfortable"
			break
		case temp > 75 && temp < 90:
			r.FeelsLikeTemp = "Warm"
			break
		case temp > 90:
			r.FeelsLikeTemp = "Hot"
			break
		default:
			return errors.New("Invalid temperature")
	}
	return nil
}

func (r *Response) SetWeather(weather string) error {
	switch(weather) {
		case "Clear":
			r.Weather = "Sunny"
			break
		case "Clouds":
			r.Weather = "Cloudy"
			break
		case "Rain":
			r.Weather = "Rainy"
			break
		case "Snow":
			r.Weather = "Snowy"
			break
		default:
			return errors.New("Invalid weather")
	}
	return nil
}
