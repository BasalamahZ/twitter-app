package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/BasalamahZ/twitter-app/helper/httplib"
	"github.com/BasalamahZ/twitter-app/internal/tweet"
)

type tweetsHandler struct {
	tweet tweet.Service
}

func (h *tweetsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGetAllTweets(w, r)
	case http.MethodPost:
		h.handleCreateTweet(w, r)
	default:
	}
}

func (h *tweetsHandler) handleGetAllTweets(w http.ResponseWriter, r *http.Request) {
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
			log.Printf("[tweet HTTP][handleGetAllTweets] Failed to get tweets. Source: %s, Err: %s\n", source, err.Error())
			httplib.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		// success
		httplib.WriteResponse(w, resBody, statusCode, httplib.JSONContentTypeDecorator)
	}()

	// prepare channels for main go routine
	resChan := make(chan []tweet.Tweet, 1)
	errChan := make(chan error, 1)

	go func() {
		// get result tweet
		res, err := h.tweet.GetAllTweets(ctx)
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
				log.Printf("[tweet HTTP][handleGetAllTweets] Internal error from handleGetAllTweets. Err: %s\n", err.Error())
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
		// format each tweet
		studyPrograms := make([]tweetHTTP, 0)
		for _, r := range res {
			var sp tweetHTTP
			sp, err = formatTweet(r)
			if err != nil {
				return
			}
			studyPrograms = append(studyPrograms, sp)
		}

		// construct response data
		resBody, err = json.Marshal(httplib.ResponseEnvelope{
			Data: studyPrograms,
		})
	}
}

func (h *tweetsHandler) handleCreateTweet(w http.ResponseWriter, r *http.Request) {
	// add timeout to context
	ctx, cancel := context.WithTimeout(r.Context(), 3000*time.Millisecond)
	defer cancel()

	var (
		err        error           // stores error in this handler
		source     string          // stores request source
		resBody    []byte          // stores response body to write
		statusCode = http.StatusOK // stores response status code
	)

	// write response
	defer func() {
		// ctx = contextlib.SetHTTPStatusCode(ctx, statusCode)
		// *r = *(r.WithContext(ctx))
		// error
		if err != nil {
			log.Printf("[tweet HTTP][handleCreateTweet] Failed to create tweets. Source: %s, Err: %s\n", source, err.Error())
			httplib.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		// success
		httplib.WriteResponse(w, resBody, statusCode, httplib.JSONContentTypeDecorator)
	}()

	// prepare channels for main go routine
	resChan := make(chan int64, 1)
	errChan := make(chan error, 1)

	go func() {
		// read body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- errBadRequest
			return
		}

		// unmarshall body
		request := tweetHTTP{}
		err = json.Unmarshal(body, &request)
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- errBadRequest
			return
		}

		// format HTTP request into service object
		reqTweet, err := parseTweetFromCreateRequest(request)
		if err != nil {
			statusCode = http.StatusBadRequest
			errChan <- err
			return
		}

		tweetID, err := h.tweet.CreateTweet(ctx, reqTweet)
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
				log.Printf("[tweet HTTP][handleCreateTweet] Internal error from CreateTweet. Err: %s\n", err.Error())
			}

			errChan <- parsedErr
			return
		}

		resChan <- tweetID
	}()

	// wait and handle main go routine
	select {
	case <-ctx.Done():
		statusCode = http.StatusGatewayTimeout
		err = errRequestTimeout
	case err = <-errChan:
	case tweetID := <-resChan:
		resBody, err = json.Marshal(httplib.ResponseEnvelope{
			Data: tweetID,
		})
	}
}

func parseTweetFromCreateRequest(th tweetHTTP) (tweet.Tweet, error) {
	result := tweet.Tweet{}

	if th.Title != nil {
		result.Title = *th.Title
	}

	if th.Description != nil {
		result.Description = *th.Description
	}

	return result, nil
}
