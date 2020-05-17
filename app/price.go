package app

// Pricing describes the an entry
// in the pricing table, i.e.:
// N          - number of payments
// NValue     - the value of each payment
// Interest   - the monthly interest rate
// TotalValue - the total value to be payed
type Pricing struct {
	N          uint   `json:"num_parcelas""`
	NValue     string `json:"val_parcela"`
	Interest   string `json:"juros"`
	TotalValue string `json:"pgto_total"`
}
