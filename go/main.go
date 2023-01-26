package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
	WindspeedHigh float32 `json:"windspeedHigh"`
	WindspeedLow  float32 `json:"windspeedLow"`
	WindspeedAvg  float32 `json:"windspeedAvg"`
	WindgustHigh  float32 `json:"windgustHigh"`
	WindgustLow   float32 `json:"windgustLow"`
	WindgustAvg   float32 `json:"windgustAvg"`
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
	PrecipRate    float32 `json:"precipRate"`
	PrecipTotal   float32 `json:"precipTotal"`
}

func main() {
	fmt.Println("Hello World")
	// Open our jsonFile
	var filename string = "20190123.json"
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + filename)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var observations Observations
	json.Unmarshal(byteValue, &observations)

	for i := 0; i < len(observations.Observations); i++ {
		fmt.Println("Time        : " + observations.Observations[i].ObsTimeUtc)
		fmt.Println("Average Temp: " + strconv.Itoa(observations.Observations[i].Metric.TempAvg))
		fmt.Println("Low Temp    : " + strconv.Itoa(observations.Observations[i].Metric.TempLow))
		fmt.Println("High Temp   : " + strconv.Itoa(observations.Observations[i].Metric.TempHigh))
	}

}
