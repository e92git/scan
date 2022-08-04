package helper

import (
	"time"
)

func StrToTime(datetime string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", datetime, time.Local)
}

func NowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
