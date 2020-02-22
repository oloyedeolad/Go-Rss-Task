package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "net/http/pprof"
	"task1/controllers"
	spider "task1/getnews"
)



func main() {

	collection := spider.ConnectDB()
	router := mux.NewRouter()
	router.HandleFunc("/person", controllers.SearchRssFeed).Methods("POST")
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", router))
	}()

	go spider.StartSpider(collection)
	//http.ListenAndServe(":3000", nil)
	go fmt.Println("I have gotten here")
	select {}

}
