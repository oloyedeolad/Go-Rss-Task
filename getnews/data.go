package rss

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//This is the method for connecting to the mongodb database
func ConnectDB() *mongo.Collection {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://feed:CvlSll83tNW5vpVD@cluster0-5nktq.mongodb.net/test?retryWrites=true&w=majority")
	clientOptions = clientOptions.SetMaxPoolSize(50)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("Feeds").Collection("Items")
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{
		Keys: bsonx.Doc{{Key: "title", Value: bsonx.String("text")}},

		Options: opt,
	}

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)

	_, err = collection.Indexes().CreateOne(context.Background(), index, opts)
	if err != nil {
		log.Println(err.Error())
	}
	return collection

}

//This method saves the received news into th database and also prevent duplication
func SaveToDb(feeds []interface{}, collection *mongo.Collection) (*mongo.InsertManyResult, error) {

	var opt options.InsertManyOptions
	opt.SetOrdered(false)
	insertManyResult, err := collection.InsertMany(context.Background(), feeds, &opt)

	if err != nil {
		/*fmt.Println(err)*/
	}

	return insertManyResult, nil
	//fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}
