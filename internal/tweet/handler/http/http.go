package http

import (
	"errors"
	"net/http"

	"github.com/BasalamahZ/twitter-app/internal/tweet"
	"github.com/gorilla/mux"
)

var (
	errUnknownConfig = errors.New("unknown config name")
)

// Handler contains tweet HTTP-handlers.
type Handler struct {
	handlers map[string]*handler
	tweet    tweet.Service
}

// handler is the HTTP handler wrapper.
type handler struct {
	h        http.Handler
	identity HandlerIdentity
}

// HandlerIdentity denotes the identity of an HTTP hanlder.
type HandlerIdentity struct {
	Name string
	URL  string
}

// Followings are the known HTTP handler identities
var (
	// HandlerTweets denotes HTTP handler to interact
	// with a student group.
	HandlerTweets = HandlerIdentity{
		Name: "tweets",
		URL:  "/tweets",
	}

	HandlerTweet = HandlerIdentity{
		Name: "tweet",
		URL:  "/tweets/{id}",
	}
)

// Option controls the behavior of Handler.
type Option func(*Handler) error

// WithHandler returns Option to add HTTP handler.
func WithHandler(identity HandlerIdentity) Option {
	return Option(func(h *Handler) error {

		return nil
	})
}

// New creates a new Handler.
//
// For the given Option, WithScopeSetting() should come first
// before WithHandler()
func New(tweet tweet.Service, identities []HandlerIdentity) (*Handler, error) {
	h := &Handler{
		handlers: make(map[string]*handler),
		tweet:    tweet,
	}

	// apply options
	for _, identity := range identities {
		if h.handlers == nil {
			h.handlers = map[string]*handler{}
		}

		h.handlers[identity.Name] = &handler{
			identity: identity,
		}

		handler, err := h.createHTTPHandler(identity.Name)
		if err != nil {
			return nil, err
		}

		h.handlers[identity.Name].h = handler
	}

	return h, nil
}

// createHTTPHandler creates a new HTTP handler that
// implements http.Handler.
func (h *Handler) createHTTPHandler(configName string) (http.Handler, error) {
	var httpHandler http.Handler
	switch configName {
	case HandlerTweets.Name:
		httpHandler = &tweetsHandler{
			tweet: h.tweet,
		}
	case HandlerTweet.Name:
		httpHandler = &tweetHandler{
			tweet: h.tweet,
		}
	default:
		return httpHandler, errUnknownConfig
	}
	return httpHandler, nil
}

// Start starts all HTTP handlers.
func (h *Handler) Start(multiplexer *mux.Router) error {
	for _, handler := range h.handlers {
		multiplexer.Handle(handler.identity.URL, handler.h)
	}
	return nil
}

type tweetHTTP struct {
	ID          *int64  `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
