package models

import "time"

type Post struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	Status     int       `json:"status"`
	BackGround string    `json:"background"`
	UserID     int       `json:"user_id"`
	Feeling    int       `json:"feeling"`
	Checkin    string    `json:"checkin"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

func (Post) TableName() string {
	return "posts"
}

type QueryPostParam struct {
	IDs       []int
	ID        int
	Content   string
	Status    int
	UserID    int
	Feeling   int
	Checkin   string
	DeletedAt bool

	LimitParam
	SortByParam
	SelectParam

	Preload *PreloadPost
}

type PreloadPost struct {
}
