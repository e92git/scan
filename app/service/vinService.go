package service

import (
	"fmt"
	// "net/http"
	// "scan/app/helper"
	"scan/app/model"
	"scan/app/store"
	"time"
)

type VinService struct {
	store *store.Store
}

func NewVin(store *store.Store) *VinService {
	return &VinService{
		store: store,
	}
}

func (s *VinService) VinByPlate(plate string, authorUserId int64) (*model.Vin, error) {
	vin, err := s.FirstOrCreate(plate, authorUserId)
	if err != nil {
		return nil, err
	}

	fmt.Println(vin.CreatedAt)
	fmt.Println(time.Now())

	if vin.NeedSend() {
		// c := helper.HttpClient()
		// statusCode, body, err := helper.SendRequest(c, http.MethodPost,
		// 	// "https://b2b-api.spectrumdata.ru/b2b/api/v1/user/reports/report_check_vehicle/_make",
		// 	"https://b2b-_make",
		// 	"F",
		// )
		// if err != nil {
		// 	errMessage := err.Error()
		// 	vin.ResponseError = &errMessage
		// 	vin.StatusId = model.VinStatuses.SendError
		// 	saveErr := s.store.Vin().Save(vin)
		// 	if saveErr != nil {
		// 		return nil, saveErr
		// 	}
		// 	return nil, err
		// }
		// fmt.Println(*statusCode)
		// fmt.Println(string(*body))
	}

	return vin, nil
}

func (s *VinService) FirstOrCreate(plate string, authorUserId int64) (*model.Vin, error) {
	newVin := &model.Vin{
		Plate:        plate,
		AuthorUserId: authorUserId,
		StatusId:     model.VinStatuses.Created,
		UpdatedAt:    time.Now(),
		CreatedAt:    time.Now(),
	}

	return newVin, s.store.Vin().FirstOrCreate(newVin)
}
