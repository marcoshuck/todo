package domain

import "gorm.io/gorm"

// MigrateModels migrates the domain models using the given DB connection.
func MigrateModels(db *gorm.DB) error {
	return db.Migrator().AutoMigrate(
		&Task{},
	)
}
