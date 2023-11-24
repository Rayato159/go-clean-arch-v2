package repositories

import "github.com/Rayato159/go-clean-arch-v2/cockroach/repositories/entities"

type CockroachRepository interface {
	InsertCockroachData(in entities.InsertCockroachDto) error
}
