package repositories

import (
	"context"
	"fmt"
	"github.com/ungerik/go-rss"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"rssfeed/datapack"
)

//The repository for the Rest API
type RssRepository interface {
	List(query string) ([]*rss.Item, error)
	SaveToDb() (*mongo.InsertManyResult, error)
}

var collection *mongo.Collection

func init() {
	collection = datapack.GetCollection()
}

//The interface implementation for the repository
func List(query string) ([]*rss.Item, error) {
	var results []*rss.Item
	searchQuery := bson.M{
		"$text": bson.M{
			"$search": query,
		},
	}
	findOptions := options.Find()
	findOptions.SetLimit(100)
	findOptions.SetAllowPartialResults(true)
	findOptions.SetProjection(bson.M{

		"score": bson.M{"$meta": "textScore"},
	})
	findOptions.SetSort(bson.M{"score": bson.M{"$meta": "textScore"}})
	list, err := collection.Find(context.Background(), searchQuery, findOptions)
	if err != nil {
		return nil, err
	}
	for list.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem rss.Item
		err := list.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}
	fmt.Println(results)
	return results, nil
}

//Methods to save in the database
func SaveToDb(feeds []interface{}) (*mongo.InsertManyResult, error) {

	var opt options.InsertManyOptions
	opt.SetOrdered(false)
	insertManyResult, err := collection.InsertMany(context.Background(), feeds, &opt)

	if err != nil {
		/*fmt.Println(err)*/
	}

	return insertManyResult, nil

}
