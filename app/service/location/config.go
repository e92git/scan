package locationService

import (
	// "database/sql"
	// "scan/app/service/location/repository"
	"scan/app/store"
)

type Config struct {
	store 	 *store.Store
	// repository *repository.Config
}

func New(store *store.Store) *Config {
	return &Config{
		store: store,
		// repository: repository.New(store),
	}
}