package main

import (
	"flag"
	"net/http"

	"github.com/narmitag/weather/go/extract_data"
)

func main() {
	httpPtr := flag.Bool("http", false, "Start Http Server")
	dataPath := flag.String("datapath", "", "Path to the data files")

	flag.Parse()

	if *httpPtr {
		http.HandleFunc("/", extract_data.TestHandler(*dataPath))
		http.ListenAndServe(":8081", nil)
	} else {
		extract_data.ExtractData(*dataPath)
	}
}
