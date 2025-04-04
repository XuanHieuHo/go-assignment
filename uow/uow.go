package uow

import (
	friendshipRepo "github.com/XuanHieuHo/go-assignment/repositories/friendship"
	userRepo "github.com/XuanHieuHo/go-assignment/repositories/user"
	"gorm.io/gorm"
)

type UnitOfWork interface {
	UserRepo() userRepo.UserRepository
	FriendshipRepo() friendshipRepo.FriendshipRepository
}

type UnitOfWorkImpl struct {
	db *gorm.DB
	userRepo userRepo.UserRepository
	friendshipRepo friendshipRepo.FriendshipRepository
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

