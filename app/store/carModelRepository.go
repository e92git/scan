package store

import (
	"scan/app/model"
	"github.com/gookit/validate"
)

type CarModelRepository struct {
	store *Store
}


// FirstOrCreate ...
func (r *CarModelRepository) FirstOrCreate(m *model.CarModel) error {
	v := validate.Struct(m)
	if !v.Validate() {
		return v.Errors
	}

	res := r.store.db.Where(model.CarModel{Name: m.Name, MarkId: m.MarkId}).FirstOrCreate(m)
	return res.Error
}
