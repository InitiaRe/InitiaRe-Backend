package repository

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/models"
)

func (t *todoRepo) Create(ctx context.Context, todo *models.Todo) (int64, error) {
	result := t.db.Create(todo)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
