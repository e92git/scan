package service

import (
	"scan/app/model"
	"scan/app/store"
)

type TireService struct {
	store *store.Store
}

func NewTire(store *store.Store) *TireService {
	return &TireService{
		store: store,
	}
}

func (s *TireService) GetTireAnalytics() (*model.TireAnalyticsResponse, error) {
	// сформировать TireAnalyticsResponse
	r, err := s.store.Tire().TireAnalytics()
	if err != nil {
		return nil, err
	}
	// Подсчитать параметры
	r.CaclSizesParams()

	// расскоментировать для получения csv из командной строки
	// d := ";"
	// for _, size := range r.Sizes {
	// 	fmt.Println(strconv.Itoa(size.Rank) +d+ size.Size +d+ fmt.Sprintf("%f", size.Percent) +d+ fmt.Sprintf("%f", size.Index) +d+ strings.Join(size.Plates.List[:],",") +d+ strings.Join(size.Cars.List[:],","))
	// }

	return r, nil
}

type GetTireSyncResponse struct {
	Logs []string
}
func (s *TireService) GetTireSync() (*model.TireAnalyticsResponse, error) {
	r, err := s.store.Tire().TireAnalytics()
	if err != nil {
		return nil, err
	}
	return r, nil
}
