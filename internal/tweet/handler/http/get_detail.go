package http

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/alwisisva/twitter-app/internal/tweet"
)

type getDetailHandler struct {
	tweetSvc tweet.Service
}

func (h *getDetailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile(`/([0-9]+)`)
	group := regex.FindAllStringSubmatch(r.URL.Path, -1)

	if len(group) != 1 || len(group[0]) != 2 {
		http.Error(w, "Invalid URI", http.StatusBadRequest)
		return
	}

	idString := group[0][1]
	newId, _ := strconv.Atoi(idString)
	tweet, err := h.tweetSvc.GetDetailTweet(newId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(tweet)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
