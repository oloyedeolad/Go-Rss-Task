package rss

import (
	"github.com/ungerik/go-rss"
	"testing"
)

func TestGetRss(t *testing.T) {

	c := make(chan rss.Channel, 100)
	url := "http://rss.cnn.com/rss/edition_world.rss"

	p, _ := GetRss(c, url)

	if len(p) < 1 {
		t.Errorf("The length should be %d", len(p))
	}
}

func TestGetRssWrongUrl(t *testing.T)  {
	c := make(chan rss.Channel, 100)
	url := "cnn.com/rss/edition_world.rss"

	_, err := GetRss(c, url)

	if err == nil {
		t.Error("wrong url not detected")
	}

}

func TestReceiveFromChannel(t *testing.T) {
	c := make(chan rss.Channel, 100)
	url := "http://rss.cnn.com/rss/edition_world.rss"

	p, _ := GetRss(c, url)

	feeds := ReceiveFromChannel(p)

	if feeds == nil{
		t.Error("feeds returned empty")
	}
	if len(feeds) < 1 {
		t.Error("feeds returned empty")
	}
}


func TestSpider(t *testing.T) {
	collection := ConnectDB()
	c := Spider(collection)

	if !c {
		t.Error("success")
	}
}
