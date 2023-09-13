package usecase

import (
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
