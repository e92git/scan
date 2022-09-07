package service

import (
	"scan/app/helper"
	"scan/app/model"
	"scan/app/store"
	"time"
)

type ScanService struct {
	store           *store.Store
	locationService *LocationService
}

func NewScan(store *store.Store, locationService *LocationService) *ScanService {
	return &ScanService{
		store:           store,
		locationService: locationService,
	}
}

func (s *ScanService) AddScanWithPrepare(locationCode string, plate string, scannedAt string, userId int64) (*model.Scan, error) {
	l, err := s.locationService.FindByCode(locationCode)
	if err != nil {
		return nil, err
	}

	sAt, err := helper.StrToTime(scannedAt)
	if err != nil {
		return nil, err
	}

	return s.AddScan(l.ID, plate, sAt, userId)
}

func (s *ScanService) AddScan(locationId int64, plate string, scannedAt time.Time, userId int64) (*model.Scan, error) {
	newScan := &model.Scan{
		LocationId: locationId,
		Plate:      helper.ClearPlate(plate),
		ScannedAt:  scannedAt,
		UserId:     userId,
	}

	return newScan, s.Create(newScan)
}

func (s *ScanService) Create(scan *model.Scan) error {
	return s.store.Scan().Create(scan)
}

type Scans struct {
	Plate string `json:"plate" example:"Т237АС142" validate:"required"`
	Date  string `json:"date" example:"2022-07-06 10:31:12" validate:"required"`
}

// CreateBulk
func (s *ScanService) CreateBulk(locationId int64, scans *[]Scans, userId int64) error {
	var scanModels []model.Scan
	for _, scan := range *scans {
		scannedAt, err := helper.StrToTime(scan.Date)
		if err != nil {
			return err
		}
		scanModels = append(scanModels, model.Scan{
			LocationId: locationId,
			Plate:      helper.ClearPlate(scan.Plate),
			ScannedAt:  scannedAt,
			UserId:     userId,
		})
	}
	return s.store.Scan().CreateBulk(&scanModels)
}
