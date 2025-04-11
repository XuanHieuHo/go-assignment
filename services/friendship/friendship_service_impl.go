package services

import (
	"github.com/XuanHieuHo/go-assignment/uow"
)

type FriendshipServiceImpl struct {
	uow uow.UnitOfWork
}

func NewFriendshipServiceImpl(uow uow.UnitOfWork) FriendshipService {
	return &FriendshipServiceImpl{uow: uow}
}
