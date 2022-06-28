package store

import (
	"database/sql"
	"scan/app/model"
)

type LocationRepository struct {
	store *Store
}

func (r *LocationRepository) All() ([]model.Location, error) {
	rows, err := r.store.db.Query("SELECT id, name, code FROM location")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []model.Location

	for rows.Next() {
		var loc model.Location
		if err := rows.Scan(&loc.Id, &loc.Name, &loc.Code); err != nil {
			return locations, err
		}
		locations = append(locations, loc)
	}
	if err = rows.Err(); err != nil {
		return locations, err
	}
	return locations, nil
}

// Find ...
func (r *LocationRepository) FindByCode(code string) (*model.Location, error) {
	u := &model.Location{}
	if err := r.store.db.QueryRow(
		"SELECT id, code, name FROM location WHERE code = ?",
		code,
	).Scan(
		&u.Id,
		&u.Code,
		&u.Name,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

