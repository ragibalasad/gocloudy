package main

import (
	"flag"
)

type WeatherFlags struct {
	Detailed     bool
	ForecastDays int
	Unit         string
}

func DefineFlags() *WeatherFlags {
	wf := &WeatherFlags{}

	// Define the flags
	flag.BoolVar(&wf.Detailed, "detailed", false, "Show extended weather details (humidity, wind speed, etc.)")
	flag.IntVar(&wf.ForecastDays, "days", 3, "Number of days for weather forecast")
	flag.StringVar(&wf.Unit, "unit", "C", "Temperature unit (C or F)")

	return wf
}
