package internal

import (
	"fmt"
	"math"
	"strconv"

	"github.com/yagossc/price_api/app"
)

type priceModality struct {
	n        uint
	interest float64
}

var modalities = []priceModality{
	{n: 36, interest: 0.0146},
	{n: 60, interest: 0.0145},
	{n: 72, interest: 0.0144},
	{n: 120, interest: 0.0139},
}

// Each payment value based on the Price method is given by:
// Total*(Interest/1 - (1 + Interest)^N).
// Here we call 1 - (1 + Interest)^N), 'Denominator'
func GetPriceTable(value float64) []app.Pricing {

	var pricePlans []app.Pricing

	for _, modality := range modalities {

		// calculate pricing table entry
		denominator := getDenominator(modality.interest, modality.n)
		tableValue := value * modality.interest / denominator
		interestString := strconv.FormatFloat(modality.interest*100, 'f', 2, 64)

		pricePlans = append(pricePlans, app.Pricing{
			N:          modality.n,
			NValue:     strconv.FormatFloat(tableValue, 'f', 2, 64),
			Interest:   fmt.Sprintf("%s%%a.m", interestString),
			TotalValue: strconv.FormatFloat(tableValue*float64(modality.n), 'f', 2, 64),
		})

	}

	fmt.Printf("Return table: %v\n", pricePlans)
	return pricePlans
}

func getDenominator(i float64, n uint) float64 {
	d1 := math.Pow((1 + (i)), float64(n))
	d2 := 1 / d1
	return float64(1 - d2)
}
