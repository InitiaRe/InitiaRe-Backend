package repository

import (
	"context"
	"fmt"

	"github.com/Ho-Minh/InitiaRe-website/constant"
	"github.com/Ho-Minh/InitiaRe-website/internal/article/entity"
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

func (r *repo) Create(ctx context.Context, obj *entity.Article) (*entity.Article, error) {
	result := r.db.Create(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *repo) CreateMany(ctx context.Context, objs []*entity.Article) (int, error) {
	result := r.db.Create(objs)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (r *repo) Update(ctx context.Context, obj *entity.Article) (*entity.Article, error) {
	result := r.db.Updates(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *repo) UpdateMany(ctx context.Context, objs []*entity.Article) (int, error) {
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
	result := r.db.Delete(&entity.Article{}, id)
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
		if err := tx.Delete(&entity.Article{}, id).Error; err != nil {
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

func (r *repo) GetById(ctx context.Context, id int) (*entity.Article, error) {
	record := &entity.Article{}
	result := r.db.Find(&record, id).Limit(1)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (r *repo) GetOne(ctx context.Context, queries map[string]interface{}) (*entity.Article, error) {
	record := &entity.Article{}
	query := r.initQuery(ctx, queries)
	result := query.Limit(1).Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (r *repo) GetList(ctx context.Context, queries map[string]interface{}) ([]*entity.Article, error) {
	records := []*entity.Article{}
	query := r.initQuery(ctx, queries)
	if err := query.Scan(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (r *repo) GetListPaging(ctx context.Context, queries map[string]interface{}) ([]*entity.Article, error) {
	records := []*entity.Article{}

	page := konversion.ReadInterface(queries, "page", constant.DEFAULT_PAGE).(int)
	size := konversion.ReadInterface(queries, "size", constant.DEFAULT_SIZE).(int)

	query := r.initQuery(ctx, queries)

	if err := query.Offset(int((page - 1) * size)).Limit(size).Scan(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (r *repo) initQuery(ctx context.Context, queries map[string]interface{}) *gorm.DB {
	query := r.db.Debug().Model(&entity.Article{})
	query = r.join(query, queries)
	query = r.filter(query, queries)
	query = r.sort(query, queries)
	return query
}

func (r *repo) join(query *gorm.DB, queries map[string]interface{}) *gorm.DB {

	query = query.
		Joins("left join \"initiaRe_status\" irs on \"initiaRe_article\".status_id = irs.status_id and irs.category = 'article'").
		Joins("left join \"initiaRe_category\" irc on \"initiaRe_article\".category_id = irc.id").
		Joins("left join \"initiaRe_user\" iru on \"initiaRe_article\".created_by = iru.id")

	query = query.Select(
		"\"initiaRe_article\".*",
		"irs.status_name",
		"irc.category_name",
		"iru.email",
	)
	return query
}

func (r *repo) sort(query *gorm.DB, queries map[string]interface{}) *gorm.DB {
	sortBy := konversion.ReadInterface(queries, "sort_by", "").(string)
	orderBy := konversion.ReadInterface(queries, "order_by", constant.DEFAULT_SORT_ORDER).(string)

	switch sortBy {
	default:
		query = query.Order("\"initiaRe_article\".id " + orderBy)
	}
	return query
}

func (r *repo) filter(query *gorm.DB, queries map[string]interface{}) *gorm.DB {

	tbName := (&entity.Article{}).TableName()
	title := konversion.ReadInterface(queries, "title", "").(string)
	email := konversion.ReadInterface(queries, "email", "").(string)
	categoryIds := konversion.ReadInterface(queries, "category_ids", []int{}).([]int)
	statusId := konversion.ReadInterface(queries, "status_id", 0).(int)
	typeId := konversion.ReadInterface(queries, "type_id", 0).(int)
	fromDate := konversion.ReadInterface(queries, "from_date", 0).(int)
	toDate := konversion.ReadInterface(queries, "to_date", 0).(int)
	createdBy := konversion.ReadInterface(queries, "created_by", 0).(int)

	if title != "" {
		query = query.Where(fmt.Sprintf("\"%s\".title ilike ?", tbName), "%"+title+"%")
	}
	if email != "" {
		query = query.Where("iru.email ilike ?", "%"+email+"%")
	}
	if statusId != 0 {
		query = query.Where(fmt.Sprintf("\"%s\".status_id = ?", tbName), statusId)
	}
	if typeId != 0 {
		query = query.Where(fmt.Sprintf("\"%s\".type_id = ?", tbName), typeId)
	}
	if createdBy != 0 {
		query = query.Where(fmt.Sprintf("\"%s\".created_by = ?", tbName), createdBy)
	}
	if fromDate != 0 {
		query = query.Where(fmt.Sprintf("\"%s\".created_at >= timestamp(?)", tbName), konversion.FormatUnixToString(fromDate, konversion.DD_MM_YYYY_HH_MM_SS))
	}
	if toDate != 0 {
		query = query.Where(fmt.Sprintf("\"%s\".created_at < timestamp(?)", tbName), konversion.FormatUnixToString(toDate, konversion.DD_MM_YYYY_HH_MM_SS))
	}
	if len(categoryIds) > 0 {
		query = query.Where(fmt.Sprintf("\"%s\".category_id in ?", tbName), categoryIds)
	}
	return query
}
