package service

import (
	"scan/app/model"
	"scan/app/store"
	"time"
)

type ScanService struct {
	store *store.Store
}

func NewScan(store *store.Store) *ScanService {
	return &ScanService{
		store: store,
	}
}

func (s *ScanService) FirstOrCreate(locationCode string, plate string, scannedAt string) (*model.Scan, error) {
	l, err := s.store.Location().FindByCode(locationCode)
	if err != nil {
		return nil, err
	}

	t, err := time.ParseInLocation("2006-01-02 15:04:05", scannedAt, time.Local)
	if err != nil {
		return nil, err
	}

	newScan := &model.Scan{
		LocationId: l.ID,
		Plate:      plate,
		ScannedAt:  t,
	}

	return newScan, s.store.Scan().FirstOrCreate(newScan)
}
