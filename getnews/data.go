package rss

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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

	//create text index
	opt := options.Index()
	opt.SetUnique(true)
	opt.SetWeights(bson.M{
		"title":       5, // Word matches in the title are weighted 5× standard.
		"description": 2,
	})

	index := mongo.IndexModel{Keys: bson.M{
		"title":       "text",
		"description": "text",
		"body":        "text",
	}, Options: opt}
	if _, err := collection.Indexes().CreateOne(context.Background(), index); err != nil {
		log.Println("Could not create text index:", err)
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
