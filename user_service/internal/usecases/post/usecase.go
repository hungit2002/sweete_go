package post

import (
	"sweete/user_service/internal/domain/repositories"
)

type UseCase struct {
	postRepository repositories.IPostRepository
}

func New(
	postRepository repositories.IPostRepository,
) *UseCase {
	return &UseCase{
		postRepository: postRepository,
	}
}
