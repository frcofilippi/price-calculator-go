package main

import (
	"fmt"

	"frcofilippi.com/price-tax-calculator/filemanager"
	"frcofilippi.com/price-tax-calculator/prices"
)

func main() {

	taxes := []float64{10, 20, 30, 50, 70, 80}

	doneChans := make([]chan bool, len(taxes))
	errorChans := make([]chan error, len(taxes))
	for tIndex, tax := range taxes {
		manager := filemanager.New("prices.txt", fmt.Sprintf("Result_%0.f.json", tax))
		priceJob := prices.NewTaxPriceIncludedJob(tax, manager)
		doneChans[tIndex] = make(chan bool)
		errorChans[tIndex] = make(chan error)
		go priceJob.Process(doneChans[tIndex], errorChans[tIndex])

	}

	for index := range taxes {
		select {
		case err := <-errorChans[index]:
			fmt.Println(err)
		case <-doneChans[index]:
			fmt.Printf("T: %v > Task completed\n`", index)
		}
	}
}
