package service

import (
	"errors"
	"log"
	"scan/app/helper"
	"scan/app/model"
	"scan/app/store"

	"gorm.io/gorm"
)

type VinService struct {
	store       *store.Store
	vinAutocode *VinAutocodeService
}

func NewVin(store *store.Store, vinAutocode *VinAutocodeService) *VinService {
	return &VinService{
		store:       store,
		vinAutocode: vinAutocode,
	}
}

// VinByPlate получить или создать новую запись с поиском vin.
// immediately - true: Данные появятся сразу в ответе. false: Отложенное получение данных
func (s *VinService) VinByPlate(plate string, authorUserId int64, immediately bool) (*model.Vin, error) {
	vin, err := s.firstOrCreateByPlate(plate, authorUserId, immediately)
	if err != nil {
		return nil, err
	}

	// return if found
	if vin.IsSuccessStatus() {
		return vin, nil
	}

	// find vin if immediately == true
	if immediately == true {
		// find vin in autocode now
		err = s.findVin(vin)
		if err != nil {
			return nil, err
		}
	}

	err = s.loadRelatives(vin)
	if err != nil {
		return nil, err
	}

	return vin, nil
}

// VinByPlateBulk получить или создать массив новых записей с поиском vin.
// Данные появятся потом (по крону FindDeffered)
func (s *VinService) VinByPlateBulk(plates []string, authorUserId int64) ([]*model.Vin, error) {
	var vins []*model.Vin
	for _, plate := range plates {
		vin, err := s.VinByPlate(plate, authorUserId, false)
		if err != nil {
			return nil, err
		}

		vins = append(vins, vin)
	}

	return vins, nil
}

// Status модель статуса по статус ид
func (s *VinService) Status(id int) (*model.VinStatus, error) {
	newVinStatus := &model.VinStatus{
		ID: id,
	}

	return newVinStatus, s.store.Vin().StatusFirst(newVinStatus)
}

// PutDeffered найти отложенные поиски по госномеру (status_id=5) и выполнить их поиск в стороннем сервисе
func (s *VinService) FindDeffered(count int) error {
	var i = 0
	var err error
	for i < count {
		// найти шаблон vin
		vin := &model.Vin{
			StatusId: model.VinStatuses.CreatedDeferred,
		}
		err = s.store.Vin().First(vin)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil // если нет записей для анализа - выход
		}
		if err != nil {
			return err // если ошибка в получении - выход с ошибкой
		}

		// сменить статус
		vin.StatusId = model.VinStatuses.Created
		err := s.store.Vin().Save(vin)
		if err != nil {
			return err
		}

		// найти vin
		err = s.findVin(vin)
		if err != nil {
			return err
		}

		// следующий шаг (i из count)
		i++
	}
	return nil
}

// CronFindDeffered найти отложенные поиски по госномеру (status_id=5). count - 12 штук за раз
func (s *VinService) CronFindDeffered() {
	var count = 12 // количесто штук за раз
	err := s.FindDeffered(count)
	if err != nil {
		log.Println("CronError: CronFindDeffered" + err.Error())
	}
}

//
// private
//

// firstOrCreateByPlate создать пустую запись в таблице vin со статусом Created или CreatedDeferred
// или вернуть если уже есть по этому госномеру
func (s *VinService) firstOrCreateByPlate(plate string, authorUserId int64, immediately bool) (*model.Vin, error) {
	statusId := model.VinStatuses.Created
	if !immediately {
		statusId = model.VinStatuses.CreatedDeferred
	}
	newVin := &model.Vin{
		Plate:        helper.ClearPlate(plate),
		AuthorUserId: authorUserId,
		StatusId:     statusId,
	}

	return newVin, s.store.Vin().FirstOrCreateByPlate(newVin)
}

// дополнить объект vin, с вин-кодом и др. данными (по грз vin.plate)
func (s *VinService) findVin(vin *model.Vin) error {
	c := helper.HttpClient()
	err := s.vinAutocode.AutocodePutUid(c, vin)
	if err != nil {
		return err
	}
	err = s.vinAutocode.AutocodePutReport(c, vin)
	if err != nil {
		return err
	}
	return nil
}

// loadRelatives загрузить Author и Status зависимости
func (s *VinService) loadRelatives(vin *model.Vin) error {
	if vin.Author == nil {
		author, err := s.store.User().First(vin.AuthorUserId)
		if err != nil {
			return err
		}
		vin.Author = author
	}
	if vin.Status == nil {
		status, err := s.Status(vin.StatusId)
		if err != nil {
			return err
		}
		vin.Status = status
	}
	return nil
}
