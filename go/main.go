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

func find_temp(observations Observations,
	lowest int,
	highest int) (int, int) {

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

	years := []int{2019, 2020, 2021, 2022}
	months := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}

	for _, year := range years {

		println("")
		println("Year : " + strconv.Itoa(year))
		println("----------------------------")
		for _, month := range months {
			highest_temp = -99
			lowest_temp = 99

			println("")
			println("+ Month : " + month)
			err := filepath.Walk("../data/hourly/"+strconv.Itoa(year)+"/"+month,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					//fmt.Println(path)
					if filepath.Ext(path) == ".json" {

						observations = process_file(path)
						lowest_temp, highest_temp = find_temp(observations, lowest_temp, highest_temp)

					}
					return nil
				})
			if err != nil {
				log.Println(err)
			}
			println("  Highest Temp = " + strconv.Itoa(highest_temp))
			println("  Lowest Temp  = " + strconv.Itoa(lowest_temp))
		}
	}
}
