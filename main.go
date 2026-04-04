package main

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
	log.Println("Day:", day)
	models.ReportForecast(weatherForecast)

}
