package main

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func extract_data() {

	var observations Observations
	var highest_temp int
	var lowest_temp int

	var highest_temp_year int
	var lowest_temp_year int

	years := []int{2019, 2020, 2021, 2022}
	months := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	headers := []string{"Timestamp", "tempLow", "tempHigh", "WindspeedHigh", "WindspeedLow",
		"WindspeedAvg", "WindgustHigh", "WindgustLow", "WindgustAvg", "PrecipRate", "PrecipTotal"}
	csvFile, err := os.Create("data.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	w := csv.NewWriter(csvFile)
	defer w.Flush()
	if err := w.Write(headers); err != nil {
		log.Fatalln("error writing record to file", err)
	}

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

			err := filepath.Walk("../data/hourly/"+strconv.Itoa(year)+"/"+month,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					//fmt.Println(path)
					if filepath.Ext(path) == ".json" {

						observations = process_file(path)
						lowest_temp, highest_temp = find_data(observations, lowest_temp, highest_temp, *w)

					}
					return nil
				})
			if err != nil {
				log.Println(err)
			}
			println("  Highest Temp = " + strconv.Itoa(highest_temp))
			println("  Lowest Temp  = " + strconv.Itoa(lowest_temp))

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
	w.Flush()
	csvFile.Close()
}

func find_data(observations Observations,
	lowest int,
	highest int, w csv.Writer) (int, int) {

	for i := 0; i < len(observations.Observations); i++ {
		if observations.Observations[i].Metric.TempHigh > highest {
			highest = observations.Observations[i].Metric.TempHigh
		}
		if observations.Observations[i].Metric.TempLow < lowest {
			lowest = observations.Observations[i].Metric.TempLow
		}
		row := []string{observations.Observations[i].ObsTimeUtc,
			strconv.Itoa(observations.Observations[i].Metric.TempLow),
			strconv.Itoa(observations.Observations[i].Metric.TempHigh),
			strconv.Itoa(observations.Observations[i].Metric.WindspeedHigh),
			strconv.Itoa(observations.Observations[i].Metric.WindspeedLow),
			strconv.Itoa(observations.Observations[i].Metric.WindspeedAvg),
			strconv.Itoa(observations.Observations[i].Metric.WindgustHigh),
			strconv.Itoa(observations.Observations[i].Metric.WindgustLow),
			strconv.Itoa(observations.Observations[i].Metric.WindgustAvg),
			strconv.FormatFloat(observations.Observations[i].Metric.PrecipRate, 'E', -1, 32),
			strconv.FormatFloat(observations.Observations[i].Metric.PrecipTotal, 'E', -1, 32),
		}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	return lowest, highest
}
