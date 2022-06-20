package service

import "scan/app/model"

func (s *Server) GetAll()  ([]model.Location, error) {
	locations, err := s.Store.Location().All()
    if err != nil {
        return nil, err
    }
	return locations, nil
}
