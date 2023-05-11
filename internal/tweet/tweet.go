package tweet

import "github.com/alwisisva/twitter-app/internal/tweet/service"

type Service interface {
	// TODO: define methods to be used by HTTP handlers to
	// interact with tweet functionalities.
	CreateTweet(tweet *service.Tweet) (service.Tweet, error)
	GetAllTweet() ([]service.Tweet, error)
	GetDetailTweet(id int) (service.Tweet, error)
}
