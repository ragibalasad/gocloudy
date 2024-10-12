package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fatih/color"
)

var version string

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		TempC     float64 `json:"temp_c"`
		WindKph   float64 `json:"wind_kph"`
		Humidity  int64   `json:"humidity"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`

	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				WindKph   float64 `json:"wind_kph"`
				Humidity  int64   `json:"humidity"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	wf := DefineFlags()
	flag.Parse()

	if wf.Version {
		fmt.Printf("GoCloudy version: %s\n", version)
		return
	}

	city := "Rangpur" // default city
	if flag.NArg() > 0 {
		city = flag.Arg(0) // get the first non-flag argument
	}

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=1a79fe38d99243d89ac92056240110&q=" + city + "&days=1&aqi=no&alerts=no")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour
	fmt.Printf(
		"%s, %s: %.0f°C, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

	if wf.Detailed {
		fmt.Printf("Humidity: %d%%\n", current.Humidity)
		fmt.Printf("Wind Speed: %.0f km/h\n", current.WindKph)
	}

	// Print hourly forecast of the remaining day
	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		forecast_hour := fmt.Sprintf(
			"%s: %.0f°C",
			date.Format("15:04"),
			hour.TempC,
		)

		if wf.Detailed {
			forecast_hour += fmt.Sprintf(", %d%% Humidity, Wind Speed: %.0f km/h", hour.Humidity, hour.WindKph)
		}

		forecast_hour += fmt.Sprintf(
			", %.0f%% Chance of Rain, %s\n",
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

		if hour.ChanceOfRain > 40 {
			color.Red(forecast_hour)
		} else {
			fmt.Print(forecast_hour)
		}
	}
}
