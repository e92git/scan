package repository

import (
	"database/sql"
	"scan/app/service/location/model"
)

type Location struct {
	Db *sql.DB
}

func (r *Location) All() ([]model.Location, error) {
    rows, err := r.Db.Query("SELECT id, name, code FROM location WHERE 0")
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
