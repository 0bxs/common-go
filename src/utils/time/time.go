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

func TodayEnd1(now time.Time) int64 {
	return time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()).UnixMilli() - 1
}

func TodayEnd2(ms int64) int64 {
	now := time.UnixMilli(ms)
	return time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()).UnixMilli() - 1
}

func TodayStart() int64 {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).UnixMilli() - 1
}

func TodayStart1(now time.Time) int64 {
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).UnixMilli() - 1
}

func TodayStart2(ms int64) int64 {
	now := time.UnixMilli(ms)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).UnixMilli() - 1
}

func WeekStart() int64 {
	now := time.Now()
	loc := now.Location()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	start := now.AddDate(0, 0, -(weekday - 1))
	return time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, loc).UnixMilli()
}

func WeekStart1(now time.Time) int64 {
	loc := now.Location()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	start := now.AddDate(0, 0, -(weekday - 1))
	return time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, loc).UnixMilli()
}

func WeekStart2(ms int64) int64 {
	now := time.UnixMilli(ms)
	loc := now.Location()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	start := now.AddDate(0, 0, -(weekday - 1))
	return time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, loc).UnixMilli()
}

func WeekEnd() int64 {
	now := time.Now()
	loc := now.Location()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	offset := 7 - weekday
	end := now.AddDate(0, 0, offset)
	return time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), loc).UnixMilli()
}

func WeekEnd1(now time.Time) int64 {
	loc := now.Location()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	offset := 7 - weekday
	end := now.AddDate(0, 0, offset)
	return time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), loc).UnixMilli()
}

func WeekEnd2(ms int64) int64 {
	now := time.UnixMilli(ms)
	loc := now.Location()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	offset := 7 - weekday
	end := now.AddDate(0, 0, offset)
	return time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), loc).UnixMilli()
}

func MontyStart() int64 {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).UnixMilli()
}

func MontyStart1(now time.Time) int64 {
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).UnixMilli()
}

func MontyStart2(ms int64) int64 {
	now := time.UnixMilli(ms)
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).UnixMilli()
}

func MonthEnd() int64 {
	now := time.Now()
	nextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())
	return nextMonth.Add(-time.Nanosecond).UnixMilli()
}

func MonthEnd1(now time.Time) int64 {
	nextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())
	return nextMonth.Add(-time.Nanosecond).UnixMilli()
}
func MonthEnd2(ms int64) int64 {
	now := time.UnixMilli(ms)
	nextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())
	return nextMonth.Add(-time.Nanosecond).UnixMilli()
}
