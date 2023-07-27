package postgresql

import (
	"context"
	"fmt"

	"github.com/BasalamahZ/twitter-app/internal/tweet"
	"github.com/jmoiron/sqlx"
)

func (sc *storeClient) CreateTweet(ctx context.Context, tweet tweet.Tweet) (int64, error) {
	// construct arguments filled with fields for the query
	argsKV := map[string]interface{}{
		"title":       tweet.Title,
		"description": tweet.Description,
		"create_time": tweet.CreateTime,
	}

	// prepare query
	query, args, err := sqlx.Named(queryCreateTweet, argsKV)
	if err != nil {
		return 0, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return 0, err
	}
	query = sc.q.Rebind(query)

	// execute query
	var tweetID int64
	err = sc.q.QueryRowx(query, args...).Scan(&tweetID)
	if err != nil {
		return 0, err
	}

	return tweetID, nil
}

func (sc *storeClient) GetAllTweets(ctx context.Context) ([]tweet.Tweet, error) {
	query := fmt.Sprintf(queryGetTweet, "")

	// prepare query
	query, args, err := sqlx.Named(query, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	query = sc.q.Rebind(query)

	// query to database
	rows, err := sc.q.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// read study program
	tweets := make([]tweet.Tweet, 0)
	for rows.Next() {
		var row tweetDB
		err = rows.StructScan(&row)
		if err != nil {
			return nil, err
		}

		tweets = append(tweets, row.format())
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tweets, nil
}

func (sc *storeClient) GetTweetByID(ctx context.Context, id int64) (tweet.Tweet, error) {
	query := fmt.Sprintf(queryGetTweet, "WHERE t.id = $1")

	// query single row
	var tdb tweetDB
	err := sc.q.QueryRowx(query, id).StructScan(&tdb)
	if err != nil {
		return tweet.Tweet{}, err
	}

	return tdb.format(), nil
}
