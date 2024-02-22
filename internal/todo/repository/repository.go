package repository

import (
	"context"
	"fmt"

	"InitiaRe-website/constant"
	"InitiaRe-website/internal/todo/entity"

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

func (r *repo) Create(ctx context.Context, obj *entity.Todo) (*entity.Todo, error) {
	result := r.db.Create(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *repo) CreateMany(ctx context.Context, objs []*entity.Todo) ([]*entity.Todo, error) {
	result := r.db.Create(objs)
	if result.Error != nil {
		return nil, result.Error
	}
	return objs, nil
}

func (r *repo) Update(ctx context.Context, obj *entity.Todo) (*entity.Todo, error) {
	result := r.db.Updates(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *repo) UpdateMany(ctx context.Context, objs []*entity.Todo) (int, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for _, v := range objs {
		if err := tx.Updates(v).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return len(objs), nil
}

func (r *repo) Delete(ctx context.Context, id int) (int, error) {
	result := r.db.Delete(&entity.Todo{}, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (r *repo) DeleteMany(ctx context.Context, ids []int) (int, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for _, id := range ids {
		if err := tx.Delete(&entity.Todo{}, id).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return len(ids), nil
}

func (r *repo) Count(ctx context.Context, queries map[string]interface{}) (int, error) {
	var count int64
	if err := r.initQuery(ctx, queries).Select("count(1)").Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *repo) GetById(ctx context.Context, id int) (*entity.Todo, error) {
	record := &entity.Todo{}
	result := r.db.Find(&record, id).Limit(1)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (r *repo) GetOne(ctx context.Context, queries map[string]interface{}) (*entity.Todo, error) {
	record := &entity.Todo{}
	query := r.initQuery(ctx, queries)
	result := query.Limit(1).Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (r *repo) GetList(ctx context.Context, queries map[string]interface{}) ([]*entity.Todo, error) {
	records := []*entity.Todo{}
	query := r.initQuery(ctx, queries)
	if err := query.Scan(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (r *repo) GetListPaging(ctx context.Context, queries map[string]interface{}) ([]*entity.Todo, error) {
	records := []*entity.Todo{}

	page := conversion.ReadInterface(queries, "page", constant.DEFAULT_PAGE).(int)
	size := conversion.ReadInterface(queries, "size", constant.DEFAULT_SIZE).(int)

	query := r.initQuery(ctx, queries)

	if err := query.Offset(int((page - 1) * size)).Limit(size).Scan(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (r *repo) initQuery(ctx context.Context, queries map[string]interface{}) *gorm.DB {
	query := r.db.Model(&entity.Todo{})
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
		"*",
	)
	return query
}

func (r *repo) sort(query *gorm.DB, queries map[string]interface{}) *gorm.DB {
	sortBy := conversion.ReadInterface(queries, "sort_by", "").(string)
	orderBy := conversion.ReadInterface(queries, "order_by", constant.DEFAULT_SORT_ORDER).(string)

	switch sortBy {
	default:
		query = query.Order("id " + orderBy)
	}
	return query
}

func (r *repo) filter(query *gorm.DB, queries map[string]interface{}) *gorm.DB {

	tbName := (&entity.Todo{}).TableName()
	fromDate := conversion.ReadInterface(queries, "from_date", 0).(int)
	toDate := conversion.ReadInterface(queries, "to_date", 0).(int)
	createdBy := conversion.ReadInterface(queries, "created_by", 0).(int)

	if createdBy != 0 {
		query = query.Where(fmt.Sprintf("\"%s\".created_by = ?", tbName), createdBy)
	}
	if fromDate != 0 {
		query = query.Where(fmt.Sprintf("\"%s\".created_at >= timestamp(?)", tbName), conversion.FormatUnixToString(int64(fromDate), conversion.DD_MM_YYYY_HH_MM_SS))
	}
	if toDate != 0 {
		query = query.Where(fmt.Sprintf("\"%s\".created_at < timestamp(?)", tbName), conversion.FormatUnixToString(int64(toDate), conversion.DD_MM_YYYY_HH_MM_SS))
	}
	return query
}
