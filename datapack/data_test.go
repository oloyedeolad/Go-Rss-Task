package datapack

import (
	"github.com/ungerik/go-rss"
	rss2 "rssfeed/getnews"
	"testing"
)

func TestConnectDB(t *testing.T) {
	collection := ConnectDB()

	if collection == nil {
		t.Error("collection not available")
	}
}

func TestSaveToDb(t *testing.T) {

	c := make(chan rss.Channel, 100)
	url := "http://rss.cnn.com/rss/edition_world.rss"

	p, _ := rss2.GetRss(c, url)
	collection := ConnectDB()
	feeds := rss2.ReceiveFromChannel(p)

	sv, _ := SaveToDb(feeds, collection)

	if sv == nil {
		t.Error("save failed")
	}
}

func TestSaveToDbNoDuplicate(t *testing.T) {

	c := make(chan rss.Channel, 100)
	url := "http://rss.cnn.com/rss/edition_world.rss"

	p, _ := rss2.GetRss(c, url)
	collection := ConnectDB()
	feeds := rss2.ReceiveFromChannel(p)

	_, err := SaveToDb(feeds, collection)

	if err != nil {
		t.Error("save failed")
	}
}
