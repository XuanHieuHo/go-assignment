package services

import (
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserServiceImpl(db *gorm.DB) UserService {
	return &UserServiceImpl{db: db}
}
