package uow

import (
	"context"

	friendshipRepo "github.com/XuanHieuHo/go-assignment/repositories/friendship"
	userRepo "github.com/XuanHieuHo/go-assignment/repositories/user"
	"gorm.io/gorm"
)

type UnitOfWork interface {
	Do(ctx context.Context, fn func(uow UnitOfWork) error) error
	UserRepo() userRepo.UserRepository
	FriendshipRepo() friendshipRepo.FriendshipRepository
}

type UnitOfWorkImpl struct {
	db *gorm.DB
	userRepo userRepo.UserRepository
	friendshipRepo friendshipRepo.FriendshipRepository
}

func NewUnitOfWorkImpl(db *gorm.DB) UnitOfWork {
	return &UnitOfWorkImpl{db: db}
}

func (u *UnitOfWorkImpl) Do(ctx context.Context, fn func(uow UnitOfWork) error) error {
	return u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		uow := &UnitOfWorkImpl{db: tx}
		return fn(uow)
	})
}

func (u *UnitOfWorkImpl) UserRepo() userRepo.UserRepository {
	if u.userRepo == nil {
		u.userRepo = userRepo.NewUserRepositoryImpl(u.db)
	}
	return u.userRepo
}

func (u *UnitOfWorkImpl) FriendshipRepo() friendshipRepo.FriendshipRepository {
	if u.friendshipRepo == nil {
		u.friendshipRepo = friendshipRepo.NewFriendshipRepository(u.db)
	}
	return u.friendshipRepo
}

