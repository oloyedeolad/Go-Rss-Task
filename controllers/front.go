package controllers

import "net/http"

func FrontControllers() {


	http.HandleFunc("/users", SearchRssFeed)
	http.HandleFunc("/users/", SearchRssFeed)
}
