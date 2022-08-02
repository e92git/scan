package store

import (
	"scan/app/model"
)

type VinRepository struct {
	store *Store
}


// Create ...
func (r *VinRepository) Create(s *model.Vin) error {
	// if err := s.Validate(); err != nil {
	// 	return err
	// }

	res := r.store.db.Create(s)
	return res.Error
}
