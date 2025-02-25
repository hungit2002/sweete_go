package models

import "time"

type Friend struct {
	UserID    int        `gorm:"column:user_id;primary_key;" json:"user_id"`
	FriendID  int        `gorm:"column:friend_id;primary_key;" json:"friend_id"`
	Status    int        `gorm:"column:status;" json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Users *User `gorm:"foreignKey:id;references:friend_id;" json:"user"`
}

func (Friend) TableName() string {
	return "friends"
}

type QueryFriendParam struct {
	UserID    int
	FriendID  int
	Status    int
	DeletedAt bool

	SwapUserIDFriendID
	LimitParam
	SortByParam
	SelectParam
	Preload *PreloadFriend
}

type SwapUserIDFriendID struct {
	UserID   int
	FriendID int
}
type PreloadFriend struct {
	Users *PreloadFriendUsers
}
type PreloadFriendUsers struct {
	Query *QueryUserParam
}
