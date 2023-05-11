package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"

	tweethttphandler "github.com/alwisisva/twitter-app/internal/tweet/handler/http"
	tweetservice "github.com/alwisisva/twitter-app/internal/tweet/service"
	tweetpgstore "github.com/alwisisva/twitter-app/internal/tweet/store/postgresql"
)

func main() {
	// TODO: change connection string if needed
	storeClient, err := tweetpgstore.New("postgres://postgres:root@localhost:5432/tweetapp?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	service := tweetservice.New(storeClient)

	handler := tweethttphandler.New(service)
	handler.Start()

	// start http server with handlers from http.DefaultServeMux
	http.ListenAndServe(":8080", nil)
}
