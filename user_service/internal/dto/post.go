package dto

type GetPostsByParamInputDTO struct {
	IDs     []int  `json:"ids"`
	Content string `json:"content"`
	Checkin string `json:"checkin"`
	UserID  int    `json:"user_id"`
}

type CreatePostInputDTO struct {
}
