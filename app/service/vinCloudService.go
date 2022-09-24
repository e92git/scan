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

	"gorm.io/gorm"
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
	}
}

type cloudCar struct {
	Vin       *string
	Body      *string
	MarkName  *string
	ModelName *string
	Year      *int
}

// saveSuccess распарсить ответ r, занести в vin, сохранить vin в бд
func (s *VinCloudService) saveSuccess(v *model.Vin, r *responseCloud) error {
	c := s.getCloudCar(r)
	v.Vin = c.Vin
	v.Year = c.Year
	v.StatusId = model.VinStatuses.Success
	v.ResponseError = nil

	// Body
	if c.Body != nil && c.Body != c.Vin {
		v.Body = c.Body
	}
	// find mark
	if c.MarkName != nil {
		carMark, err := s.carService.FirstMarkByName(*c.MarkName)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if carMark != nil {
			v.Mark = carMark
			v.MarkId = &carMark.ID
		}

	}
	// find model
	if c.ModelName != nil {
		carModel, err := s.carService.FirstModelByName(*c.ModelName)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if carModel != nil {
			v.Model = carModel
			v.ModelId = &carModel.ID
		}
	}

	// если не найден vin. Сохранить с ошибкой
	if v.Vin == nil && v.Vin2 == nil && v.Body == nil && v.ModelId == nil {
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

func (s *VinCloudService) getCloudCar(r *responseCloud) *cloudCar {
	c := &cloudCar{}
	if r.Response[0].Vin != "" {
		c.Vin = &r.Response[0].Vin
	}
	if r.Response[0].Body != "" {
		c.Body = &r.Response[0].Body
	}
	if r.Response[0].Year != 0 {
		c.Year = &r.Response[0].Year
	}
	if r.Response[0].Mark != "" {
		c.MarkName = &r.Response[0].Model
	}
	if r.Response[0].Model != "" {
		c.ModelName = &r.Response[0].Model
	}
	return c
}

// getCloudCarByVin получи данные из vin.ResponseCloud в виде структуры cloudCar
func (s *VinCloudService) getCloudCarByVin(vin *model.Vin) (*cloudCar, error) {
	if vin.ResponseCloud == nil {
		return &cloudCar{}, nil
	}
	r := &responseCloud{}
	err := json.Unmarshal([]byte(*vin.ResponseCloud), r)
	if err != nil {
		return nil, err
	}

	return s.getCloudCar(r), nil
}

// updateSynonyms 
// Добавить в поле `car_marks.name_synonyms` синоним марки из `vins.response_cloud``.
// При условии, что такого синонима марки еще нет там (с учетом регистра).
// Тоже самое и с моделью (`car_models.name_synonyms``)
func (s *VinCloudService) updateSynonyms(vin *model.Vin) error {
	cloudCar, err := s.getCloudCarByVin(vin)
	if err != nil {
		return err
	}	

	if cloudCar.MarkName != nil && vin.Mark != nil && vin.Mark.AddSynonym(*cloudCar.MarkName) {
		s.carService.SaveMark(vin.Mark)
	}
	if cloudCar.ModelName != nil && vin.Model != nil && vin.Model.AddSynonym(*cloudCar.ModelName) {
		s.carService.SaveModel(vin.Model)
	}
	return nil
}
