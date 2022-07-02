package apiserver

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Gorm connect (default)
func ConnectGorm(dsn string, logLevel string) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Настроить уровень логирования в командной строке
	if logLevel == "debug" {
		db.Config.Logger = logger.Default.LogMode(logger.Info)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// SetMaxIdleConns устанавливает максимальное количество соединений в пуле бездействия.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns устанавливает максимальное количество открытых соединений с БД.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime устанавливает максимальное время повторного использования соединения.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

// sql connect (not used)
func ConnectSql(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
