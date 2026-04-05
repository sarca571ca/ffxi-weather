package main

// NOTE: Weather forecasters are correct but its not usefull for us as KV spawns at the start
// of weather. Meaning if there is none | none | none or sunshine | sunshine | sunshine there
// is zero chance of a spawn. As overlapping weather doesn't help in this case.

// WARNING: If this module is to be expanded to cover all weathers and all zones we will need to
// expand the logic also. Need to give 2 forecasts one for on weather change and a normal
// overlapping weather forecast that ends at 0700 or 1300.

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sarca571ca/ffxi-weather/internal/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	weatherBlob := os.Getenv("BLOB")

	wb, err := models.NewWeatherBlob(weatherBlob)
	if err != nil {
		log.Fatal(err)
	}

	vt, err := models.FetchCurrentVanaTime()
	var day models.DayWeather

	if err == nil {
		day, err = wb.WeatherForVanaDay(models.AbsoluteVanaDay(vt.Year, vt.Month, vt.Day))
	} else {
		day, err = wb.WeatherForNowFallback(time.Now())
	}

	if err != nil {
		log.Fatal(err)
	}
	weatherForecast, err := wb.BuildWeatherForecast(10, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Day:", day.VanaDay%2160)
	models.ReportForecast(weatherForecast)

}
