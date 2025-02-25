package models

import "time"

type Friend struct {
	UserID    int        `gorm:"column:user_id;primary_key;" json:"user_id"`
	FriendID  int        `gorm:"column:role_id;primary_key;" json:"friend_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (Friend) TableName() string {
	return "friends"
}

type QueryFriendParam struct {
	UserID    int
	FriendID  int
	Status    int
	DeletedAt bool

	LimitParam
	SortByParam
	SelectParam
	Preload *PreloadFriend
}

type PreloadFriend struct{}
