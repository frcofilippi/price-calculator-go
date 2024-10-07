package prices

import (
	"fmt"

	"frcofilippi.com/price-tax-calculator/conversions"
	"frcofilippi.com/price-tax-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

func NewTaxPriceIncludedJob(taxRate float64, ioManager iomanager.IOManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		IOManager: ioManager,
	}
}

func (job TaxIncludedPriceJob) Process() error {

	err := job.loadPrices()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", calculateFinalPrice(price, job.TaxRate))
	}

	job.TaxIncludedPrices = result

	return job.saveData()

}

func (job TaxIncludedPriceJob) saveData() error {
	return job.IOManager.WriteJSON(job)
}

func (job *TaxIncludedPriceJob) loadPrices() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversions.StringsToFloat(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func calculateFinalPrice(price, taxRate float64) float64 {
	return price + (price * taxRate / 100)
}
