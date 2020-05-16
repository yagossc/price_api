package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// database contains both a mongo Client
// and a mongo Database handler to facilitate
// connecting/disconnecting to a given db
// and accessing colletions
type dbConnection struct {
	client   *mongo.Client
	database *mongo.Database
}

// Server is a wrapper around echo.Echo and the database connection
type Server struct {
	storage *dbConnection
	e       *echo.Echo
}

// NewServer returns a new server instance
func NewServer(cli *mongo.Client, db *mongo.Database, e *echo.Echo) *Server {
	conn := &dbConnection{client: cli, database: db}
	s := &Server{storage: conn, e: e}
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
