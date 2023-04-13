package data

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

// User 用户表结构.
type User struct {
	BaseModel
	UserName  string `json:"user_name"`
	NickName  string `json:"nick_name"`
	Uid       string `json:"uid"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	AvatarUrl string `json:"avatar_url"`
	Level     int    `json:"level"`
	IsLocked  int    `json:"is_locked"`
	IsDeleted int    `json:"is_deleted"`
}

type UserRepo struct {
	tableName string
	engine    *gorm.DB
}

func NewUserRepo(engine *gorm.DB) *UserRepo {
	return &UserRepo{
		tableName: "user",
		engine:    engine,
	}
}

// Create 创建用户.
func (r *UserRepo) Create(ctx context.Context, value *User) error {
	return r.engine.Table(r.tableName).WithContext(ctx).Create(value).Error
}

// GetByCond 根据条件查询用户.
func (r *UserRepo) GetByCond(ctx context.Context, cond map[string]interface{}) (*User, error) {
	if len(cond) == 0 {
		return nil, nil
	}

	var user User
	err := r.engine.Table(r.tableName).WithContext(ctx).Where(cond).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*User, error) {
	cond := make(map[string]interface{})
	cond["email"] = email
	cond["is_deleted"] = 0
	return r.GetByCond(ctx, cond)
}

// UpdateByIdAndCond 根据id和条件更新.
func (r *UserRepo) UpdateByIdAndCond(ctx context.Context, id uint, cond map[string]interface{}) error {
	if id <= 0 || len(cond) == 0 {
		return nil
	}

	return r.engine.Table(r.tableName).WithContext(ctx).Where("id = ?", id).Updates(cond).Error
}

// ExistByCond 根据条件判断是否存在.
func (r *UserRepo) ExistByCond(ctx context.Context, cond map[string]interface{}) (has bool, err error) {
	user, err := r.GetByCond(ctx, cond)
	if err != nil || user == nil {
		return false, err
	}
	return true, nil
}

func (r *UserRepo) ExistByEmail(ctx context.Context, email string) (has bool, err error) {
	cond := make(map[string]interface{})
	cond["email"] = email
	cond["is_deleted"] = 0
	return r.ExistByCond(ctx, cond)
}

func (r *UserRepo) FindByIds(ctx context.Context, ids []uint) ([]*User, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var users []*User
	err := r.engine.Table(r.tableName).WithContext(ctx).
		Where("id IN ?", ids).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) FindMapByIds(ctx context.Context, ids []uint) (map[uint]*User, error) {
	users, err := r.FindByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	userMap := make(map[uint]*User)
	for _, user := range users {
		userMap[user.ID] = user
	}
	return userMap, nil
}
