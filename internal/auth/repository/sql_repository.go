package repository

import (
	"context"
	"fmt"

	"github.com/Ho-Minh/InitiaRe-website/internal/auth/entity"
	"github.com/vukyn/go-kuery/konversion"

	"gorm.io/gorm"
)

func (r *repo) Create(ctx context.Context, obj *entity.User) (*entity.User, error) {
	result := r.db.Create(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *repo) GetById(ctx context.Context, id int) (*entity.User, error) {
	record := &entity.User{}
	result := r.db.Limit(1).Find(&record, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (r *repo) GetOne(ctx context.Context, queries map[string]interface{}) (*entity.User, error) {
	record := &entity.User{}
	query := r.initQuery(ctx, queries)
	result := query.Limit(1).Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (r *repo) initQuery(ctx context.Context, queries map[string]interface{}) *gorm.DB {
	query := r.db.Model(&entity.User{})
	query = r.join(query, queries)
	query = r.filter(query, queries)
	return query
}

func (r *repo) join(query *gorm.DB, queries map[string]interface{}) *gorm.DB {
	query = query.Select(
		"*",
	)
	return query
}

func (r *repo) filter(query *gorm.DB, queries map[string]interface{}) *gorm.DB {

	userTbName := (&entity.User{}).TableName()
	email := konversion.ReadInterface(queries, "email", "").(string)

	if email != "" {
		query = query.Where(fmt.Sprintf("\"%s\".email = ?", userTbName), email)
	}
	return query
}
