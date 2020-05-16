package app

// Product defines a product's attributes
type Product struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

// Quotation defines a quotation's
// request attributes
type Quotation struct {
	ProductName string `json:"name"`
	Quantity    uint   `json:"quant"`
}
