package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagossc/price_api/app"
	"github.com/yagossc/price_api/store"
)

func (s *Server) quotation(c echo.Context) error {
	var err error
	var products []app.Product
	var quoteRequest []app.Quotation

	if err = c.Bind(&quoteRequest); err != nil {
		return err
	}

	for _, quote := range quoteRequest {
		p, err := store.FindProductByName(s.storage.database, quote.ProductName)
		if err != nil {
			return err
		}
		products = append(products, p)
	}

	fmt.Printf("%v\n", quoteRequest)
	fmt.Printf("%v\n", products)

	return c.JSON(http.StatusOK, products)

}
