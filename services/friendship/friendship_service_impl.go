package services

import (
	"gorm.io/gorm"
)

type FriendshipServiceImpl struct {
	db *gorm.DB
}

func NewFriendshipServiceImpl(db *gorm.DB) FriendshipService {
	return &FriendshipServiceImpl{db: db}
}
