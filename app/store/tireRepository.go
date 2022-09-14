package store

import (
	"scan/app/helper"
	"scan/app/model"
	"strings"
)

type TireRepository struct {
	store *Store
}

// TireAnalytics
func (r *TireRepository) TireAnalytics() (*model.TireAnalyticsResponse, error) {
	response := &model.TireAnalyticsResponse{}

	rows, err := r.store.db.Raw(`
		SELECT v.plate, t.vendor as mark, t.model_seria as model, t.year, GROUP_CONCAT(DISTINCT t.tyres_factory ORDER BY t.tyres_factory ASC SEPARATOR '|') as sizes
		FROM vins v
		LEFT JOIN car_models m ON m.id = v.model_id
		LEFT JOIN car_marks cm ON cm.id = m.mark_id
		LEFT JOIN tires t ON t.vendor = cm.name_in_tires AND t.model_seria = m.name_in_tires AND t.year = v.year
		WHERE m.name_in_tires IS NOT NULL AND t.year is NOT NULL
		GROUP BY v.plate
	;`).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var plate, mark, carModel, year, sizeConcat string
	for rows.Next() {
		rows.Scan(&plate, &mark, &carModel, &year, &sizeConcat)
		// size
		findSizes := getSizes(sizeConcat)
		count := len(findSizes)
		if count == 0 {
			continue
		}
		//index
		index := 1 / float32(count)
		// plates, cars
		car := mark + " " + carModel + " " + year
		for _, findSize := range findSizes {
			// get
			size := response.FirstOrCreateSize(findSize)
			// set
			size.Index += index
			size.Plates.Count++
			size.Plates.List = append(size.Plates.List, plate)
			if !helper.InArray(&size.Cars.List, car) {
				size.Cars.Count++
				size.Cars.List = append(size.Cars.List, car)
			}
		}
	}
	
	return response, nil
}

/// private

// getSizes - получить массив уникльных размеров шин для данного госномера
func getSizes(sizeConcat string) []string {
	sizes := strings.Split(sizeConcat, "|")
	var res []string
	for _, size := range sizes {
		size = strings.Trim(size, " ")
		if size != "" && !helper.InArray(&res, size) {
			res = append(res, size)
		}
	}
	return res
}
