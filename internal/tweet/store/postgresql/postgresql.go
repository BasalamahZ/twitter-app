package postgresql

import "github.com/jmoiron/sqlx"

// storeClient implements service.StoreClient.
type storeClient struct {
	db *sqlx.DB
}

// New constructs a new storeClient
func New(connectionString string) (*storeClient, error) {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &storeClient{
		db: db,
	}, nil
}

// TODO: implements service.StoreClient with storeClient.
