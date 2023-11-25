package usecases

import (
	"time"

	"github.com/Rayato159/go-clean-arch-v2/cockroach/entities"
	"github.com/Rayato159/go-clean-arch-v2/cockroach/models"
	"github.com/Rayato159/go-clean-arch-v2/cockroach/repositories"
)

type cockroachUsecaseImpl struct {
	cockroachRepository repositories.CockroachRepository
	cockroachMessaging  repositories.CockroachMessaging
}

func NewCockroachUsecaseImpl(
	cockroachRepository repositories.CockroachRepository,
	cockroachMessaging repositories.CockroachMessaging,
) CockroachUsecase {
	return &cockroachUsecaseImpl{
		cockroachRepository: cockroachRepository,
		cockroachMessaging:  cockroachMessaging,
	}
}

func (u *cockroachUsecaseImpl) CockroachDataProcessing(in *models.AddCockroachData) error {
	insertCockroachData := &entities.InsertCockroachDto{
		Amount: in.Amount,
	}

	if err := u.cockroachRepository.InsertCockroachData(insertCockroachData); err != nil {
		return err
	}

	pushCockroachData := &entities.CockroachPushNotificationDto{
		Title:        "Cockroach Detected 🪳 !!!",
		Amount:       in.Amount,
		ReportedTime: time.Now().Local().Format("2006-01-02 15:04:05"),
	}

	if err := u.cockroachMessaging.PushNotification(pushCockroachData); err != nil {
		return err
	}

	return nil
}
