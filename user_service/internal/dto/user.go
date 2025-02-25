package dto

type UserParamsDTO struct {
	FullName string `form:"full_name" json:"full_name"`
	Email    string `form:"email" json:"email"`
	Phone    string `form:"phone" json:"phone"`
	IDs      []int  `form:"ids" json:"ids"`
}

type UserIdDTO struct {
	ID int `form:"id" json:"id"`
}

type InputInviteFriendDTO struct {
	UserID   int `form:"user_id" json:"user_id"`
	FriendID int `form:"friend_id" json:"friend_id"`
	Status   int `form:"status" json:"status"`
}
