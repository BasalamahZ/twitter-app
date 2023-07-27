package service

import (
	"context"

	"github.com/BasalamahZ/twitter-app/internal/tweet"
)

// PGStore is the PostgreSQL store for configuration service.
type PGStore interface {
	NewClient(useTx bool) (PGStoreClient, error)
}

type PGStoreClient interface {
	// Commit commits the transaction.
	Commit() error
	// Rollback aborts the transaction.
	Rollback() error
	
	CreateTweet(ctx context.Context, tweet tweet.Tweet) (int64, error)
	GetAllTweets(ctx context.Context) ([]tweet.Tweet, error)
	GetTweetByID(ctx context.Context, id int64) (tweet.Tweet, error)
}
