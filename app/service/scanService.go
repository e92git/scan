package service

import (
	"scan/app/model"
	"scan/app/store"
	"scan/app/helper"
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

func strToTime(scannedAt string) {
	panic("unimplemented")
}
