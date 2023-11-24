package repositories

import (
	"github.com/Rayato159/go-clean-arch-v2/cockroach/repositories/entities"
	"github.com/labstack/gommon/log"
)

type cockroachFCMMessaging struct{}

func NewCockroachFCMMessaging() CockroachMessaging {
	return &cockroachFCMMessaging{}
}

func (c *cockroachFCMMessaging) PushNotification(m entities.CockroachPushNotificationDto) error {
	// ... handle logic to push FCM notification here ...
	log.Info("Pushing notification to FCM")

	return nil
}
