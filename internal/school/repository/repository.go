package repository

import (
	"InitiaRe-website/constant"
	"InitiaRe-website/internal/school/entity"
	"context"
	"fmt"

	"github.com/vukyn/kuery/conversion"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func InitRepo(db *gorm.DB) IRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) GetListPaging(ctx context.Context, queries map[string]interface{}) ([]*entity.School, error) {
	records := []*entity.School{}

	page := conversion.ReadInterface(queries, "page", constant.DEFAULT_PAGE).(int)
	size := conversion.ReadInterface(queries, "size", constant.DEFAULT_SIZE).(int)

	query := r.initQuery(ctx, queries)

	if err := query.Offset(int((page - 1) * size)).Limit(size).Scan(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (r *repo) initQuery(ctx context.Context, queries map[string]interface{}) *gorm.DB {
	query := r.db.Debug().Model(&entity.School{})
	query = r.join(query, queries)
	query = r.column(query, queries)
	query = r.filter(query, queries)
	query = r.sort(query, queries)
	return query
}

func (r *repo) join(query *gorm.DB, queries map[string]interface{}) *gorm.DB {
	return query
}

func (r *repo) column(query *gorm.DB, queries map[string]interface{}) *gorm.DB {
	query = query.Select(
		"\"initiaRe_school\".*",
	)
	return query
}

func (r *repo) sort(query *gorm.DB, queries map[string]interface{}) *gorm.DB {
	sortBy := conversion.ReadInterface(queries, "sort_by", "").(string)
	orderBy := conversion.ReadInterface(queries, "order_by", constant.DEFAULT_SORT_ORDER).(string)

	switch sortBy {
	default:
		query = query.Order("\"initiaRe_school\".id " + orderBy)
	}
	return query
}

func (r *repo) filter(query *gorm.DB, queries map[string]interface{}) *gorm.DB {
	tbName := (&entity.School{}).TableName()
	statusId := conversion.ReadInterfaceV2(queries, "status_id", 0)

	if statusId != 0 {
		query = query.Where(fmt.Sprintf("\"%s\".status_id = ?", tbName), statusId)
	}
	return query
}

func (r *repo) Count(ctx context.Context, queries map[string]interface{}) (int, error) {
	var count int64
	if err := r.initQuery(ctx, queries).Select("count(1)").Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
