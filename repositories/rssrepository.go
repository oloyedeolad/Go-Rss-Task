package repositories

import (
	"context"
	"fmt"
	"github.com/ungerik/go-rss"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	spider "rssfeed/getnews"
)

//The repository for the Rest API
type RssRepository interface {
	List( query string) ([]*rss.Item, error)
	Get(ID string, dest interface{}, query string) (rss.Item, error)
}


//The interface implementation for the repository
func List( query string) ([]*rss.Item, error) {
	var results []*rss.Item
	collection := spider.ConnectDB()
	fmt.Println(query)
	searchQuery := bson.M{
		"$text": bson.M{
			"$search": query,
		},
	}
	findOptions := options.Find()
	findOptions.SetLimit(100)
	findOptions.SetAllowPartialResults(true)
	findOptions.SetProjection(bson.M{

		"score":       bson.M{"$meta": "textScore"},
	})
	findOptions.SetSort(bson.M{"score": bson.M{"$meta": "textScore"}})
	list, err := collection.Find(context.Background(), searchQuery, findOptions)
	//defer list.Close(context.Background())
	if err != nil{
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
	return results,  nil
}

//method to get a single rssfeed
func Get(ID string, dest interface{}, query string) (rss.Item, error) {
	collection := spider.ConnectDB()
	var result rss.Item;
	 err := collection.FindOne(context.Background(), ID).Decode(&result)

	if err != nil {
		return result, err
	}

	 return result, nil
}
