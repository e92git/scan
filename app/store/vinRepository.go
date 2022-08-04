package store

import (
	"scan/app/model"
)

type VinRepository struct {
	store *Store
}


// FirstOrCreate ...
func (r *VinRepository) FirstOrCreate(m *model.Vin) error {
	if err := m.Validate(); err != nil {
		return err
	}

	res := r.store.db.Where(model.Vin{Plate: m.Plate}).FirstOrCreate(m)
	return res.Error
}

// Save 
func (r *VinRepository) Save(m *model.Vin) error {
	if err := m.Validate(); err != nil {
		return err
	}

	res := r.store.db.Save(m)
	return res.Error
}
