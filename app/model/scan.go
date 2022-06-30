package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Scan struct {
	ID         int64 
	LocationId int64  
	Plate      string 
	VinId      int64
	ScannedAt  string 
	CreatedAt  string
}

// Validate ...
func (s *Scan) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.LocationId, validation.Required),
		validation.Field(&s.Plate, validation.Required),
		validation.Field(&s.ScannedAt, validation.Required),
	)
}
