package http

import (
	"encoding/json"
	"net/http"

	"github.com/alwisisva/twitter-app/internal/tweet"
)

type getAllHandler struct {
	tweetSvc tweet.Service
}

func (h *getAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tweets, err := h.tweetSvc.GetAllTweet()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		response, _ := json.Marshal(tweets)

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}
