package models

import "time"

type User struct {
	ID            int        `json:"id"`
	Phone         string     `json:"phone"`
	Email         string     `json:"email"`
	FullName      string     `json:"full_name"`
	Address       string     `json:"address"`
	EducationInfo string     `json:"education_info"`
	WorkInfo      string     `json:"work_info"`
	Gender        int        `json:"gender"`
	Relationship  int        `json:"relationship"`
	Dob           time.Time  `json:"dob"`
	Avatar        string     `json:"avatar"`
	Poster        string     `json:"poster"`
	Password      string     `json:"password"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}

type QueryUserParam struct {
	IDs       []int
	ID        int
	Email     string
	Phone     string
	FullName  string
	DeletedAt bool

	LimitParam
	EmailOrPhone *EmailOrPhone
	SortByParam
	SelectParam
	Preload *PreloadUser
}
type EmailOrPhone struct {
	Email string
	Phone string
}
type LimitParam struct {
	Limit  *int
	Offset *int
	Page   *int
}

type SortByParam struct {
	SortBy []string
}
type SelectParam struct {
	Select []string
}
type PreloadUser struct {
}
