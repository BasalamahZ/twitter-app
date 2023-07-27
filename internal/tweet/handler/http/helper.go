package http

import "github.com/BasalamahZ/twitter-app/internal/tweet"

// formatTweet formats the given tweet into the
// respective HTTP-format object.
func formatTweet(t tweet.Tweet) (tweetHTTP, error) {
	return tweetHTTP{
		ID:          &t.ID,
		Title:       &t.Title,
		Description: &t.Description,
	}, nil
}
