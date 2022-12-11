package response

import (
	"terminal/models"
	"time"
)

type UserDetailResp struct {
	ID              int          `json:"id"`
	Avatar          string       `json:"avatar"`
	Nickname        string       `json:"nickname"`
	Tags            []models.Tag `json:"tags"`
	FollowerNum     int          `json:"follower_num"`
	SubscribedNum   int          `json:"subscribed_num"`
	SubscribeStatus bool         `json:"subscribe_status"`
	CreatedTime     time.Time    `json:"created_time"`
}
