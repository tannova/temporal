package weather

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/temporal"

	"temporal-server/model"
	"temporal-server/store"
)

func GetWeatherActivity(ctx context.Context, cityName string) (result *model.WeatherData, err error) {
	if cityName == "retry" {
		return nil, fmt.Errorf("trigger retry")
	}
	result, err = store.GetCurrentWeather(ctx, cityName)
	if err != nil {
		return result, temporal.NewApplicationError("unable to get weather data", "GET_WEATHER", err)
	}
	return result, nil
}
