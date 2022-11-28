package utils

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/extark/gateway_microservice/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(Cfg.DBDSN), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(
		models.CasbinRule{},
	); err != nil {
		return nil, err
	}

	return db, nil
}

func initCasbinAdapter() *gormadapter.Adapter {
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(Cfg.SQLDB, &models.CasbinRule{})

	if err != nil {
		log.Fatalln(err)
	}

	return adapter
}
