package database

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"shipping-management/internal/infrastructure/adapters/database"
	"shipping-management/internal/infrastructure/config"
	"shipping-management/internal/infrastructure/repositories/dtos"
)

func SeedCarriers() {
	ctx := context.Background()
	appConfig := config.NewConfig()

	mongoDB, err := database.NewMongoDatabase(ctx, appConfig)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = mongoDB.Client().Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	jsonPath := filepath.Join("internal", "infrastructure", "database", "carriers.json")
	jsonData, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("Error to read json: %v", err)
	}

	var carriers []dtos.Carrier
	if err := json.Unmarshal(jsonData, &carriers); err != nil {
		log.Fatalf("Error: %v", err)
	}

	documents := make([]interface{}, len(carriers))
	for i, carrier := range carriers {
		documents[i] = carrier
	}

	collection := mongoDB.Collection("carriers")

	if _, err := collection.DeleteMany(ctx, map[string]interface{}{}); err != nil {
		log.Printf("Canot clean collection: %v", err)
	}

	_, err = collection.InsertMany(ctx, documents)
	if err != nil {
		log.Fatalf("Error to insert data: %v", err)
	}

	log.Printf("Inserted carriers into the collection")
}
