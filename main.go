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
	"github.com/yagossc/price_api/app"
	"github.com/yagossc/price_api/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// load configuration
	cfg := loadConfig()

	// database connection
	cli, db, err := openDBConnection(cfg)
	if err != nil {
		os.Exit(1)
	}

	defer cli.Disconnect(context.TODO())

	// server configuration
	e := echo.New()

	// Create server
	s := api.NewServer(cli, db, e)

	// Setup Routes
	s.Routes()

	seed(cfg, db)

	// Start server
	log.Fatal(s.Start(":" + strconv.FormatUint(cfg.Port, 10)))
}

func openDBConnection(cfg config) (*mongo.Client, *mongo.Database, error) {
	var cli *mongo.Client
	var err error

	// Set client options
	clientOptions := options.Client().ApplyURI(cfg.DBURL)

	for i := 0; i < 5; i++ {
		if i > 0 {
			time.Sleep(time.Duration(i+1) * time.Second)
		}

		fmt.Printf("Connecting to database (tries=%d)... ", i+1)
		cli, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			fmt.Printf("ERROR!\n%v\n\n", err)
			continue
		}
		// Check the connection
		err = cli.Ping(context.TODO(), nil)
		if err != nil {
			fmt.Printf("ERROR!\ndatabase error: %v\n", err)
		} else {
			break
		}
	}

	db := cli.Database(cfg.DBNAME)

	return cli, db, err
}

func seed(cfg config, db *mongo.Database) error {
	initialSeed := []app.Product{
		{Name: "MÓDULO POLI 330W", Price: "657.3798"},
		{Name: "INVERSOR MONO 3KW 220V", Price: "3763.59375"},
		{Name: "CABO SOLAR 6MM2 VM", Price: "5.32"},
		{Name: "DISJUNTOR TRIPOLAR 63A", Price: "23.4348375"},
		{Name: "PERFIS ALUMÍNIO 4150MM TELHADO 04 PLACAS", Price: "144.6571875"},
		{Name: "ESTRUTURA SOLAR TELHA ONDULADA 4 PLACAS", Price: "165.615625"},
	}

	for _, seed := range initialSeed {
		err := store.InsertProduct(db, seed)
		if err != nil {
			fmt.Printf("Error seeding: %v\n", err)
		} else {
			fmt.Printf("Seeded: %v\n", seed)
		}
	}

	return nil
}
