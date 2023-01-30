package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
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

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
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

func httpserver(w http.ResponseWriter, _ *http.Request) {

	var observations Observations
	var highest_temp int
	var lowest_temp int

	var highest_temp_year int
	var lowest_temp_year int

	low_line := make([]opts.LineData, 0)
	high_line := make([]opts.LineData, 0)
	var xaxis []string

	years := []int{2019, 2020, 2021, 2022}
	months := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}

	for _, year := range years {

		println("")
		println("Year : " + strconv.Itoa(year))
		println("----------------------------")
		lowest_temp_year = 99
		highest_temp_year = -99

		for _, month := range months {
			highest_temp = -99
			lowest_temp = 99

			println("")
			println("+ Month : " + month)

			//xaxis = append(xaxis, strconv.Itoa(year)+month)
			err := filepath.Walk("../data/hourly/"+strconv.Itoa(year)+"/"+month,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					//fmt.Println(path)
					if filepath.Ext(path) == ".json" {

						observations = process_file(path)
						lowest_temp, highest_temp = find_temp(observations, lowest_temp, highest_temp)
						if lowest_temp != 99 {
							low_line = append(low_line, opts.LineData{Value: lowest_temp})
							high_line = append(high_line, opts.LineData{Value: highest_temp})
							xaxis = append(xaxis, fileNameWithoutExtension(filepath.Base(path)))
						}

					}
					return nil
				})
			if err != nil {
				log.Println(err)
			}
			println("  Highest Temp = " + strconv.Itoa(highest_temp))
			println("  Lowest Temp  = " + strconv.Itoa(lowest_temp))

			// low_line = append(low_line, opts.LineData{Value: lowest_temp})
			// high_line = append(high_line, opts.LineData{Value: highest_temp})

			if lowest_temp < lowest_temp_year {
				lowest_temp_year = lowest_temp
			}
			if highest_temp > highest_temp_year {
				highest_temp_year = highest_temp
			}
		}

		println("")
		println("++ Summary for " + strconv.Itoa(year))
		println("   Highest Temp = " + strconv.Itoa(highest_temp_year))
		println("   Lowest Temp  = " + strconv.Itoa(lowest_temp_year))

	}
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "High and Low Daily Temp",
		}))

	// Put data into instance
	line.SetXAxis(xaxis).
		AddSeries("Lowest Temp", low_line).
		AddSeries("Highest Temp", high_line).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}
