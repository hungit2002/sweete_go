package dto

type GetFriendByParamInputDTO struct {
	UserID   int    `form:"user_id" json:"user_id"`
	FullName string `form:"full_name" json:"full_name"`
}
