package store

import (
	"scan/app/model"
)

type LocationRepository struct {
	store *Store
}

// All ...
func (r *LocationRepository) All() ([]model.Location, error) {
	var locations []model.Location
	res := r.store.db.Find(&locations)
	return locations, res.Error
}

// FindByCode ...
func (r *LocationRepository) FindByCode(code string) (*model.Location, error) {
	loc := &model.Location{}
	res := r.store.db.Where("code = ?", code).First(loc)
	return loc, res.Error
}

// func (r *LocationRepository) All() ([]model.Location, error) {
// 	rows, err := r.store.db.Query("SELECT id, name, code FROM location")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var locations []model.Location

// 	for rows.Next() {
// 		var loc model.Location
// 		if err := rows.Scan(&loc.ID, &loc.Name, &loc.Code); err != nil {
// 			return locations, err
// 		}
// 		locations = append(locations, loc)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return locations, err
// 	}
// 	return locations, nil
// }

// // Find ...
// func (r *LocationRepository) FindByCode(code string) (*model.Location, error) {
// 	m := &model.Location{}
// 	if err := r.store.db.QueryRow(
// 		"SELECT id, code, name FROM locations WHERE code = ?",
// 		code,
// 	).Scan(
// 		&m.ID,
// 		&m.Code,
// 		&m.Name,
// 	); err != nil {
// 		return nil, err
// 	}
// 	return m, nil
// }

