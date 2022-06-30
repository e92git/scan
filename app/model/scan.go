package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Scan struct {
	ID         int64  `json:"id,omitempty"`
	LocationId int64  `json:"location_id"`
	Plate      string `json:"plate"`
	VinId      int64  `json:"vin_id,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	ScannedAt  string `json:"scanned_at"`
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
