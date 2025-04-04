package repositories

import (
	"gorm.io/gorm"
)

type FriendshipRepositoryImpl struct {
	db *gorm.DB
}

func NewFriendshipRepository(tx *gorm.DB) FriendshipRepository {
	return &FriendshipRepositoryImpl{db: tx}
}
