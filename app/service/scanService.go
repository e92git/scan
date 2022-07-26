package service

import (
	"scan/app/helper"
	"scan/app/model"
	"scan/app/store"
)

type ScanService struct {
	store *store.Store
	locationService *LocationService
}

func NewScan(store *store.Store, locationService *LocationService) *ScanService {
	return &ScanService{
		store: store,
		locationService: locationService,
	}
}

// func (s *ScanService) CreateBulk(locationCode string, plate string, scannedAt string) (*model.Scan, error) {

// }

func (s *ScanService) FirstOrCreate(locationCode string, plate string, scannedAt string) (*model.Scan, error) {
	l, err := s.locationService.FindByCode(locationCode)
	if err != nil {
		return nil, err
	}

	ScannedAtTime, err := helper.StrToTime(scannedAt)
	if err != nil {
		return nil, err
	}

	newScan := &model.Scan{
		LocationId: l.ID,
		Plate:      plate,
		ScannedAt:  ScannedAtTime,
	}

	return newScan, s.store.Scan().FirstOrCreate(newScan)
}
