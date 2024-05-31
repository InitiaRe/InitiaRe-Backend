package repository

import (
	"InitiaRe-website/internal/school/entity"
	"context"
)

type IRepository interface {
	GetListPaging(ctx context.Context, queries map[string]interface{}) ([]*entity.School, error)
	Count(ctx context.Context, queries map[string]interface{}) (int, error)
}
