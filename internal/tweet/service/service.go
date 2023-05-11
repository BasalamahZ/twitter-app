package service

// service implements tweet.Service.
type service struct {
	storeClient StoreClient
}

// New construts a new service.
func New(storeClient StoreClient) *service {
	s := &service{
		storeClient: storeClient,
	}

	return s
}

// TODO: implements tweet.Service with service.
// TODO: define methods to be used by HTTP handlers to
// interact with tweet functionalities.
func (s *service) CreateTweet(tweet *Tweet) (Tweet, error) {
	t, err := s.storeClient.CreateTweet(tweet)
	if err != nil {
		return t, err
	}
	return t, nil
}

func (s *service) GetAllTweet() ([]Tweet, error) {
	t, err := s.storeClient.GetAllTweet()
	if err != nil {
		return t, err
	}
	return t, nil
}

func (s *service) GetDetailTweet(id int) (Tweet, error) {
	t, err := s.storeClient.GetDetailTweet(id)
	if err != nil {
		return t, err
	}
	return t, nil
}
