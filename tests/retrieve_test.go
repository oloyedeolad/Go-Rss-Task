package tests

import (
	"github.com/ungerik/go-rss"
	rss2 "rssfeed/getnews"
	"testing"
)

func TestGetRss(t *testing.T) {

	c := make(chan rss.Channel, 100)
	url := "http://rss.cnn.com/rss/edition_world.rss"

	p, _ := rss2.GetRss(c, url)

	if len(p) < 1 {
		t.Errorf("The length should be %d", len(p))
	}
}

func TestGetRssWrongUrl(t *testing.T) {
	c := make(chan rss.Channel, 100)
	url := "cnn.com/rss/edition_world.rss"

	_, err := rss2.GetRss(c, url)

	if err == nil {
		t.Error("wrong url not detected")
	}

}

func TestReceiveFromChannel(t *testing.T) {
	c := make(chan rss.Channel, 100)
	url := "http://rss.cnn.com/rss/edition_world.rss"

	p, _ := rss2.GetRss(c, url)

	feeds := rss2.ReceiveFromChannel(p)

	if feeds == nil {
		t.Error("feeds returned empty")
	}
	if len(feeds) < 1 {
		t.Error("feeds returned empty")
	}
}

func TestSpider(t *testing.T) {

	c := rss2.Spider()

	if !c {
		t.Error("success")
	}
}
