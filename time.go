package gears

import "time"

// DaysAgo judge t, if it is n days ago.
func DaysAgo(t time.Time, n int) bool {
	return time.Now().Unix()-t.Unix() < 24*60*60*int64(n)
}
