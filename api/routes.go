package api

func (s *Server) Routes() {

	// create quotation route
	s.e.POST("/quotation", s.quotation)

	// create products route
	s.e.GET("/products", s.getAllProducts)
}
