package data

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

// Article 文章表结构.
type Article struct {
	BaseModel
	UserId      uint   `json:"user_id"`
	Title       string `json:"title"`
	Abstract    string `json:"abstract"`
	Content     string `json:"content"`
	IsTop       int32  `json:"is_top"`
	IsPublished int32  `json:"is_published"`
	IsDeleted   int32  `json:"is_deleted"`
}

type ArticleRepo struct {
	tableName string
	engine    *gorm.DB
}

func NewArticleRepo(engine *gorm.DB) *ArticleRepo {
	return &ArticleRepo{
		tableName: "article",
		engine:    engine,
	}
}

// Create .
func (r *ArticleRepo) Create(ctx context.Context, value *Article) error {
	return r.engine.Table(r.tableName).WithContext(ctx).Create(value).Error
}

// GetByCond .
func (r *ArticleRepo) GetByCond(ctx context.Context, cond map[string]interface{}) (*Article, error) {
	if len(cond) == 0 {
		return nil, nil
	}

	var article Article
	err := r.engine.Table(r.tableName).WithContext(ctx).Where(cond).First(&article).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &article, nil
}

// UpdateByIdAndCond 根据id和条件更新.
func (r *ArticleRepo) UpdateByIdAndCond(ctx context.Context, id uint, cond map[string]interface{}) error {
	if id <= 0 || len(cond) == 0 {
		return nil
	}

	return r.engine.Table(r.tableName).WithContext(ctx).Where("id = ?", id).Updates(cond).Error
}

// ExistByCond 根据条件判断是否存在.
func (r *ArticleRepo) ExistByCond(ctx context.Context, cond map[string]interface{}) (has bool, err error) {
	user, err := r.GetByCond(ctx, cond)
	if err != nil || user == nil {
		return false, err
	}
	return true, nil
}

func (r *ArticleRepo) FindOrderByIdAndLimit(ctx context.Context, id uint, limit int) ([]*Article, error) {
	var articles []*Article
	db := r.engine.Table(r.tableName).WithContext(ctx)
	if id > 0 {
		db = db.Debug().Where("id < ?", id)
	}

	if err := db.Order("id desc").Limit(limit).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
