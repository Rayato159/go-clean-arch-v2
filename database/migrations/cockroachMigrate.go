package main

import (
	"github.com/Rayato159/go-clean-arch-v2/cockroach/repositories/entities"
	"github.com/Rayato159/go-clean-arch-v2/config"
	"github.com/Rayato159/go-clean-arch-v2/database"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	cockroachMigrate(db)
}

func cockroachMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Cockroach{})
	db.GetDb().CreateInBatches([]entities.Cockroach{
		{Amount: 1},
		{Amount: 2},
		{Amount: 2},
		{Amount: 5},
		{Amount: 3},
	}, 10)
}
