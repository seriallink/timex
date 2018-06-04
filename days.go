package timex

import (
	"time"
)

// Returns date with the time portion of the day truncated
func Trunc(t time.Time) (time.Time) {
	_, offset := t.Zone()
	seconds := t.Unix()
	return time.Unix(seconds-(seconds+int64(offset))%86400, 0)
}

// Returns today's time to midnight
func Today() (time.Time) {
	return Trunc(time.Now())
}

// Returns tomorrow's time by adding one day to the given time
func Tomorrow(t time.Time) (time.Time) {
	return AddDays(t, 1)
}

// Returns tomorrow's time by subtracting one day from the given time
func Yesterday(t time.Time) (time.Time) {
	return SubDays(t, 1)
}

func Days(t time.Time) int {
	_, offset := t.Zone()
	return int((t.Unix() + int64(offset)) / int64(86400))
}

// Calculates the difference between two times in days
func DaysBetween(t1, t2 time.Time) (int) {
	return Days(t2) - Days(t1)
}

// Adds "n" days to a given time
func AddDays(t time.Time, n int) (ret time.Time) {
	ret = Trunc(t)
	if n != 0 {
		ret = time.Unix(ret.Unix() + int64(86400*n), 0)
	}
	return
}

// Subtracts "n" days to a given time
func SubDays(t time.Time, n int) (ret time.Time) {
	ret = Trunc(t)
	if n != 0 {
		ret = time.Unix(ret.Unix() - int64(86400*n), 0)
	}
	return
}
