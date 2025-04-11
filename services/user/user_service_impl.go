package services

import (
	"github.com/XuanHieuHo/go-assignment/uow"
)

type UserServiceImpl struct {
	uow uow.UnitOfWork
}

func NewUserServiceImpl(uow uow.UnitOfWork) UserService {
	return &UserServiceImpl{uow: uow}
}
