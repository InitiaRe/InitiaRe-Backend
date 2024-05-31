package usecase

import (
	"InitiaRe-website/internal/school/models"
	"context"
)

type IUseCase interface {
	GetListPaging(ctx context.Context, params *models.RequestList) (*models.ListPaging, error)
}
