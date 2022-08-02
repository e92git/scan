package model

type Vin struct {
	ID           int64  `json:"id" example:"7635"`
	Plate        string `json:"plate" example:"О245КМ142"`
	Vin          string `json:"vin" example:"XTA219170K0330071"`
	Vin2         string `json:"vin2" example:"XTA219170K0330071"`
	Body         string `json:"body" example:"KGC100005240"`
	MarkId       int    `json:"mark_id" example:"23"`
	ModelId      int    `json:"model_id" example:"231"`
	Year         int    `json:"year" example:"2012"`
	Data         string `json:"data" example:"{...}"`
	StatusId     int    `json:"status_id" example:"3"`
	ErrorMessage string `json:"error_message" example:"Номер не удалось найти"`
	AuthorUserId int64  `json:"author_user_id" example:"234"`
	UpdatedAt    string `json:"updated_at" example:"2022-07-23 11:23:55"`
	CreatedAt    string `json:"created_at" example:"2022-07-28 11:23:55"`

	Author User `json:"-" gorm:"foreignKey:AuthorUserId"`
}
