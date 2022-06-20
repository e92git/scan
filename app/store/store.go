package store

import (
	"database/sql"
)

// Store ...
type Store struct {
	db             		*sql.DB
	userRepository 		*UserRepository
	locationRepository 	*LocationRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Location ...
func (s *Store) Location() *LocationRepository {
	if s.locationRepository != nil {
		return s.locationRepository
	}

	s.locationRepository = &LocationRepository{
		store: s,
	}

	return s.locationRepository
}

// User ...
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}