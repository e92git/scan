package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"scan/app/helper"
	"scan/app/model"
	"scan/app/store"
	"time"
)

type VinAutocodeService struct {
	store      *store.Store
	carService *CarService
}

func NewVinAutocode(store *store.Store, carService *CarService) *VinAutocodeService {
	return &VinAutocodeService{
		store:      store,
		carService: carService,
	}
}

var constApiKey string = "AR-REST aV90cm9maW1vdl9pbnRlZ3JhdGlvbkBlOTI6MTY0MTQwMTEyMzo5OTk5OTk5OTk6VHBFcGRlMm5tdzVwcW0zbnExZ0o0dz09"

// find and put autocodeUid by plate into vin (api request)
func (s *VinAutocodeService) AutocodePutUid(c *http.Client, vin *model.Vin) error {
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
		return s.saveError(vin, nil, err)
	}
	// error 400, 500
	if *statusCode != 200 {
		response := string(*body)
		responseError := fmt.Sprintf("%d", *statusCode) + ": autocodePutUid request error"
		return s.saveError(vin, &response, errors.New(responseError))
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
		return nil
	}

	return errors.New("autocodePutUid statusCode not found")
}

type response struct {
	Size int              `json:"size"`
	Data []responseReport `json:"data"`
}
type responseReport struct {
	VehicleId  string `json:"vehicle_id"`
	ProgressOk int    `json:"progress_ok"`
	Content    struct {
		Identifiers struct {
			Vehicle struct {
				Vin    string `json:"vin"`
				Body   string `json:"body"`
				RegNum string `json:"reg_num"`
			} `json:"vehicle"`
			Manufacture struct {
				Vin string `json:"vin,omitempty"`
			} `json:"manufacture"`
		} `json:"identifiers"`
		TechData struct {
			Brand struct {
				Name struct {
					Normalized string `json:"normalized"`
				} `json:"name"`
			} `json:"brand"`
			Model struct {
				Name struct {
					Normalized string `json:"normalized"`
				} `json:"name"`
			} `json:"model"`
			Year int `json:"year"`
		} `json:"tech_data"`
	} `json:"content"`
}

// find and put report by uid into vin (api request)
func (s *VinAutocodeService) AutocodePutReport(c *http.Client, vin *model.Vin) error {
	autocodeUid, err := vin.GetAutocodeUid()
	if err != nil {
		return err
	}

	var maxAttempts int = 5
	var interval time.Duration = 1 * time.Second

	// получить report с maxAttempts попыток
	for i := 1; i <= maxAttempts; i++ {
		statusCode, body, err := helper.SendRequest(c, http.MethodGet,
			fmt.Sprintf("https://b2b-api.spectrumdata.ru/b2b/api/v1/user/reports/%s?_content=true", *autocodeUid),
			nil,
			constApiKey,
		)
		// big error (fake domen)
		if err != nil {
			return s.saveError(vin, vin.Response, err)
		}
		// error 400, 500
		if *statusCode != 200 {
			responseJson := string(*body)
			responseError := fmt.Sprintf("%d", *statusCode) + ": autocodePutReport request error"
			return s.saveError(vin, &responseJson, errors.New(responseError))
		}
		// success 200
		if *statusCode == 200 {
			responseString := string(*body)
			r := response{}
			err := json.Unmarshal([]byte(*body), &r)
			if err != nil {
				return s.saveError(vin, &responseString, err)
			}

			// если результат еще не получен, подождать 2 секунды и проверить еще раз
			if r.Size == 0 || r.Data[0].ProgressOk == 0 {
				time.Sleep(interval)
				continue
			}

			vin.Response = &responseString
			return s.saveSuccess(vin, &r)
		}
	}

	return errors.New("autocodePutReport not found")
}

// saveSuccess распарсить ответ r, занести в vin, сохранить vin в бд
func (s *VinAutocodeService) saveSuccess(vin *model.Vin, r *response) error {
	if len(r.Data) == 0 {
		return errors.New("saveSuccess response is empty")
	}

	var vin1 *string = nil
	if r.Data[0].Content.Identifiers.Vehicle.Vin != "" {
		vin1 = &r.Data[0].Content.Identifiers.Vehicle.Vin
	}
	var vin2 *string = nil
	if r.Data[0].Content.Identifiers.Manufacture.Vin != "" {
		vin2 = &r.Data[0].Content.Identifiers.Manufacture.Vin
	}
	var body *string = nil
	if r.Data[0].Content.Identifiers.Vehicle.Body != "" {
		body = &r.Data[0].Content.Identifiers.Vehicle.Body
	}
	var year *int = nil
	if r.Data[0].Content.TechData.Year != 0 {
		year = &r.Data[0].Content.TechData.Year
	}
	var err error
	var markId *int = nil
	var mark *model.CarMark = nil
	if r.Data[0].Content.TechData.Brand.Name.Normalized != "" {
		mark, err = s.carService.FirstOrCreateMark(r.Data[0].Content.TechData.Brand.Name.Normalized)
		if err != nil {
			return err
		}
		markId = &mark.ID
	}
	var modelId *int = nil
	var carModel *model.CarModel = nil
	if markId != nil && r.Data[0].Content.TechData.Model.Name.Normalized != "" {
		carModel, err = s.carService.FirstOrCreateModel(*markId, r.Data[0].Content.TechData.Model.Name.Normalized)
		if err != nil {
			return err
		}
		modelId = &carModel.ID
	}

	vin.ResponseError = nil
	vin.StatusId = model.VinStatuses.Success
	vin.Vin = vin1
	vin.Vin2 = vin2
	vin.Body = body
	vin.Year = year
	vin.MarkId = markId
	vin.ModelId = modelId

	// если не найден vin. Сохранить с ошибкой
	if vin.Vin == nil && vin.Vin2 == nil && vin.Body == nil && vin.MarkId == nil {
		responseError := "200: vins are empty"
		vin.ResponseError = &responseError
		vin.StatusId = model.VinStatuses.SendError
	}

	// сохранить
	saveErr := s.store.Vin().Save(vin)
	if saveErr != nil {
		return saveErr
	}

	// load data
	if mark != nil {
		vin.Mark = mark
	}
	if carModel != nil {
		vin.Model = carModel
	}

	// успешный выход
	return nil
}

func (s *VinAutocodeService) saveError(vin *model.Vin, response *string, err error) error {
	responseError := err.Error()
	vin.Response = response
	vin.ResponseError = &responseError
	vin.StatusId = model.VinStatuses.SendError
	saveErr := s.store.Vin().Save(vin)
	if saveErr != nil {
		return saveErr
	}
	return err
}
