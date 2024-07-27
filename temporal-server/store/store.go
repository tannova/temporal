package store

import (
	"context"

	"temporal-server/model"
)

func GetCurrentWeather(ctx context.Context, cityName string) (*model.WeatherData, error) {
	// code to fetch the current weather for a given city
	return &model.WeatherData{
		Temperature: 36,
		Humidity:    40,
		WindSpeed:   21,
		City:        cityName,
	}, nil
}
