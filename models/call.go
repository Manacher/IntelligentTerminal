package models

import "time"

type Call struct {
	ID          int
	UserID      int
	ExpiredTime time.Time
}
