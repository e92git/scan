package store

import (
	"scan/app/model"

	"github.com/gookit/validate"
	"gorm.io/gorm"
)

type CarModelRepository struct {
	store *Store
}

// FirstOrCreate ...
func (r *CarModelRepository) FirstOrCreate(m *model.CarModel) error {
	v := validate.Struct(m)
	if !v.Validate() {
		return v.Errors
	}

	res := r.store.db.Where(model.CarModel{Name: m.Name, MarkId: m.MarkId}).FirstOrCreate(m)
	return res.Error
}

// ImportFromTires - import всех моделей (у кого найдена марка) из tires в car_model.name_in_tires.
// param clear - очистить все car_model.name_in_tires и получить заново
func (r *CarModelRepository) ImportFromTires(clear bool) ([]string, error) {
	var logs []string
	var res *gorm.DB
	if clear == true {
		res = r.store.db.Exec("UPDATE car_models SET name_in_tires = NULL")
		if res.Error != nil {
			return nil, res.Error
		}
	}

	rows, err := r.store.db.Raw(`
	WITH tires_model AS (
		SELECT vendor AS mark, model_seria AS model, COUNT(*) AS year_count
		FROM tires
		GROUP BY vendor, model_seria
		ORDER BY mark
		)
		SELECT m.*
		FROM car_models m
		LEFT JOIN car_marks cm ON cm.id = m.mark_id
		LEFT JOIN tires_model tm ON tm.mark = cm.name_in_tires AND tm.model = m.name
		WHERE cm.name_in_tires IS NOT NULL AND m.name_in_tires IS NULL AND tm.model IS NOT NULL
	;`).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var model model.CarModel
	for rows.Next() {
		r.store.db.ScanRows(rows, &model)
		model.NameInTires = &model.Name
		res = r.store.db.Save(model)
		if res.Error != nil {
			return nil, res.Error
		}
		logs = append(logs, "Добавлена модель (полное совпадение): "+*model.NameInTires)
	}

	updateSql := "UPDATE car_models m, car_marks cm SET m.name_in_tires = ? WHERE cm.id = m.mark_id AND m.name = ? AND cm.name = ?"
	res = r.store.db.Exec(updateSql, "Нива 4X4", "2121 (4x4)", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "Нива 4X4", "2131 (4x4)", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2101-2107", "2101", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2101-2107", "2102", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2101-2107", "2103", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2101-2107", "2104", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2101-2107", "2105", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2101-2107", "2106", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2101-2107", "2107", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2108, 2109, 21099", "2108", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2108, 2109, 21099", "2109", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2108, 2109, 21099", "21099", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2110, 2111, 2112", "2110", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2110, 2111, 2112", "2111", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2110, 2111, 2112", "2112", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2113, 2114, 2115", "2113", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2113, 2114, 2115", "2114", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "2113, 2114, 2115", "2115", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "Приора", "Priora", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "Калина", "Kalina", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "Веста", "Vesta", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "Ока", "1111 Ока", "LADA (ВАЗ)")
	res = r.store.db.Exec(updateSql, "1 series", "1 серия", "BMW")
	res = r.store.db.Exec(updateSql, "2 series", "2 серия", "BMW")
	res = r.store.db.Exec(updateSql, "3 series", "3 серия", "BMW")
	res = r.store.db.Exec(updateSql, "4 series", "4 серия", "BMW")
	res = r.store.db.Exec(updateSql, "5 series", "5 серия", "BMW")
	res = r.store.db.Exec(updateSql, "CS35 Plus", "CS35PLUS", "Changan")
	res = r.store.db.Exec(updateSql, "Tiggo", "Tiggo (T11)", "Chery")
	res = r.store.db.Exec(updateSql, "Amulet", "Amulet (A15)", "Chery")
	res = r.store.db.Exec(updateSql, "Hover", "Hover H5", "Great Wall")
	res = r.store.db.Exec(updateSql, "Stepwgn", "N-WGN", "Honda")
	res = r.store.db.Exec(updateSql, "Bluebird Sylphy", "Bluebird", "Nissan")
	res = r.store.db.Exec(updateSql, "3102", "3102 «Волга»", "ГАЗ")
	res = r.store.db.Exec(updateSql, "31105", "31105 «Волга»", "ГАЗ")
	res = r.store.db.Exec(updateSql, "Патриот", "Patriot", "УАЗ")

	return logs, nil
}
