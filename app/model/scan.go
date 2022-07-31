package model

import (
	// "database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Scan struct {
	ID         int64         `json:"id" example:"76352"`
	LocationId int64         `json:"location_id" example:"12"`
	Plate      string        `json:"plate" example:"О245КМ142"`
	// VinId      sql.NullInt64 `json:"vin_id,omitempty" swaggertype:"integer"`
	UserId     int64 		 `json:"user_id,omitempty" example:"234"`
	ScannedAt  time.Time     `json:"scanned_at" example:"2022-07-23 11:23:55"`
	CreatedAt  time.Time     `json:"created_at" example:"2022-07-28 11:23:55"`
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
