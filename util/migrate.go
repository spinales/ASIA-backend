package util

import (
	"fmt"
	"log"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spinales/ASIA-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MigrateDB(migrationPath string, config *Config) error {
	path := fmt.Sprintf("file:/%s", migrationPath)

	m, err := migrate.New(path, strings.ReplaceAll(config.DBURI, "postgresql", "cockroachdb"))
	if err != nil {
		return err
	}

	ver, _, _ := m.Version()
	if ver < config.DBVersion {
		if err := m.Up(); err != nil {
			return err
		}
	}

	return nil
}

func MigrateDB2(migrationPath string, config *Config) error {
	db, err := gorm.Open(postgres.Open(config.DBURI), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
		return err
	}

	db.Migrator().DropTable(models.User{}, models.Trimester{}, models.Nacionality{},
		models.Teacher{}, models.Admin{}, models.Admin{},
		models.Course{}, models.Rating{}, models.Section{}, models.Selection{}, models.SelectionRecord{}, models.Student{})
	db.AutoMigrate(models.User{}, models.Trimester{}, models.Nacionality{},
		models.Teacher{}, models.Admin{}, models.Admin{}, models.Student{},
		models.Course{}, models.Rating{}, models.Section{}, models.Selection{}, models.SelectionRecord{})

	return nil
}
