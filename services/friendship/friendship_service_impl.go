package services

import "github.com/XuanHieuHo/go-assignment/uow"

type FriendshipServiceImpl struct {
	Manager *uow.Manager
}

func NewFriendshipServiceImpl(manager *uow.Manager) FriendshipService {
	return &FriendshipServiceImpl{Manager: manager}
}
