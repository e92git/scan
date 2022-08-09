package service

import (
	"errors"
	"fmt"
	"net/http"
	"scan/app/helper"
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

var constApiKey string = "AR-REST aV90cm9maW1vdl9pbnRlZ3JhdGlvbkBlOTI6MTY0MTQwMTEyMzo5OTk5OTk5OTk6VHBFcGRlMm5tdzVwcW0zbnExZ0o0dz09"

func (s *VinService) VinByPlate(plate string, authorUserId int64) (*model.Vin, error) {
	vin, err := s.FirstOrCreate(plate, authorUserId)
	if err != nil {
		return nil, err
	}

	// return if found
	if vin.IsSuccessStatus() || vin.IsErrorStatus() {
		return vin, nil
	}

	c := helper.HttpClient()
	err = s.autocodePutUid(c, vin)
	if err != nil {
		return nil, err
	}
	err = s.autocodePutReport(c, vin)
	if err != nil {
		return nil, err
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

// private

// find and put autocodeUid by plate into vin (api request)
func (s *VinService) autocodePutUid(c *http.Client, vin *model.Vin) error {
	type jsonType struct {
		QueryType      string `json:"queryType"`
		Query          string `json:"query"`
		IdempotenceKey string `json:"idempotenceKey,omitempty"`
	}
	jsonData := jsonType{
		QueryType:      "GRZ",
		Query:          vin.Plate,
		IdempotenceKey: "any_key",
	}
	statusCode, body, err := helper.SendRequest(c, http.MethodPost,
		"https://b2b-api.spectrumdata.ru/b2b/api/v1/user/reports/report_check_vehicle/_make",
		jsonData,
		constApiKey,
	)

	// big error (fake domen)
	if err != nil {
		errMessage := err.Error()
		vin.Response = nil
		vin.ResponseError = &errMessage
		vin.StatusId = model.VinStatuses.SendError
		saveErr := s.store.Vin().Save(vin)
		if saveErr != nil {
			return saveErr
		}
		return err
	}
	// fmt.Println(*statusCode)
	// fmt.Println(string(*body))
	// error 400, 500
	if *statusCode != 200 {
		response := string(*body)
		errMessage := fmt.Sprintf("%d", *statusCode) + ": api request error"
		vin.Response = &response
		vin.ResponseError = &errMessage
		vin.StatusId = model.VinStatuses.SendError
		saveErr := s.store.Vin().Save(vin)
		if saveErr != nil {
			return saveErr
		}
		return errors.New(errMessage)
	}
	// success 200
	if *statusCode == 200 {
		response := string(*body)
		vin.Response = &response
		vin.ResponseError = nil
		vin.StatusId = model.VinStatuses.SendSuccess
		saveErr := s.store.Vin().Save(vin)
		if saveErr != nil {
			return saveErr
		}
	}

	return nil
}

// find and put report by uid into vin (api request)
func (s *VinService) autocodePutReport(c *http.Client, vin *model.Vin) error {
	autocodeUid, err := vin.GetAutocodeUid()
	if err != nil {
		return err
	}
	fmt.Println(*autocodeUid)

	for i := 1; i < 3; i++ {
		statusCode, body, err := helper.SendRequest(c, http.MethodGet,
			fmt.Sprintf("https://b2b-api.spectrumdata.ru/b2b/api/v1/user/reports/%s?_content=true", *autocodeUid),
			nil,
			constApiKey,
		)
		// big error (fake domen)
		if err != nil {
			errMessage := err.Error()
			vin.ResponseError = &errMessage
			vin.StatusId = model.VinStatuses.SendError
			saveErr := s.store.Vin().Save(vin)
			if saveErr != nil {
				return saveErr
			}
			return err
		}
		// error 400, 500
		if *statusCode != 200 {
			response := string(*body)
			errMessage := fmt.Sprintf("%d", *statusCode) + ": get uid request error"
			vin.Response = &response
			vin.ResponseError = &errMessage
			vin.StatusId = model.VinStatuses.SendError
			saveErr := s.store.Vin().Save(vin)
			if saveErr != nil {
				return saveErr
			}
			return errors.New(errMessage)
		}
		// success 200
		if *statusCode == 200 {
			response := string(*body)
			vin.Response = &response
			vin.ResponseError = nil
			vin.StatusId = model.VinStatuses.Success
			// парсинг response, выделение vin, mark и других данных
			// ....
			saveErr := s.store.Vin().Save(vin)
			if saveErr != nil {
				return saveErr
			}
			// если данные еще не найдены 
			time.Sleep(2 * time.Second)
		}
	}

	return nil
}
