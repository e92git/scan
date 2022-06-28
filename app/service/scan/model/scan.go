package model

type Scan struct {
	ID         int64  `json:"id,omitempty"`
	LocationId int64  `json:"location_id"`
	Plate      string `json:"plate"`
	VinId      int64  `json:"vin_id,omitempty"`
	ScannedAt  string `json:"scanned_at"`
	CreatedAt  string `json:"created_at,omitempty"`
}