package prices

import (
	"example.com/price-calculator-app/conversion"
	"example.com/price-calculator-app/filemanager"
	"fmt"
)

const fileName = "prices.txt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) Loaddata() {

	lines, err := filemanager.ReadLines(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices

}

func (job *TaxIncludedPriceJob) Process() {
	job.Loaddata()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		temp := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", temp)
	}

	job.TaxIncludedPrices = result
	filemanager.WriteJson(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
