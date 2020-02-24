package controllers

import (
	"encoding/json"
	"net/http"
	"rssfeed/repositories"
)

type Topic struct {
	Topic string `json:"topic" bson:"topic,omitempty"`
}

// controller for searching the database. Note: each feed is referred to as Items
func SearchRssFeed(response http.ResponseWriter, request *http.Request) {

	//set the header for the response
	response.Header().Set("content-type", "application/json")
	var searchText Topic
	//Decode the search text
	_ = json.NewDecoder(request.Body).Decode(&searchText)
	// return a bad response for empty topic
	if searchText.Topic == "" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + "Ensure your request has a topic" + `" }`))
		return
	}

	// Use the repository to make the search
	items, err := repositories.List(searchText.Topic)
	// catching any error that might return from the search
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	//Encoding and returning response
	json.NewEncoder(response).Encode(items)
}
