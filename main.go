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

	router := mux.NewRouter()
	router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	router.HandleFunc("/search", controllers.SearchRssFeed).Methods("POST")
	go func() {
		log.Println(http.ListenAndServe(GetPort(), router))
	}()

	go rss.StartSpider()

	// This meant to allow the program run forever
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
