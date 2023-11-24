package usecases

import "github.com/Rayato159/go-clean-arch-v2/cockroach/models"

type CockroachUsecase interface {
	CockroachDataManaging(in *models.AddCockroachData) error
}
