package tweet

import (
	"context"
	"time"
)

type Service interface {
	CreateTweet(ctx context.Context, tweet Tweet) (int64, error)
	GetAllTweets(ctx context.Context) ([]Tweet, error)
	GetTweetByID(ctx context.Context, id int64) (Tweet, error)
}

type Tweet struct {
	ID          int64
	Title       string
	Description string
	CreateBy    string
	CreateTime  time.Time
	UpdateBy    string
	UpdateTime  time.Time
}
