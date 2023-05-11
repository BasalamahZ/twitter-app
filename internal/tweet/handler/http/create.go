package http

import (
	"encoding/json"
	"net/http"

	"github.com/alwisisva/twitter-app/internal/tweet"
	"github.com/alwisisva/twitter-app/internal/tweet/service"
)

type createHandler struct {
	tweetSvc tweet.Service
}

func (h *createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		p := service.Tweet{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		tweet, err := h.tweetSvc.CreateTweet(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, _ := json.Marshal(tweet)

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}
