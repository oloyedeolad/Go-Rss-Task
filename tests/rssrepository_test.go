package tests

import (
	"github.com/ungerik/go-rss"
	rss2 "rssfeed/getnews"
	"rssfeed/repositories"
	"testing"
)

func TestSaveToDb(t *testing.T) {

	c := make(chan rss.Channel, 100)
	url := "http://rss.cnn.com/rss/edition_world.rss"

	p, _ := rss2.GetRss(c, url)

	feeds := rss2.ReceiveFromChannel(p)

	sv, _ := repositories.SaveToDb(feeds)

	if sv == nil {
		t.Error("save failed")
	}
}

func TestSaveToDbNoDuplicate(t *testing.T) {

	c := make(chan rss.Channel, 100)
	url := "http://rss.cnn.com/rss/edition_world.rss"

	p, _ := rss2.GetRss(c, url)
	feeds := rss2.ReceiveFromChannel(p)

	_, err := repositories.SaveToDb(feeds)

	if err != nil {
		t.Error("save failed")
	}
}
