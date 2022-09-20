package store

import (
	"scan/app/model"

	"github.com/gookit/validate"
	"gorm.io/gorm/clause"
)

type VinRepository struct {
	store *Store
}

// FirstOrCreateByPlate ...
func (r *VinRepository) FirstOrCreateByPlate(m *model.Vin) error {
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

func (r *VinRepository) First(m *model.Vin) error {
	res := r.store.db.Where(m).First(m)
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

	res := r.store.db.Omit(clause.Associations).Save(m)
	return res.Error
}
