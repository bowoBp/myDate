package repository

import (
	"context"
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
			user *domains.User,
		) (*domains.User, error)
		GetUserByID(
			ctx context.Context,
			id uint,
		) (user *domains.User, err error)
		GetUserByEmail(
			ctx context.Context,
			email string,
		) (user *domains.User, err error)
		UpdateSelectedField(
			ctx context.Context,
			user *domains.User,
			fields ...string,
		) (*domains.User, error)
	}
)

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return UserRepo{db: db}
}

func (repo UserRepo) AddUser(
	ctx context.Context,
	user *domains.User,
) (*domains.User, error) {
	err := repo.db.WithContext(ctx).
		Create(&user).
		Error
	return user, err
}

func (repo UserRepo) GetUserByID(
	ctx context.Context,
	id uint,
) (user *domains.User, err error) {
	err = repo.db.WithContext(ctx).
		First(&user, id).
		Error
	return user, err
}

func (repo UserRepo) GetUserByEmail(
	ctx context.Context,
	email string,
) (user *domains.User, err error) {
	err = repo.db.WithContext(ctx).
		First(&user, email).
		Error
	return user, err
}

func (repo UserRepo) UpdateSelectedField(
	ctx context.Context,
	user *domains.User,
	fields ...string,
) (*domains.User, error) {
	err := repo.db.WithContext(ctx).
		Model(user).
		Select(fields).
		Omit(clause.Associations).
		Updates(*user).
		Error

	return user, err
}
