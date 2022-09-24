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
	vinCloud    *VinCloudService
}

var vinSourse string = "cloud" // источник данных по умолчанию "autocode" | "cloud"

func NewVin(store *store.Store, vinAutocode *VinAutocodeService, vinCloud *VinCloudService) *VinService {
	return &VinService{
		store:       store,
		vinAutocode: vinAutocode,
		vinCloud:    vinCloud,
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
		// find vin in autocode or cloud now
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
		vin.StatusId = model.VinStatuses.InProcess
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

// CronFindDeffered найти отложенные поиски по госномеру (status_id=5)
func (s *VinService) CronFindDeffered() {
	var count = 60 // количесто штук за раз (12 в минуту)
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
	statusId := model.VinStatuses.InProcess
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
	// TODO: удалить vinSourse
	// всегда вызывать s.vinCloud.Find
	// готово: дополнить s.vinCloud.Find поиском марки и модели из name и name_synonyms
	// если vin.isStatusError или не получен Вин или не нашлась марка и модель - вызывать s.vinAutocode.find
	// сверять марку и модель и дополнять name_synonyms из response_cloud (name тоже должна быть в name_synonyms)
	switch vinSourse {
	case "cloud":
		err := s.vinCloud.Find(c, vin)
		if err != nil {
			return err
		}
		return nil
	case "autocode":
		err := s.vinAutocode.Find(c, vin)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Undefined source")
}

// loadRelatives подгрузить Author и Status зависимости, если их нет или они не актуальные
func (s *VinService) loadRelatives(vin *model.Vin) error {
	if vin.Author == nil || vin.Author.ID != vin.AuthorUserId {
		author, err := s.store.User().First(vin.AuthorUserId)
		if err != nil {
			return err
		}
		vin.Author = author
	}
	if vin.Status == nil || vin.Status.ID != vin.StatusId {
		status, err := s.Status(vin.StatusId)
		if err != nil {
			return err
		}
		vin.Status = status
	}
	return nil
}
