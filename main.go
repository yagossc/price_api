package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/yagossc/price_api/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// load configuration
	cfg := loadConfig()

	// database connection
	db, err := openDBConnection(cfg)
	if err != nil {
		os.Exit(1)
	}

	defer db.Disconnect(context.TODO())

	// server configuration
	e := echo.New()

	// Create server
	s := api.NewServer(db, e)

	// Start server
	log.Fatal(s.Start(":" + strconv.FormatUint(cfg.Port, 10)))
}

func openDBConnection(cfg config) (*mongo.Client, error) {
	var db *mongo.Client
	var err error

	// Set client options
	clientOptions := options.Client().ApplyURI(cfg.DBURL)

	for i := 0; i < 5; i++ {
		if i > 0 {
			time.Sleep(time.Duration(i+1) * time.Second)
		}

		fmt.Printf("Connecting to database (tries=%d)... ", i+1)
		db, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			fmt.Printf("ERROR!\n%v\n\n", err)
			continue
		}
		// Check the connection
		err = db.Ping(context.TODO(), nil)
		if err != nil {
			fmt.Printf("ERROR!\ndatabase error: %v\n", err)
		} else {
			break
		}
	}

	return db, err
}
