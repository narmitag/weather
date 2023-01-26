package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

func process_file(filename string) Observations {
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var observations Observations
	json.Unmarshal(byteValue, &observations)

	return observations
}

func find_temp(observations Observations) (int, int) {
	var highest int = -99
	var lowest int = 99

	for i := 0; i < len(observations.Observations); i++ {
		if observations.Observations[i].Metric.TempHigh > highest {
			highest = observations.Observations[i].Metric.TempHigh
		}
		if observations.Observations[i].Metric.TempLow < lowest {
			lowest = observations.Observations[i].Metric.TempLow
		}
	}
	return lowest, highest
}

func main() {

	var observations Observations
	var highest_temp int
	var lowest_temp int
	var file_lowest int
	var file_highest int

	err := filepath.Walk("../data",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// fmt.Println(path)
			if filepath.Ext(path) == ".json" {

				observations = process_file(path)

				file_lowest, file_highest = find_temp(observations)
				if file_highest > highest_temp {
					highest_temp = file_highest
				}
				if file_lowest < lowest_temp {
					lowest_temp = file_lowest
				}

			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	println("Highest Temp = " + strconv.Itoa(highest_temp))
	println("Lowest Temp = " + strconv.Itoa(lowest_temp))
}
