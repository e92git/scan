package store

import (
	"scan/app/apiserver"

	"gorm.io/gorm"
)

// Store ...
type Store struct {
	db                 *gorm.DB
	userRepository     *UserRepository
	locationRepository *LocationRepository
	scanRepository     *ScanRepository
	vinRepository      *VinRepository
	carMarkRepository  *CarMarkRepository
	carModelRepository *CarModelRepository
	tireRepository     *TireRepository
}

// New ...
func New(config *apiserver.Config) (*Store, error) {
	db, err := apiserver.ConnectGorm(config.Dsn, config.LogLevel)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
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

// CarMark ...
func (s *Store) CarMark() *CarMarkRepository {
	if s.carMarkRepository != nil {
		return s.carMarkRepository
	}
	s.carMarkRepository = &CarMarkRepository{
		store: s,
	}
	return s.carMarkRepository
}

// CarModel ...
func (s *Store) CarModel() *CarModelRepository {
	if s.carModelRepository != nil {
		return s.carModelRepository
	}
	s.carModelRepository = &CarModelRepository{
		store: s,
	}
	return s.carModelRepository
}
// Tire ...
func (s *Store) Tire() *TireRepository {
	if s.tireRepository != nil {
		return s.tireRepository
	}
	s.tireRepository = &TireRepository{
		store: s,
	}
	return s.tireRepository
}
