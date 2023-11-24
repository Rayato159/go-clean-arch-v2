package repositories

import "github.com/Rayato159/go-clean-arch-v2/cockroach/entities"

type CockroachRepository interface {
	InsertCockroachData(in entities.InsertCockroachDto) error
}
