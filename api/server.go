package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server is a wrapper around echo.Echo and the database connection
type Server struct {
	db *mongo.Client
	e  *echo.Echo
}

// NewServer returns a new server instance
func NewServer(db *mongo.Client, e *echo.Echo) *Server {
	s := &Server{db: db, e: e}
	return s
}

// Start starts the server instance.
func (s *Server) Start(address string) error {
	return s.e.Start(address)
}

// ServeHTTP implements `http.Handler` interface, which serves HTTP requests.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.e.ServeHTTP(w, r)
}
