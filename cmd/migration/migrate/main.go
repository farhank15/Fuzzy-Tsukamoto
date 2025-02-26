package main

import (
	"go-tsukamoto/config"
	"go-tsukamoto/internal/app/models"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := config.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	modelsToMigrate := models.GetModelsToMigrate()

	// Log the models being migrated
	log.Println("Migrating the following models:")
	for _, model := range modelsToMigrate {
		log.Printf("- %T\n", model)
	}

	// Migrate the schema
	err = db.AutoMigrate(modelsToMigrate...)
	if err != nil {
		log.Fatal("failed to migrate database")
	}

	// Seed data for Predicate
	models.SeedPredicates(db)

	log.Println("Database migration completed successfully")
}
