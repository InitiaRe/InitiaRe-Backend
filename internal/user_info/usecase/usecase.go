package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/user_info/repository"
)

type usecase struct {
	repo repository.IRepository
}

func InitUsecase(repo repository.IRepository) IUseCase {
	return &usecase{
		repo: repo,
	}
}


func (u *usecase) Enable(ctx context.Context, userId int) error {
	return nil
}

func (u *usecase) Disable(ctx context.Context, userId int) error {
	return nil
}
