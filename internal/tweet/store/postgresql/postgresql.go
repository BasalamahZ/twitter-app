package postgresql

import (
	"fmt"
	"log"

	"github.com/alwisisva/twitter-app/internal/tweet/service"
	"github.com/jmoiron/sqlx"
)

// storeClient implements service.StoreClient.
type storeClient struct {
	db *sqlx.DB
}

// New constructs a new storeClient
func New(connectionString string) (*storeClient, error) {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Println(err.Error())
	}

	return &storeClient{
		db: db,
	}, nil
}

// TODO: implements service.StoreClient with storeClient.
// TODO: define methods to be used by HTTP handlers to
// interact with tweet storage (PostgreSQL).
func (s *storeClient) CreateTweet(tweet *service.Tweet) (service.Tweet, error) {
	if err := s.db.Get(tweet, `INSERT INTO tweets VALUES ($1, $2, $3) RETURNING *`,
		tweet.ID,
		tweet.Title,
		tweet.Description); err != nil {
		return service.Tweet{}, fmt.Errorf("error getting tweet: %w", err)
	}
	return *tweet, nil
}

func (s *storeClient) GetAllTweet() ([]service.Tweet, error) {
	var tweet []service.Tweet
	if err := s.db.Select(&tweet, `SELECT * FROM tweets`); err != nil {
		return []service.Tweet{}, fmt.Errorf("error getting tweet: %w", err)
	}
	return tweet, nil
}

func (s *storeClient) GetDetailTweet(id int) (service.Tweet, error) {
	var tweet service.Tweet
	if err := s.db.Get(&tweet, `SELECT * FROM tweets WHERE id = $1`, id); err != nil {
		return service.Tweet{}, fmt.Errorf("error getting tweet: %w", err)
	}
	return tweet, nil
}
