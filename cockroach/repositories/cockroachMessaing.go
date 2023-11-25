package repositories

import "github.com/Rayato159/go-clean-arch-v2/cockroach/entities"

type CockroachMessaging interface {
	PushNotification(m *entities.CockroachPushNotificationDto) error
}
