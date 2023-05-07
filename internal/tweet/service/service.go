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
