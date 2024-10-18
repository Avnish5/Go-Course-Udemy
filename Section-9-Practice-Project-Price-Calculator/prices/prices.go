package prices

import (
	"fmt"

	"example.com/price-calculator-app/conversion"
	"example.com/price-calculator-app/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_price"`
}

func (job *TaxIncludedPriceJob) Loaddata() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil

}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.Loaddata()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		temp := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", temp)
	}

	job.TaxIncludedPrices = result
	return job.IOManager.WriteJson(job)
}

func NewTaxIncludedPriceJob(io iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		IOManager:   io,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
