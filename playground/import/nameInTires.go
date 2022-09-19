package main

import (
	"fmt"
	"log"
	"scan/app/apiserver"
)

// Пример скрипта для тестов
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

	// run code
	db.Exec("SELECT 1")

	fmt.Println("Успешно завершено Всё.")
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
