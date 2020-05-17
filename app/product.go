package app

// Product defines a product's attributes
type Product struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

// QuotationRequest defines a
// quotation's request attributes
type QuotationRequest struct {
	ProductName string `json:"name"`
	Quantity    uint   `json:"quant"`
}
