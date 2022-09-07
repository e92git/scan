package store

import (
	"scan/app/model"

	"github.com/gookit/validate"
)

type ScanRepository struct {
	store *Store
}

// Create ...
func (r *ScanRepository) Create(s *model.Scan) error {
	v := validate.Struct(s)
	if !v.Validate() {
		return v.Errors
	}

	res := r.store.db.Create(s)
	return res.Error
}

// FirstOrCreate ...
func (r *ScanRepository) FirstOrCreate(s *model.Scan) error {
	v := validate.Struct(s)
	if !v.Validate() {
		return v.Errors
	}

	res := r.store.db.Where(s).FirstOrCreate(s)
	return res.Error
}

// First ...
func (r *ScanRepository) First(s *model.Scan) error {
	res := r.store.db.First(s)
	return res.Error
}

// CreateBulk
func (r *ScanRepository) CreateBulk(s *[]model.Scan) error {
	for _, scan := range *s {
		v := validate.Struct(scan)
		if !v.Validate() {
			return v.Errors
		}
	}
	res := r.store.db.CreateInBatches(s, 1000)
	return res.Error
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
