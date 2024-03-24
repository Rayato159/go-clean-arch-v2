package main

import (
	"github.com/Rayato159/go-clean-arch-v2/cockroach/entities"
	"github.com/Rayato159/go-clean-arch-v2/config"
	"github.com/Rayato159/go-clean-arch-v2/database"
)

func main() {
	conf := config.GetConfig()

	db := database.NewPostgresDatabase(conf)

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
