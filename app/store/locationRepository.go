package store

import (
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
