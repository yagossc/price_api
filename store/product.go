package store

import (
	"context"
	"fmt"

	"github.com/yagossc/price_api/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	productsColletion = "products"
	Name              = "name"
)

// FindProductByName finds an entry in the "products" colletion by filter "name"
func FindProductByName(db *mongo.Database, name string) (app.Product, error) {
	var product app.Product

	collection := db.Collection(productsColletion)

	err := collection.FindOne(context.TODO(), bson.D{{Name, name}}).Decode(&product)
	if err != nil {
		return app.Product{}, err
	}

	fmt.Printf("Found document: %v\n", product)

	return product, nil
}

// InsertProduct inserts a new entry in the "products" collection
func InsertProduct(db *mongo.Database, p app.Product) error {

	collection := db.Collection(productsColletion)

	insertResult, err := collection.InsertOne(context.TODO(), p)
	fmt.Printf("Inserted result ID: %v\n", insertResult.InsertedID)
	if err != nil {
		return err
	}

	return nil

}
