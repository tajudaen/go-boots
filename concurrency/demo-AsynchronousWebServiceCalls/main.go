package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"runtime"
)

type QuoteResponse struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}

func main() {
	runtime.GOMAXPROCS(3) // using Parallelism to improve performance
	start := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	numComplete := 0

	for _, symbol := range stockSymbols {
		go func(symbol string) { // concurrency added
			// make request to endpoint
			resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + symbol)
			// postpone the close of connection
			defer resp.Body.Close()

			// read data from the returned response
			body, _ := ioutil.ReadAll(resp.Body)

			// create a variable of type QuoteResponse
			quote := new(QuoteResponse)

			// the raw data(body variable) is store in a destination object(&quote, which is the memory location for quote)
			xml.Unmarshal(body, &quote)

			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol)
	}

	// Pert of concurrency implementation (delay termination of the main routine)
	for numComplete < len(stockSymbols) {
		time.Sleep(10 * time.Millisecond)
	}
	elapsed := time.Since(start)

	fmt.Printf("Execution time: %s", elapsed)
}
