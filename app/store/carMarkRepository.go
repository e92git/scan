package store

import (
	"scan/app/model"
	"github.com/gookit/validate"
)

type CarMarkRepository struct {
	store *Store
}


// FirstOrCreate ...
func (r *CarMarkRepository) FirstOrCreate(m *model.CarMark) error {
	v := validate.Struct(m)
	if !v.Validate() {
		return v.Errors
	}

	res := r.store.db.Where(model.CarMark{Name: m.Name}).FirstOrCreate(m)
	return res.Error
}
