package entities

import "time"

type (
	InsertCockroachDto struct {
		Id        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
		Amount    uint32    `json:"number"`
		CreatedAt time.Time `json:"createdAt"`
	}

	Cockroach struct {
		Id        uint32    `json:"id"`
		Amount    uint32    `json:"number"`
		CreatedAt time.Time `json:"createdAt"`
	}

	CockroachPushNotificationDto struct {
		Title        string `json:"title"`
		Amount       uint32 `json:"amount"`
		ReportedTime string `json:"createdAt"`
	}
)
