package main

import (
	"fmt"

	"frcofilippi.com/price-tax-calculator/commandmanager"
	"frcofilippi.com/price-tax-calculator/filemanager"
	"frcofilippi.com/price-tax-calculator/prices"
)

func main() {

	taxes := []float64{10, 20, 30, 50, 70, 80}

	for _, tax := range taxes {
		manager := filemanager.New("pricess.txt", fmt.Sprintf("Result_%0.f.json", tax))
		priceJob := prices.NewTaxPriceIncludedJob(tax, manager)
		manager2 := commandmanager.New()
		priceJob2 := prices.NewTaxPriceIncludedJob(tax, manager2)

		err := priceJob2.Process()

		if err != nil {
			fmt.Println("Error processing")
			fmt.Print(err)
			return
		}
		err = priceJob.Process()

		if err != nil {
			fmt.Println("Error processing")
			fmt.Print(err)
			return
		}

	}
}
