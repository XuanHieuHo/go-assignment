package repositories

import (
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(tx *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: tx}
}
