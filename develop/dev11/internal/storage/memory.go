package storage

import (
	"dev11/internal/models"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	hoursDay = time.Duration(24)
	daysWeek = time.Duration(7)
)

type MemoryStorage struct {
	sync.RWMutex
	Data map[string]map[string]models.Event
}

func NewMemoryStorage() (*MemoryStorage, error) {
	m := &MemoryStorage{
		Data: make(map[string]map[string]models.Event),
	}
	return m, nil
}

func (m *MemoryStorage) CreateEvent(event models.Event, userId string) (string, error) {
	m.Lock()
	defer m.Unlock()

	data, ok := m.Data[userId]
	if !ok {
		data = make(map[string]models.Event)
	}

	eventId := uuid.New().String()
	event.Id = eventId
	data[eventId] = event
	m.Data[userId] = data
	return eventId, nil
}

func (m *MemoryStorage) DeleteEvent(event models.Event, userId string) error {
	m.Lock()
	defer m.Unlock()

	data, ok := m.Data[userId]
	if !ok {
		return errors.New("user doesn't exist")
	}

	_, ok = data[event.Id]
	if !ok {
		return errors.New("event doesn't exist")
	}
	delete(data, event.Id)
	m.Data[userId] = data
	return nil
}

func (m *MemoryStorage) UpdateEvent(event models.Event, userId string) error {
	m.Lock()
	defer m.Unlock()

	data, ok := m.Data[userId]
	if !ok {
		return errors.New("user doesn't exist")
	}
	_, ok = data[event.Id]
	if !ok {
		return errors.New("event doesn't exist")
	}
	data[event.Id] = event
	m.Data[userId] = data
	return nil
}

func (m *MemoryStorage) GetEventForDay(userId string, date time.Time) ([]models.Event, error) {
	m.Lock()
	defer m.Unlock()

	data, ok := m.Data[userId]
	if !ok {
		return nil, errors.New("user doesn't exist")
	}

	var events []models.Event
	for _, v := range data {
		if date.Truncate(hoursDay * time.Hour).Equal(v.Date.Truncate(hoursDay * time.Hour)) {
			events = append(events, v)
		}
	}

	return events, nil
}

func (m *MemoryStorage) GetEventForWeek(userId string, startDate time.Time) ([]models.Event, error) {
	m.Lock()
	defer m.Unlock()

	data, ok := m.Data[userId]
	if !ok {
		return nil, errors.New("user doesn't exist")
	}
	lastDay := startDate.AddDate(0, 0, 7)
	var events []models.Event
	for _, v := range data {
		if (startDate.Before(v.Date) || startDate.Equal(v.Date)) && v.Date.Before(lastDay) {
			events = append(events, v)
		}
	}

	return events, nil
}

func (m *MemoryStorage) GetEventForMonth(userId string, date time.Time) ([]models.Event, error) {
	m.Lock()
	defer m.Unlock()
	data, ok := m.Data[userId]
	if !ok {
		return nil, errors.New("user doesn't exist")
	}

	var events []models.Event
	month := date.Month()
	year := date.Year()
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, 0)

	for _, v := range data {
		if (firstDay.Before(v.Date) || firstDay.Equal(v.Date)) && v.Date.Before(lastDay) {
			events = append(events, v)
		}
	}

	return events, nil
}
