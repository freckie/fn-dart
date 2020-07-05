package utils

import (
	"time"
)

func MakeDate(days int) string {
	now := time.Now()
	now = now.AddDate(0, 0, days)
	return now.Format("20060102")
}
