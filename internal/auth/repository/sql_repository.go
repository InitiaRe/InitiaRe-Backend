package repository

import (
	"context"
	"fmt"

	"github.com/Ho-Minh/InitiaRe-website/internal/auth/entity"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils/conversion"

	"gorm.io/gorm"
)

func (r *repo) dbWithContext(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx)
}

func (r *repo) Create(ctx context.Context, obj *entity.User) (*entity.User, error) {
	result := r.dbWithContext(ctx).Create(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *repo) GetById(ctx context.Context, id int) (*entity.User, error) {
	record := &entity.User{}
	result := r.dbWithContext(ctx).Find(&record, id).Limit(1)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (r *repo) GetOne(ctx context.Context, queries map[string]interface{}) (*entity.User, error) {
	record := &entity.User{}
	query := r.initQuery(ctx, queries)
	result := query.Offset(0).Limit(1).Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (r *repo) initQuery(ctx context.Context, queries map[string]interface{}) *gorm.DB {
	query := r.dbWithContext(ctx).Model(&entity.User{})
	query = r.join(query, queries)
	query = r.filter(query, queries)
	return query
}

func (r *repo) join(query *gorm.DB, queries map[string]interface{}) *gorm.DB {
	query = query.Select(
		"users.*",
	)
	return query
}

func (r *repo) filter(query *gorm.DB, queries map[string]interface{}) *gorm.DB {

	userTbName := (&entity.User{}).TableName()
	email := conversion.GetFromInterface(queries, "email", "").(string)

	if email != "" {
		query = query.Where(fmt.Sprintf("%s.email = ?", userTbName), email)
	}
	return query
}
