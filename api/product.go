package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yagossc/price_api/app"
	"github.com/yagossc/price_api/internal"
	"github.com/yagossc/price_api/store"
)

func (s *Server) quotation(c echo.Context) error {
	var err error
	var totalPrice float64
	var quoteRequest []app.QuotationRequest

	if err = c.Bind(&quoteRequest); err != nil {
		return err
	}

	for _, quote := range quoteRequest {
		p, err := store.FindProductByName(s.storage.database, quote.ProductName)
		if err != nil {
			return err
		}

		f, err := strconv.ParseFloat(p.Price, 64)
		if err != nil {
			return err
		}

		totalPrice += f
	}

	priceTable := internal.GetPriceTable(totalPrice)

	return c.JSON(http.StatusOK, priceTable)
}

func (s *Server) getAllProducts(c echo.Context) error {
	res, err := store.FindAllProducts(s.storage.database)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
