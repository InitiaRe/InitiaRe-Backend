package repository

import (
	"context"
	"fmt"

	"github.com/Ho-Minh/InitiaRe-website/constant"
	categoryEntity "github.com/Ho-Minh/InitiaRe-website/internal/category/entity"
	"github.com/vukyn/go-kuery/konversion"
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

func (r *repo) Create(ctx context.Context, obj *categoryEntity.Category) (*categoryEntity.Category, error) {
	result := r.db.Create(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *repo) CreateMany(ctx context.Context, objs []*categoryEntity.Category) (int, error) {
	result := r.db.Create(objs)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (r *repo) Update(ctx context.Context, obj *categoryEntity.Category) (*categoryEntity.Category, error) {
	result := r.db.Updates(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *repo) UpdateMany(ctx context.Context, objs []*categoryEntity.Category) (int, error) {
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
	result := r.db.Delete(&categoryEntity.Category{}, id)
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
		if err := tx.Delete(&categoryEntity.Category{}, id).Error; err != nil {
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
	if err := r.initQuery(ctx, queries).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *repo) GetById(ctx context.Context, id int) (*categoryEntity.Category, error) {
	record := &categoryEntity.Category{}
	result := r.db.Find(&record, id).Limit(1)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (r *repo) GetOne(ctx context.Context, queries map[string]interface{}) (*categoryEntity.Category, error) {
	record := &categoryEntity.Category{}
	query := r.initQuery(ctx, queries)
	result := query.Limit(1).Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (r *repo) GetList(ctx context.Context, queries map[string]interface{}) ([]*categoryEntity.Category, error) {
	records := []*categoryEntity.Category{}
	query := r.initQuery(ctx, queries)
	if err := query.Scan(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (r *repo) GetListPaging(ctx context.Context, queries map[string]interface{}) ([]*categoryEntity.Category, error) {
	records := []*categoryEntity.Category{}

	page := konversion.ReadInterface(queries, "page", constant.DEFAULT_PAGE).(int)
	size := konversion.ReadInterface(queries, "size", constant.DEFAULT_SIZE).(int)

	query := r.initQuery(ctx, queries)

	if err := query.Offset(int((page - 1) * size)).Limit(size).Scan(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (r *repo) initQuery(ctx context.Context, queries map[string]interface{}) *gorm.DB {
	query := r.db.Model(&categoryEntity.Category{})
	query = r.join(query, queries)
	query = r.filter(query, queries)
	query = r.sort(query, queries)
	return query
}

func (r *repo) join(query *gorm.DB, queries map[string]interface{}) *gorm.DB {

	query = query.
		Joins("left join \"initiaRe_article_category\" irac on \"initiaRe_category\".id = irac.category_id")

	query = query.Select(
		"\"initiaRe_category\".*",
	)
	return query
}

func (r *repo) sort(query *gorm.DB, queries map[string]interface{}) *gorm.DB {
	sortBy := konversion.ReadInterface(queries, "sort_by", "").(string)
	orderBy := konversion.ReadInterface(queries, "order_by", constant.DEFAULT_SORT_ORDER).(string)

	switch sortBy {
	default:
		query = query.Order("\"initiaRe_category\".id " + orderBy)
	}
	return query
}

func (r *repo) filter(query *gorm.DB, queries map[string]interface{}) *gorm.DB {

	tbName := (&categoryEntity.Category{}).TableName()
	articleId := konversion.ReadInterface(queries, "article_id", 0).(int)
	createdBy := konversion.ReadInterface(queries, "created_by", 0).(int)
	fromDate := konversion.ReadInterface(queries, "from_date", 0).(int)
	toDate := konversion.ReadInterface(queries, "to_date", 0).(int)

	if createdBy != 0 {
		query = query.Where(fmt.Sprintf("\"%s\".created_by = ?", tbName), createdBy)
	}
	if articleId != 0 {
		query = query.Where("irac.article_id = ?", articleId)
	}
	if fromDate != 0 {
		query = query.Where(fmt.Sprintf("\"%s\".created_at >= timestamp(?)", tbName), konversion.FormatUnixToString(fromDate, konversion.DD_MM_YYYY_HH_MM_SS))
	}
	if toDate != 0 {
		query = query.Where(fmt.Sprintf("\"%s\".created_at < timestamp(?)", tbName), konversion.FormatUnixToString(toDate, konversion.DD_MM_YYYY_HH_MM_SS))
	}
	return query
}
