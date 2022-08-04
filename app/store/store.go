package store

import (
	"gorm.io/gorm"
)

// Store ...
type Store struct {
	db                 *gorm.DB
	userRepository     *UserRepository
	locationRepository *LocationRepository
	scanRepository     *ScanRepository
	vinRepository      *VinRepository
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

// Scan ...
func (s *Store) Scan() *ScanRepository {
	if s.scanRepository != nil {
		return s.scanRepository
	}
	s.scanRepository = &ScanRepository{
		store: s,
	}
	return s.scanRepository
}

// Vin ...
func (s *Store) Vin() *VinRepository {
	if s.vinRepository != nil {
		return s.vinRepository
	}
	s.vinRepository = &VinRepository{
		store: s,
	}
	return s.vinRepository
}
