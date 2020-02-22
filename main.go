package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"rssfeed/controllers"
	rss "rssfeed/getnews"
)

func main() {

	collection := rss.ConnectDB()

	router := mux.NewRouter()
	router.HandleFunc("/person", controllers.SearchRssFeed).Methods("POST")
	go func() {
		log.Println(http.ListenAndServe(GetPort(), router))
	}()

	go rss.StartSpider(collection)
	//http.ListenAndServe(":3000", nil)
	go fmt.Println("I have gotten here")
	select {}

}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT") // Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
