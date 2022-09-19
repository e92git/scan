package store

import (
	"scan/app/model"

	"github.com/gookit/validate"
	"gorm.io/gorm"
)

type CarMarkRepository struct {
	store *Store
}

// FirstOrCreate ...
func (r *CarMarkRepository) FirstOrCreate(m *model.CarMark) error {
	v := validate.Struct(m)
	if !v.Validate() {
		return v.Errors
	}

	res := r.store.db.Where(model.CarMark{Name: m.Name}).FirstOrCreate(m)
	return res.Error
}

// ImportFromTiresToCarMarks - import всех марок из tires в car_marks.name_in_tires.
// param clear - очистить все car_marks.name_in_tires и получить заново
func (r *CarMarkRepository) ImportFromTiresToCarMarks(clear bool) ([]string, error) {
	var logs []string
	var res *gorm.DB
	if clear == true {
		res = r.store.db.Exec("UPDATE car_marks SET name_in_tires = NULL")
		if res.Error != nil {
			return nil, res.Error
		}
	}

	rows, err := r.store.db.Raw(`
	WITH tires_marks AS (SELECT vendor as mark, COUNT(*) as model_count FROM tires GROUP BY vendor ORDER BY mark ASC)
		SELECT cm.*
		FROM car_marks cm
		LEFT JOIN tires_marks tm ON cm.name = tm.mark
		WHERE cm.name_in_tires is NULL AND tm.mark is NOT NULL
	;`).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var mark model.CarMark
	for rows.Next() {
		r.store.db.ScanRows(rows, &mark)
		mark.NameInTires = &mark.Name
		res = r.store.db.Save(mark)
		if res.Error != nil {
			return nil, res.Error
		}
		logs = append(logs, "Добавлена марка (полное совпадение): " + *mark.NameInTires)
	}

	updateSql := "UPDATE car_marks SET name_in_tires = ? WHERE name = ?"
	res = r.store.db.Exec(updateSql, "ВАЗ", "LADA (ВАЗ)")
	if res.Error != nil {
		return nil, res.Error
	}
	res = r.store.db.Exec(updateSql, "Ssang Yong", "SsangYong")
	if res.Error != nil {
		return nil, res.Error
	}
	res = r.store.db.Exec(updateSql, "Mercedes", "Mercedes-Benz")
	if res.Error != nil {
		return nil, res.Error
	}

	return logs, nil
}
