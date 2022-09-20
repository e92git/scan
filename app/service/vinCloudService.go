package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"scan/app/apiserver"
	"scan/app/helper"
	"scan/app/model"
	"scan/app/store"
)

type VinCloudService struct {
	config     *apiserver.Config
	store      *store.Store
	carService *CarService
}

func NewVinCloud(config *apiserver.Config, store *store.Store, carService *CarService) *VinCloudService {
	return &VinCloudService{
		config:     config,
		store:      store,
		carService: carService,
	}
}

// find and put vin data into vin table (api request)
func (s *VinCloudService) Find(c *http.Client, vin *model.Vin) error {
	statusCode, body, err := helper.SendRequest(c, http.MethodGet,
		"https://api.clouddata.ru/v1/car_autofill_lite?car_number="+vin.Plate,
		nil,
		s.config.ApiKeyCloud,
	)

	// big error (fake domen)
	if err != nil {
		return s.saveError(vin, nil, err)
	}
	// error 400, 500
	if *statusCode != 200 {
		response := string(*body)
		responseError := fmt.Sprintf("%d", *statusCode) + ": VinCloudService.find() request error"
		return s.saveError(vin, &response, errors.New(responseError))
	}
	// success 200
	if *statusCode == 200 {
		responseString := string(*body)
		r := responseCloud{}
		err := json.Unmarshal([]byte(*body), &r)
		if err != nil {
			return s.saveError(vin, &responseString, err)
		}

		if len(r.Response) == 0 {
			return s.saveError(vin, &responseString, errors.New("200: vins are empty. VinCloudService.find() response[] is empty"))
		}

		vin.ResponseCloud = &responseString
		return s.saveSuccess(vin, &r)
	}

	return errors.New("VinCloud statusCode not found")
}

type responseCloud struct {
	Code      int  `json:"code"`
	Status    bool `json:"status"`
	RequestId int  `json:"request_id"`
	Response  []struct {
		Plate string `json:"car_number"`
		Vin   string `json:"car_vin"`
		Body  string `json:"car_body"`
		Mark  string `json:"car_brand"`
		Model string `json:"car_model"`
		Year  int    `json:"car_year"`
	} `json:"response"`
}

// saveSuccess распарсить ответ r, занести в vin, сохранить vin в бд
func (s *VinCloudService) saveSuccess(v *model.Vin, r *responseCloud) error {
	var vin *string = nil
	if r.Response[0].Vin != "" {
		vin = &r.Response[0].Vin
	}
	var body *string = nil
	if r.Response[0].Body != "" {
		body = &r.Response[0].Body
	}
	var year *int = nil
	if r.Response[0].Year != 0 {
		year = &r.Response[0].Year
	}

	v.Vin = vin
	if body != vin {
		v.Body = body
	}
	v.Year = year
	v.StatusId = model.VinStatuses.Success
	v.ResponseError = nil

	// если не найден vin. Сохранить с ошибкой
	if v.Vin == nil && v.Vin2 == nil && v.Body == nil && v.MarkId == nil {
		responseError := "200: vins are empty"
		v.ResponseError = &responseError
		v.StatusId = model.VinStatuses.SendError
	}

	saveErr := s.store.Vin().Save(v)
	if saveErr != nil {
		return saveErr
	}
	return nil
}

func (s *VinCloudService) saveError(vin *model.Vin, response *string, err error) error {
	responseError := err.Error()
	vin.ResponseError = &responseError
	vin.StatusId = model.VinStatuses.SendError
	saveErr := s.store.Vin().Save(vin)
	if saveErr != nil {
		return saveErr
	}
	vin.Response = response
	saveErr = s.store.Vin().Save(vin)
	if saveErr != nil {
		return saveErr
	}
	return err
}
