package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Scan struct {
	ID         int64     `json:"id" example:"76352"`
	LocationId int64     `json:"location_id" example:"12"`
	Plate      string    `json:"plate" example:"О245КМ142"`
	UserId     int64     `json:"user_id" example:"234"`
	ScannedAt  time.Time `json:"scanned_at" example:"2022-07-29T11:23:55+07:00"`
	CreatedAt  time.Time `json:"created_at" example:"2022-08-04T12:23:52.372+07:00"`
}

// Validate ...
func (s *Scan) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.LocationId, validation.Required),
		validation.Field(&s.Plate, validation.Required),
		validation.Field(&s.ScannedAt, validation.Required),
	)
}

// // MarshalJSON вид структуры при конверции в json 
// func (s *Scan) MarshalJSON() ([]byte, error) {
// 	type Alias Scan
// 	return json.Marshal(&struct {
// 		*Alias
// 		ScannedAt string `json:"scanned_at"`
// 		CreatedAt string `json:"created_at"`
// 	}{
// 		Alias: (*Alias)(s),
// 		ScannedAt: s.CreatedAt.Format("2006-01-02 15:04:05"),
// 		CreatedAt: s.CreatedAt.Format("2006-01-02 15:04:05.000"),
// 	})
// }
