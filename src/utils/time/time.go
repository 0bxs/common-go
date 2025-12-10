package time

import "time"

func Now2Dawning() int64 {
	now := time.Now()
	today := now.Truncate(24 * time.Hour)
	return now.Sub(today).Milliseconds()
}
