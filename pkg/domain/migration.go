package domain

import "gorm.io/gorm"

func MigrateModels(db *gorm.DB) error {
	return db.Migrator().AutoMigrate(
		&Task{},
	)
}
