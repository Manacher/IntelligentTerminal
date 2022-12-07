package models

import "time"

type User struct {
	ID          int
	AccountID   int
	Nickname    string
	Avatar      string
	CreatedTime time.Time
}
