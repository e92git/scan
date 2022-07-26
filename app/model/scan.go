package model

import (
	"database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Scan struct {
	ID         int64         `json:"id"`
	LocationId int64         `json:"location_id"`
	Plate      string        `json:"plate"`
	VinId      sql.NullInt64 `json:"vin_id,omitempty"`
	UserId     int64 		 `json:"user_id,omitempty"`
	ScannedAt  time.Time     `json:"scanned_at"`
	CreatedAt  time.Time     `json:"created_at"`
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
