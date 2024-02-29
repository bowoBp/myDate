package repository

import (
	"context"
	"errors"
	"github.com/bowoBp/myDate/internal/domains"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	UserRepo struct {
		db *gorm.DB
	}

	UserRepoInterface interface {
		AddUser(
			ctx context.Context,
			user *domains.Users,
		) (*domains.Users, error)
		GetUserByID(
			ctx context.Context,
			id uint,
		) (user *domains.Users, err error)
		GetUserByEmail(
			ctx context.Context,
			email string,
		) (user domains.Users, err error)
		UpdateSelectedField(
			ctx context.Context,
			user *domains.Users,
			fields ...string,
		) (*domains.Users, error)
		UpdateIsVerify(
			ctx context.Context,
			user *domains.Users,
			updateData map[string]any,
		) (*domains.Users, error)
	}
)

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return UserRepo{db: db}
}

func (repo UserRepo) AddUser(
	ctx context.Context,
	user *domains.Users,
) (*domains.Users, error) {
	err := repo.db.WithContext(ctx).
		Create(&user).
		Error
	return user, err
}

func (repo UserRepo) GetUserByID(
	ctx context.Context,
	id uint,
) (user *domains.Users, err error) {
	err = repo.db.WithContext(ctx).
		Where("user_id = ?", id).
		First(&user).
		Error
	return user, err
}

func (repo UserRepo) GetUserByEmail(
	ctx context.Context,
	email string,
) (user domains.Users, err error) {
	err = repo.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return user, nil
	}
	return user, err
}

func (repo UserRepo) UpdateSelectedField(
	ctx context.Context,
	user *domains.Users,
	fields ...string,
) (*domains.Users, error) {
	err := repo.db.WithContext(ctx).
		Model(user).
		Select(fields).
		Omit(clause.Associations).
		Updates(*user).
		Error

	return user, err
}

func (repo UserRepo) UpdateIsVerify(
	ctx context.Context,
	user *domains.Users,
	updateData map[string]any,
) (*domains.Users, error) {
	err := repo.db.WithContext(ctx).
		Omit(clause.Associations).
		Model(&user).
		Updates(updateData).
		Error

	return user, err
}
