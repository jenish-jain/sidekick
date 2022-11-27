package helpers

import "time"

type Time interface {
	GetDaysBetweenTimestamps(initialTimestamp time.Time, finalTimestamp time.Time) int64
}

type timeImpl struct{}

func NewTime() Time {
	return timeImpl{}
}

func (t timeImpl) GetDaysBetweenTimestamps(initialTimestamp time.Time, finalTimestamp time.Time) int64 {
	return int64(finalTimestamp.Sub(initialTimestamp).Hours() / 24)
}
