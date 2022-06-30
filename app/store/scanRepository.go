package store

import (
	// "scan/app/model"
)

type ScanRepository struct {
	store *Store
}

// Create ...
// func (r *ScanRepository) Create(s *model.Scan) error {
// 	if err := s.Validate(); err != nil {
// 		return err
// 	}

// 	res, err := r.store.db.Exec(
// 		"INSERT INTO scans (location_id, plate, scanned_at) VALUES (?, ?, ?)",
// 		s.LocationId,
// 		s.Plate,
// 		s.ScannedAt,
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		return err
// 	}
// 	s.ID = id
// 	return nil
// }

// func (r *ScanRepository) All() ([]model.Scan, error) {
// 	rows, err := r.store.db.Query("SELECT id, name, code FROM location")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var locations []model.Location

// 	for rows.Next() {
// 		var loc model.Location
// 		if err := rows.Scan(&loc.Id, &loc.Name, &loc.Code); err != nil {
// 			return locations, err
// 		}
// 		locations = append(locations, loc)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return locations, err
// 	}
// 	return locations, nil
// }

// Find ...
// func (r *ScanRepository) FindByCode(code string) (*model.Location, error) {
// 	u := &model.Location{}
// 	if err := r.store.db.QueryRow(
// 		"SELECT id, code, name FROM location WHERE code = ?",
// 		code,
// 	).Scan(
// 		&u.Id,
// 		&u.Code,
// 		&u.Name,
// 	); err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, ErrRecordNotFound
// 		}
// 		return nil, err
// 	}
// 	return u, nil
// }
