package database

import (
	"github.com/hoangndst/cloud-observability-backend/auth/password"
	"github.com/hoangndst/cloud-observability-backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(connection, defaultUser, defaultPass string, strength int, createDefaultUserIfNotExist bool) (*GormDatabase, error) {
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&model.User{}); err != nil {
		return nil, err
	}
	var userCount *int64
	*userCount = 0
	db.Find(new(model.User)).Count(userCount)
	if createDefaultUserIfNotExist && *userCount == 0 {
		db.Create(&model.User{Name: defaultUser, Pass: password.CreatePassword(defaultPass, strength), Admin: true})
	}

	return &GormDatabase{DB: db}, nil
}

// GormDatabase is a wrapper for the gorm framework.
type GormDatabase struct {
	DB *gorm.DB
}
