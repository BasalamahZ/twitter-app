package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/BasalamahZ/twitter-app/helper/httplib"
	"github.com/BasalamahZ/twitter-app/internal/tweet"
	"github.com/gorilla/mux"
)

type tweetHandler struct {
	tweet tweet.Service
}

func (h *tweetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tweetID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Printf("[Tweet HTTP][TweetHandler] Failed to parse Tweet ID. ID: %s. Err: %s\n", vars["id"], err.Error())
		httplib.WriteErrorResponse(w, http.StatusBadRequest, []string{err.Error()})
		return
	}
	switch r.Method {
	case http.MethodGet:
		h.handleGetTweetByID(w, r, tweetID)
	default:
	}
}

func (h *tweetHandler) handleGetTweetByID(w http.ResponseWriter, r *http.Request, tweetID int64) {
	// add timeout to context
	ctx, cancel := context.WithTimeout(r.Context(), 2000*time.Millisecond)
	defer cancel()

	var (
		err        error           // stores error in this handler
		source     string          // stores request source
		resBody    []byte          // stores response body to write
		statusCode = http.StatusOK // stores response status code
	)

	// write response
	defer func() {
		// error
		if err != nil {
			log.Printf("[Tweet HTTP][handleGetTweetByID] Failed to get Tweet by ID. tweetID: %d. Source: %s, Err: %s\n", tweetID, source, err.Error())
			httplib.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		// success
		httplib.WriteResponse(w, resBody, statusCode, httplib.JSONContentTypeDecorator)
	}()

	// prepare channels for main go routine
	resChan := make(chan tweet.Tweet, 1)
	errChan := make(chan error, 1)

	go func() {
		res, err := h.tweet.GetTweetByID(ctx, tweetID)
		if err != nil {
			// determine error and status code, by default its internal error
			parsedErr := errInternalServer
			statusCode = http.StatusInternalServerError
			if v, ok := mapHTTPError[err]; ok {
				parsedErr = v
				statusCode = http.StatusBadRequest
			}

			// log the actual error if its internal error
			if statusCode == http.StatusInternalServerError {
				log.Printf("[Tweet HTTP][handleGetTweetByID] Internal error from GetTweetByID. tweetID: %d. Err: %s\n", tweetID, err.Error())
			}

			errChan <- parsedErr
			return
		}

		resChan <- res
	}()

	// wait and handle main go routine
	select {
	case <-ctx.Done():
		statusCode = http.StatusGatewayTimeout
		err = errRequestTimeout
	case err = <-errChan:
	case res := <-resChan:
		// format Tweet
		var t tweetHTTP
		t, err = formatTweet(res)
		if err != nil {
			return
		}

		// construct response data
		resBody, err = json.Marshal(httplib.ResponseEnvelope{
			Data: t,
		})
	}
}
