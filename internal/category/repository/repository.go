package repository

import (
	"context"
	"fmt"

	"github.com/Ho-Minh/InitiaRe-website/internal/constants"
	"github.com/Ho-Minh/InitiaRe-website/internal/category/entity"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils/conversion"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) IRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, obj *entity.Category) (*entity.Category, error) {
	result := r.db.Create(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *repo) CreateMany(ctx context.Context, objs []*entity.Category) (int, error) {
	result := r.db.Create(objs)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (r *repo) Update(ctx context.Context, obj *entity.Category) (*entity.Category, error) {
	result := r.db.Updates(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *repo) UpdateMany(ctx context.Context, objs []*entity.Category) (int, error) {
	result := r.db.Updates(objs)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (r *repo) Count(ctx context.Context, queries map[string]interface{}) (int, error) {
	var count int64
	if err := r.initQuery(ctx, queries).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *repo) GetById(ctx context.Context, id int) (*entity.Category, error) {
	record := &entity.Category{}
	result := r.db.Find(&record, id).Limit(1)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (r *repo) GetOne(ctx context.Context, queries map[string]interface{}) (*entity.Category, error) {
	record := &entity.Category{}
	query := r.initQuery(ctx, queries)
	result := query.Offset(0).Limit(1).Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (r *repo) GetList(ctx context.Context, queries map[string]interface{}) ([]*entity.Category, error) {
	records := []*entity.Category{}
	query := r.initQuery(ctx, queries)
	if err := query.Scan(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (r *repo) GetListPaging(ctx context.Context, queries map[string]interface{}) ([]*entity.Category, error) {
	records := []*entity.Category{}

	page := conversion.GetFromInterface(queries, "page", constants.DEFAULT_PAGE).(int)
	size := conversion.GetFromInterface(queries, "size", constants.DEFAULT_SIZE).(int)

	query := r.initQuery(ctx, queries)

	if err := query.Offset(int((page - 1) * size)).Limit(size).Scan(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (r *repo) initQuery(ctx context.Context, queries map[string]interface{}) *gorm.DB {
	query := r.db.Model(&entity.Category{})
	query = r.join(query, queries)
	query = r.filter(query, queries)
	query = r.sort(query, queries)
	return query
}

func (r *repo) join(query *gorm.DB, queries map[string]interface{}) *gorm.DB {
	query = query.Select(
		"*",
	)
	return query
}

func (r *repo) sort(query *gorm.DB, queries map[string]interface{}) *gorm.DB {
	sortBy := conversion.GetFromInterface(queries, "sort_by", "").(string)
	orderBy := conversion.GetFromInterface(queries, "order_by", constants.DEFAULT_SORT_ORDER).(string)

	switch sortBy {
	default:
		query = query.Order("id " + orderBy)
	}
	return query
}

func (r *repo) filter(query *gorm.DB, queries map[string]interface{}) *gorm.DB {

	todoTbName := (&entity.Category{}).TableName()
	fromDate := conversion.GetFromInterface(queries, "from_date", 0).(int)
	toDate := conversion.GetFromInterface(queries, "to_date", 0).(int)
	createdBy := conversion.GetFromInterface(queries, "created_by", 0).(int)

	if createdBy != 0 {
		query = query.Where(fmt.Sprintf("%s.created_by = ?", todoTbName), createdBy)
	}
	if fromDate != 0 {
		query = query.Where(fmt.Sprintf("%s.created_at >= timestamp(?)", todoTbName), conversion.FormatUnixToString(fromDate, "YYYY-MM-DD HH:mm:ss"))
	}
	if toDate != 0 {
		query = query.Where(fmt.Sprintf("%s.created_at < timestamp(?)", todoTbName), conversion.FormatUnixToString(toDate, "YYYY-MM-DD HH:mm:ss"))
	}
	return query
}
