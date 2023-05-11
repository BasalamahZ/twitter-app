package service

type Tweet struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}
type StoreClient interface {
	// TODO: define methods to be used by HTTP handlers to
	// interact with tweet storage (PostgreSQL).
	CreateTweet(tweet *Tweet) (Tweet, error)
	GetAllTweet() ([]Tweet, error)
	GetDetailTweet(id int) (Tweet, error)
}
