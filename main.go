package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	rss "rssfeed/getnews"
)



func main() {

	collection := rss.ConnectDB()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	go rss.StartSpider(collection)
	//http.ListenAndServe(":3000", nil)
	go fmt.Println("I have gotten here")
	select {}

}
