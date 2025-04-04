package models

import "time"

type FriendShip struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	User       User      `json:"user" gorm:"foreignKey:UserID"`
	FriendID   uint      `json:"friend_id" gorm:"not null"`
	Friend     User      `json:"friend" gorm:"foreignKey:FriendID"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
