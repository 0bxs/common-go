package time

import (
	"time"
)

func Now2TodayEnd() int64 {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	return today.Sub(now).Milliseconds()
}

func TodayEnd() int64 {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()).UnixMilli() - 1
}

func TodayStart() int64 {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).UnixMilli() - 1
}
