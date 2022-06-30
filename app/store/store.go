package store

import (
	"gorm.io/gorm"
)

// Store ...
type Store struct {
	db             		*gorm.DB
	userRepository 		*UserRepository
	locationRepository  *LocationRepository
	scanRepository  	*ScanRepository
}

// New ...
func New(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
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

// Location ...
func (s *Store) Scan() *ScanRepository {
	if s.scanRepository != nil {
		return s.scanRepository
	}

	s.scanRepository = &ScanRepository{
		store: s,
	}

	return s.scanRepository
}