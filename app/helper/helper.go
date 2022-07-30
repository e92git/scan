package helper

import (
	"scan/app/model"
	"time"
)

func StrToTime(datetime string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", datetime, time.Local)
}

func isMiddlewareApi(role string) bool {
	return role == model.UserRoles.ShowApi
}
