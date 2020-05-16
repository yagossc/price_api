package api

func (s *Server) Routes() {

	// create products route
	s.e.POST("/quotation", s.quotation)
}
