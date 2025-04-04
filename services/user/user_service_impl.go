package services

import "github.com/XuanHieuHo/go-assignment/uow"

type UserServiceImpl struct {
	Manager *uow.Manager
}

func NewUserServiceImpl(manager *uow.Manager) UserService {
	return &UserServiceImpl{Manager: manager}
}
