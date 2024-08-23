package util

import "time"

func IntPtr(i int) *int {
	return &i
}

func GenerateTimePtr() *time.Time {
	currentTime := time.Now()

	return &currentTime
}
