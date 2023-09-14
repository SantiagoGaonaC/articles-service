package database

import (
	"products-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDatabase() (*gorm.DB, error) {
	env := config.GetEnv()

	dsn := "host=" + env.DatabaseHost +
		" user=" + env.DatabaseUser +
		" password=" + env.DatabasePassword +
		" dbname=" + env.DatabaseName +
		" port=" + env.DatabasePort +
		" sslmode=" + env.DatabaseSSLMode +
		" TimeZone=America/Bogota"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
