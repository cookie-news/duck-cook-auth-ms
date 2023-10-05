package mongo

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	COLLETCTION_CUSTOMER = "Customer"
)

func Connect() mongo.Database {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_SRV")).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	db := mongoClient.Database(os.Getenv("MONGO_DB"))

	collection := db.Collection(COLLETCTION_CUSTOMER)

	indexUser := mongo.IndexModel{
		Keys:    bson.M{"user": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateOne(context.TODO(), indexUser)
	if err != nil {
		panic(err)
	}

	indexEmail := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateOne(context.TODO(), indexEmail)
	if err != nil {
		panic(err)
	}

	return *db
}
