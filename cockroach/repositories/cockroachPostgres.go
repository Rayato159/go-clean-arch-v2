package repositories

import (
	"github.com/Rayato159/go-clean-arch-v2/cockroach/repositories/entities"
	"gorm.io/gorm"
)

type cockroachPostgresRepository struct {
	db *gorm.DB
}

func (r *cockroachPostgresRepository) InsertCockroachData(
	in entities.InsertCockroachDto,
) error {
	return nil
}
