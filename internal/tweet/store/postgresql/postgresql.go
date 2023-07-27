package postgresql

import (
	"errors"
	"time"

	"github.com/BasalamahZ/twitter-app/internal/tweet"
	"github.com/BasalamahZ/twitter-app/internal/tweet/service"
	"github.com/jmoiron/sqlx"
)

var (
	errInvalidCommit   = errors.New("cannot do commit on non-transactional querier")
	errInvalidRollback = errors.New("cannot do rollback on non-transactional querier")
)

// store implements configuration/service.PGStore
type store struct {
	db *sqlx.DB
}

// storeClient implements configuration/service.PGStoreClient
type storeClient struct {
	q sqlx.Ext
}

// New creates a new store.
func New(db *sqlx.DB) (*store, error) {
	s := &store{
		db: db,
	}

	return s, nil
}

func (s *store) NewClient(useTx bool) (service.PGStoreClient, error) {
	var q sqlx.Ext

	// determine what object should be use as querier
	q = s.db
	if useTx {
		var err error
		q, err = s.db.Beginx()
		if err != nil {
			return nil, err
		}
	}

	return &storeClient{
		q: q,
	}, nil
}

func (sc *storeClient) Commit() error {
	if tx, ok := sc.q.(*sqlx.Tx); ok {
		return tx.Commit()
	}
	return errInvalidCommit
}

func (sc *storeClient) Rollback() error {
	if tx, ok := sc.q.(*sqlx.Tx); ok {
		return tx.Rollback()
	}
	return errInvalidRollback
}

type tweetDB struct {
	ID          int64      `db:"id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	CreateTime  time.Time  `db:"create_time"`
	UpdateTime  *time.Time `db:"update_time"`
}

// format formats database struct into domain struct.
func (tdb *tweetDB) format() tweet.Tweet {
	t := tweet.Tweet{
		ID:          tdb.ID,
		Title:       tdb.Title,
		Description: tdb.Description,
		CreateTime:  tdb.CreateTime,
	}

	if tdb.UpdateTime != nil {
		t.UpdateTime = *tdb.UpdateTime
	}

	return t
}
