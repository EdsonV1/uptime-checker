package utils

import "time"

func RepeatCheck(interval time.Duration, duration time.Duration, fn func()) {
	if duration <= 0 {
		fn()
		return
	}

	end := time.Now().Add(duration)

	for time.Now().Before(end) {
		fn()
		time.Sleep(interval)
	}
}
