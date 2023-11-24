package repositories

import (
	"github.com/Rayato159/go-clean-arch-v2/cockroach/entities"
	"gorm.io/gorm"
)

type cockroachPostgresRepository struct {
	db *gorm.DB
}

func NewCockroachPostgresRepository(db *gorm.DB) CockroachRepository {
	return &cockroachPostgresRepository{db: db}
}

func (r *cockroachPostgresRepository) InsertCockroachData(
	in entities.InsertCockroachDto,
) error {
	return nil
}
