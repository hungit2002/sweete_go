package user

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (u *UseCase) GetUserDetail(ctx context.Context, id int) (*models.User, error) {
	columns := []string{"id",
		"full_name",
		"email",
		"phone",
		"avatar",
		"poster",
		"dob",
		"gender",
		"relationship",
		"work_info",
		"education_info",
		"address"}
	user, err := u.userRepository.FirstByParams(ctx, models.QueryUserParam{
		SelectParam: models.SelectParam{Select: columns},
		ID:          id,
		DeletedAt:   true,
		Preload: &models.PreloadUser{
			FriendUsers: &models.PreloadUserFriendUser{
				Query: &models.QueryFriendParam{
					SelectParam: models.SelectParam{
						Select: columns,
					},
				},
			},
		},
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// get friends

	return user, nil
}
