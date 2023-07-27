package service

import (
	"context"

	"github.com/BasalamahZ/twitter-app/internal/tweet"
)

// TODO: implements tweet.Service with service.
// TODO: define methods to be used by HTTP handlers to
// interact with tweet functionalities.
func (s *service) CreateTweet(ctx context.Context, tweet tweet.Tweet) (int64, error) {
	// these value should be same for all users
	var (
		createTime = s.timeNow()
	)

	tweet.CreateTime = createTime

	// get pg store client without using transaction
	pgStoreClient, err := s.pgStore.NewClient(false)
	if err != nil {
		return 0, err
	}

	tweetID, err := pgStoreClient.CreateTweet(ctx, tweet)
	if err != nil {
		return 0, err
	}
	return tweetID, nil
}

func (s *service) GetAllTweets(ctx context.Context) ([]tweet.Tweet, error) {
	// get pg store client without using transaction
	pgStoreClient, err := s.pgStore.NewClient(false)
	if err != nil {
		return nil, err
	}

	// get user from pgstore
	result, err := pgStoreClient.GetAllTweets(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) GetTweetByID(ctx context.Context, id int64) (tweet.Tweet, error) {
	// validate arguments
	if id <= 0 {
		return tweet.Tweet{}, tweet.ErrDataNotFound
	}

	// get pg store client without using transaction
	pgStoreClient, err := s.pgStore.NewClient(false)
	if err != nil {
		return tweet.Tweet{}, err
	}

	// get user from pgstore
	result, err := pgStoreClient.GetTweetByID(ctx, id)
	if err != nil {
		return tweet.Tweet{}, err
	}

	return result, nil
}
