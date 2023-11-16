package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/article/models"
)

type IUseCase interface {
	Create(ctx context.Context, userId int, params *models.SaveRequest) (*models.Response, error)
	CreateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error)
	Update(ctx context.Context, userId int, params *models.SaveRequest) (*models.Response, error)
	UpdateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error)
	Delete(ctx context.Context, id int) (int, error)
	DeleteMany(ctx context.Context, ids []int) (int, error)
	GetById(ctx context.Context, id int) (*models.Response, error)
	GetList(ctx context.Context, params *models.RequestList) ([]*models.Response, error)
	GetApprovedArticle(ctx context.Context) (*models.ApprovedList, error)
	GetListPaging(ctx context.Context, params *models.RequestList) (*models.ListPaging, error)
	GetOne(ctx context.Context, params *models.RequestList) (*models.Response, error)
	Approve(ctx context.Context, id int) error
	Disable(ctx context.Context, id int) error
}
