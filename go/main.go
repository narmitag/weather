package main

import (
	"flag"
	"net/http"
)

type Observations struct {
	Observations []Observation `json:"observations"`
}

type Observation struct {
	StationID          string  `json:"stationID"`
	Tz                 string  `json:"tz"`
	ObsTimeUtc         string  `json:"obsTimeUtc"`
	ObsTimeLocal       string  `json:"obsTimeLocal"`
	Epoch              int     `json:"epoch"`
	Lat                float32 `json:"lat"`
	Lon                float32 `json:"lon"`
	SolarRadiationHigh float32 `json:"solarRadiationHigh"`
	UvHigh             float32 `json:"uvHigh"`
	WinddirAvg         int     `json:"winddirAvg"`
	HumidityHigh       float32 `json:"humidityHigh"`
	HumidityLow        float32 `json:"humidityLow"`
	HumidityAvg        float32 `json:"humidityAvg"`
	QcStatus           float32 `json:"qcStatus"`
	Metric             Metric  `json:"metric"`
}

type Metric struct {
	TempHigh      int     `json:"tempHigh"`
	TempLow       int     `json:"tempLow"`
	TempAvg       int     `json:"tempAvg"`
	WindspeedHigh int     `json:"windspeedHigh"`
	WindspeedLow  int     `json:"windspeedLow"`
	WindspeedAvg  int     `json:"windspeedAvg"`
	WindgustHigh  int     `json:"windgustHigh"`
	WindgustLow   int     `json:"windgustLow"`
	WindgustAvg   int     `json:"windgustAvg"`
	DewptHigh     float32 `json:"dewptHigh"`
	DewptLow      float32 `json:"dewptLow"`
	DewptAvg      float32 `json:"dewptAvg"`
	WindchillHigh float32 `json:"windchillHigh"`
	WindchillLow  float32 `json:"windchillLow"`
	WindchillAvg  float32 `json:"windchillAvg"`
	HeatindexHigh float32 `json:"heatindexHigh"`
	HeatindexLow  float32 `json:"heatindexLow"`
	HeatindexAvg  float32 `json:"heatindexAvg"`
	PressureMax   float32 `json:"pressureMax"`
	PressureMin   float32 `json:"pressureMin"`
	PressureTrend float32 `json:"pressureTrend"`
	PrecipRate    float64 `json:"precipRate"`
	PrecipTotal   float64 `json:"precipTotal"`
}

func main() {
	httpPtr := flag.Bool("http", false, "Start Http Server")

	flag.Parse()

	if *httpPtr {
		http.HandleFunc("/", httpserver)
		http.ListenAndServe(":8081", nil)
	} else {
		extract_data()
	}
}
