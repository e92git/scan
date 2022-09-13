package main

import (
	"fmt"
	"log"
	"scan/app/apiserver"
	"scan/app/model"

	"gorm.io/gorm"
)

// добавить в поля name_in_tires название марок и моделей как в таблице tires
func main() {
	// load config
	config, err := apiserver.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// load store
	db, err := apiserver.ConnectGorm(config.Dsn, config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	// Import Marks
	err = marks(db, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Успешно завершено Марки.")

	// Import Models
	err = models(db, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Успешно завершено Модели.")

	fmt.Println("Успешно завершено Всё.")
}

// marks - import всех марок из tires в car_marks.name_in_tires.
// param clear - очистить все car_marks.name_in_tires и получить заново
func marks(db *gorm.DB, clear bool) error {
	var res *gorm.DB
	if clear == true {
		res = db.Exec("UPDATE car_marks SET name_in_tires = NULL")
		if res.Error != nil {
			return res.Error
		}
	}

	rows, err := db.Raw(`
	WITH tires_marks AS (SELECT vendor as mark, COUNT(*) as model_count FROM tires GROUP BY vendor ORDER BY mark ASC)
		SELECT cm.*
		FROM car_marks cm
		LEFT JOIN tires_marks tm ON cm.name = tm.mark
		WHERE cm.name_in_tires is NULL AND tm.mark is NOT NULL
	;`).Rows()
	defer rows.Close()
	if err != nil {
		return err
	}
	var mark model.CarMark
	for rows.Next() {
		db.ScanRows(rows, &mark)
		mark.NameInTires = &mark.Name
		res = db.Save(mark)
		if res.Error != nil {
			return res.Error
		}
		fmt.Println("Добавлена марка (полное совпадение): " + *mark.NameInTires)
	}

	updateSql := "UPDATE car_marks SET name_in_tires = ? WHERE name = ?"
	res = db.Exec(updateSql, "ВАЗ", "LADA (ВАЗ)")
	if res.Error != nil {
		return res.Error
	}
	res = db.Exec(updateSql, "Ssang Yong", "SsangYong")
	if res.Error != nil {
		return res.Error
	}
	res = db.Exec(updateSql, "Mercedes", "Mercedes-Benz")
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// models - import всех моделей (у кого найдена марка) из tires в car_model.name_in_tires.
// param clear - очистить все car_model.name_in_tires и получить заново
func models(db *gorm.DB, clear bool) error {
	var res *gorm.DB
	if clear == true {
		res = db.Exec("UPDATE car_models SET name_in_tires = NULL")
		if res.Error != nil {
			return res.Error
		}
	}

	rows, err := db.Raw(`
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
		return err
	}
	var model model.CarModel
	for rows.Next() {
		db.ScanRows(rows, &model)
		model.NameInTires = &model.Name
		res = db.Save(model)
		if res.Error != nil {
			return res.Error
		}
		fmt.Println("Добавлена модель (полное совпадение): " + *model.NameInTires)
	}

	updateSql := "UPDATE car_models m, car_marks cm SET m.name_in_tires = ? WHERE cm.id = m.mark_id AND m.name = ? AND cm.name = ?"
	res = db.Exec(updateSql, "Нива 4X4", "2121 (4x4)", "LADA (ВАЗ)")
	if res.Error != nil {
		return res.Error
	}
	res = db.Exec(updateSql, "Нива 4X4", "2131 (4x4)", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2101-2107", "2101", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2101-2107", "2102", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2101-2107", "2103", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2101-2107", "2104", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2101-2107", "2105", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2101-2107", "2106", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2101-2107", "2107", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2108, 2109, 21099", "2108", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2108, 2109, 21099", "2109", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2108, 2109, 21099", "21099", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2110, 2111, 2112", "2110", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2110, 2111, 2112", "2111", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2110, 2111, 2112", "2112", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2113, 2114, 2115", "2113", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2113, 2114, 2115", "2114", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "2113, 2114, 2115", "2115", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "Приора", "Priora", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "Калина", "Kalina", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "Веста", "Vesta", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "Ока", "1111 Ока", "LADA (ВАЗ)")
	res = db.Exec(updateSql, "1 series", "1 серия", "BMW")
	res = db.Exec(updateSql, "2 series", "2 серия", "BMW")
	res = db.Exec(updateSql, "3 series", "3 серия", "BMW")
	res = db.Exec(updateSql, "4 series", "4 серия", "BMW")
	res = db.Exec(updateSql, "5 series", "5 серия", "BMW")
	res = db.Exec(updateSql, "CS35 Plus", "CS35PLUS", "Changan")
	res = db.Exec(updateSql, "Tiggo", "Tiggo (T11)", "Chery")
	res = db.Exec(updateSql, "Amulet", "Amulet (A15)", "Chery")
	res = db.Exec(updateSql, "Hover", "Hover H5", "Great Wall")
	res = db.Exec(updateSql, "Stepwgn", "N-WGN", "Honda")
	res = db.Exec(updateSql, "Bluebird Sylphy", "Bluebird", "Nissan")
	res = db.Exec(updateSql, "3102", "3102 «Волга»", "ГАЗ")
	res = db.Exec(updateSql, "31105", "31105 «Волга»", "ГАЗ")
	res = db.Exec(updateSql, "Патриот", "Patriot", "УАЗ")

	return nil
}

// SQL:
// WITH tires_marks AS (SELECT vendor as mark, COUNT(*) as model_count FROM `tires` GROUP BY vendor ORDER BY `mark` ASC)
// SELECT *
// FROM `car_marks` cm
// LEFT JOIN tires_marks tm ON cm.name = tm.mark;

// WITH tires_model AS (
// 	SELECT vendor AS mark, model_seria AS model, COUNT(*) AS year_count
// 	FROM tires
// 	GROUP BY vendor, model_seria
// 	ORDER BY mark
// 	)
// 	SELECT cm.*, m.*, tm.*
// 	FROM car_models m
// 	LEFT JOIN car_marks cm ON cm.id = m.mark_id
// 	LEFT JOIN tires_model tm ON tm.mark = cm.name_in_tires AND tm.model = m.name
// 	WHERE cm.name_in_tires IS NOT NULL AND m.name_in_tires IS NULL AND tm.model IS NOT NULL
// 	;
//
// SELECT *, COUNT(*) FROM `tires`
// WHERE vendor = "BMV" 
// GROUP BY model_seria, tyres_factory
// ORDER BY `tires`.`model_seria`, year DESC
// LIMIT 500;
//
// SELECT v.plate, cm.name, m.name, v.year, t.vendor as mark, t.model_seria as model, t.year, GROUP_CONCAT(DISTINCT t.tyres_factory ORDER BY t.tyres_factory ASC SEPARATOR '|') as sizes
// FROM vins v
// LEFT JOIN car_models m ON m.id = v.model_id
// LEFT JOIN car_marks cm ON cm.id = m.mark_id
// LEFT JOIN tires t ON t.vendor = cm.name_in_tires AND t.model_seria = m.name_in_tires AND t.year = v.year
// WHERE m.name_in_tires IS NOT NULL AND t.year is NOT NULL
// GROUP BY v.plate
// LIMIT 500;
