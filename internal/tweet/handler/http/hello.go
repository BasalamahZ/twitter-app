package http

import (
	"net/http"

	"github.com/alwisisva/twitter-app/internal/tweet"
)

type helloHandler struct {
	tweetSvc tweet.Service
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO!"))
}
