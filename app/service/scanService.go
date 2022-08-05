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

// func (s *ScanService) CreateBulk(locationCode string, plate string, scannedAt string) (*model.Scan, error) {

// }

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
		Plate:      plate,
		ScannedAt:  scannedAt,
		UserId:     userId,
	}

	return newScan, s.Create(newScan)
}

func (s *ScanService) Create(scan *model.Scan) error {
	return s.store.Scan().Create(scan)
}
