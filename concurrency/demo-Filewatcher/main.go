package main

import (
	"fmt"
	"os"
	"time"
	"io/ioutil"
	"strings"
	"encoding/csv"
	"strconv"
	"runtime"
)

const watchedPath = "./source"

type Invoice struct {
	Number string
	Amount float64
	PurchaseOrderNumber int
	InvoiceDate time.Time
}

func main() {
	runtime.GOMAXPROCS(2) // introducing multiple processors 
	for {
		d, _ := os.Open(watchedPath)
		files, _ := d.Readdir(-1) // return as many files as possible
		for _, fi := range files {
			filePath := watchedPath + "/" + fi.Name()
			f, _ := os.Open(filePath) // Open individual file
			data, _ := ioutil.ReadAll(f) // read content from individual file
			f.Close() // close file
			os.Remove(filePath) // delete file

			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))
				records, _ := reader.ReadAll()
				for _, r := range records {
					invoice := new(Invoice)
					invoice.Number = r[0]
					invoice.Amount, _ = strconv.ParseFloat(r[1], 64)
					invoice.PurchaseOrderNumber, _ = strconv.Atoi(r[2])
					unixTime, _ := strconv.ParseInt(r[3], 10, 64)
					invoice.InvoiceDate = time.Unix(unixTime, 0)

					fmt.Printf("Received invoice '%v' for $%.2f and submitted for processing\n", invoice.Number, invoice.Amount)
				}
			}(string(data))
		}
	}
}