package http

import (
	"net/http"

	"github.com/alwisisva/twitter-app/internal/tweet"
)

// Handler contains HTTP handlers.
type Handler struct {
	// handlers maps a URL path to its HTTP handler
	handlers map[string]http.Handler
}

// New construct a new Handler.
func New(tweetSvc tweet.Service) *Handler {
	h := &Handler{
		handlers: make(map[string]http.Handler),
	}

	// example how to add HTTP handler to h.handlers
	// h.handlers["/hello"] = &helloHandler{
	// 	tweetSvc: tweetSvc,
	// }

	// TODO: add desired HTTP handler to h.handlers
	h.handlers["/tweets"] = &getAllHandler{
		tweetSvc: tweetSvc,
	}
	
	h.handlers["/tweet"] = &createHandler{
		tweetSvc: tweetSvc,
	}
	h.handlers["/tweet/"] = &getDetailHandler{
		tweetSvc: tweetSvc,
	}

	return h
}

// Start starts all HTTP handlers.
func (h *Handler) Start() {
	for path, handler := range h.handlers {
		http.Handle(path, handler)
	}
}
