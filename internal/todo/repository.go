package todo

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/models"
)

type Repository interface {
	Create(ctx context.Context, todo *models.Todo) (int64, error)
	// Update(ctx context.Context, news *models.News) (*models.News, error)
	// GetNewsByID(ctx context.Context, newsID uuid.UUID) (*models.NewsBase, error)
	// Delete(ctx context.Context, newsID uuid.UUID) error
	// GetTodos(ctx context.Context, pq *utils.PaginationQuery) (*models.NewsList, error)
	// SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*models.NewsList, error)
}
