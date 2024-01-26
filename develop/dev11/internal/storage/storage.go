package storage

import (
	"dev11/internal/models"
	"time"
)

type Storager interface {
	CreateEvent(event models.Event, userId string) (string, error)
	UpdateEvent(event models.Event, userId string) error
	DeleteEvent(event models.Event, userId string) error
	GetEventForDay(userId string, date time.Time) ([]models.Event, error)
	GetEventForWeek(userId string, startDate time.Time) ([]models.Event, error)
	GetEventForMonth(userId string, date time.Time) ([]models.Event, error)
}
