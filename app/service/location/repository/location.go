package repository

import (
	"database/sql"
	"scan/app/service/location/model"
)

type Location struct {
	db *sql.DB
}

func New(db *sql.DB) *Location {
	return &Location{
		db: db,
	}
}

func (r *Location) All() ([]model.Location, error) {
    rows, err := r.db.Query("SELECT id, name, code FROM location")
    if err != nil {
		return nil, err
	}
    defer rows.Close()

    var locations []model.Location

	for rows.Next() {
		var loc model.Location
		if err := rows.Scan(&loc.ID, &loc.Name, &loc.Code); err != nil {
			return locations, err
		}
		locations = append(locations, loc)
	}
	if err = rows.Err(); err != nil {
		return locations, err
	}
	return locations, nil
}
