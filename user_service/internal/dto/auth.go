package dto

type LoginDTO struct {
	Email    string `form:"email" json:"email" validate:"required_without=Phone,omitempty"`
	Phone    string `form:"phone" json:"phone" validate:"required_without=Email,omitempty"`
	Password string `form:"password" json:"password" validate:"required"`
}
type AuthResponseDTO struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user_controller"`
}

type RegisterDTO struct {
	FirstName   string `form:"firstname" json:"firstname" validate:"required"`
	Surname     string `form:"surname" json:"surname" validate:"required"`
	DateOfBirth string `form:"date_of_birth" json:"date_of_birth" validate:"required"`
	Gender      int    `form:"gender" json:"gender" validate:"required"`
	Phone       string `form:"phone" json:"phone" validate:"required"`
	Email       string `form:"email" json:"email" validate:"required_without=Email,omitempty"`
	Password    string `form:"password" json:"password" validate:"required"`
}
type RegisterResponseDTO struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user_controller"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}
