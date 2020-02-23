package controllers

import (
	"encoding/json"
	"net/http"
	"rssfeed/repositories"
)

type Topic struct {
	Topic string `json:"topic" bson:"_id,omitempty"`
}

func SearchRssFeed(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var searchText Topic

	_ = json.NewDecoder(request.Body).Decode(&searchText)
	if searchText.Topic == "" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + "Ensure your request has a topic" + `" }`))
		return
	}

	items, err := repositories.List(searchText.Topic)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(items)
}
