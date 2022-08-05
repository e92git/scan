package store

import (
	"scan/app/model"
	"github.com/gookit/validate"
)

type VinRepository struct {
	store *Store
}


// FirstOrCreate ...
func (r *VinRepository) FirstOrCreate(m *model.Vin) error {
	v := validate.Struct(m)
	if !v.Validate() {
		return v.Errors
	}

	res := r.store.db.Where(model.Vin{Plate: m.Plate}).FirstOrCreate(m)
	return res.Error
}

// Save 
func (r *VinRepository) Save(m *model.Vin) error {
	v := validate.Struct(m)
	if !v.Validate() {
		return v.Errors
	}

	res := r.store.db.Save(m)
	return res.Error
}
