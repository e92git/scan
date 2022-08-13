package store

import (
	"github.com/gookit/validate"
	"scan/app/model"
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

	res := r.store.db.
		Joins("Mark").
		Joins("Model").
		Joins("Author").
		Joins("Status").
		Where(model.Vin{Plate: m.Plate}).
		FirstOrCreate(m)
	return res.Error
}

func (r *VinRepository) StatusFirst(m *model.VinStatus) error {
	v := validate.Struct(m)
	if !v.Validate() {
		return v.Errors
	}

	res := r.store.db.First(m)
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
