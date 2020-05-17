package store

import (
	"context"
	"fmt"

	"github.com/yagossc/price_api/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	return product, nil
}

// FindAllProducts finds all available products
func FindAllProducts(db *mongo.Database) ([]app.Product, error) {
	var results []app.Product

	collection := db.Collection(productsColletion)

	findOptions := options.Find()

	cursor, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return []app.Product{}, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var p app.Product
		err := cursor.Decode(&p)
		if err != nil {
			return []app.Product{}, err
		}

		results = append(results, p)
	}

	if err := cursor.Err(); err != nil {
		return []app.Product{}, err
	}

	return results, nil
}

// InsertProduct inserts a new entry in the "products" collection
func InsertProduct(db *mongo.Database, p app.Product) error {

	collection := db.Collection(productsColletion)

	_, err := FindProductByName(db, p.Name)
	if err == nil {
		return fmt.Errorf("already existent")
	}

	insertResult, err := collection.InsertOne(context.TODO(), p)
	fmt.Printf("Inserted result ID: %v\n", insertResult.InsertedID)
	if err != nil {
		return err
	}

	return nil
}
