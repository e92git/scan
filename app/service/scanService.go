package service

import (
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

func (s *ScanService) Create(locationCode string, plate string, scannedAt string) (*model.Scan, error) {
	l, err := s.locationService.FindByCode(locationCode)
	if err != nil {
		return nil, err
	}

	newScan := &model.Scan{
		LocationId: l.ID,
		Plate:      plate,
		ScannedAt:  scannedAt,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05.999"),
	}

	return newScan, s.store.Scan().Create(newScan)
}
