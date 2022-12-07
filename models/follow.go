package models

import "time"

type Follow struct {
	ID           int
	SubscribedID int
	FollowerID   int
	CreatedTime  time.Time
}
