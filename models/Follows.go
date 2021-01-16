package models

import (
	"time"
)

type Follow struct {
	FollowerID  int   `gorm:"primaryKey;"`
	Follower    Actor `gorm:"foreignKey:FollowerID;references:ID;"`
	FollowingID int   `gorm:"primaryKey;"`
	Following   Actor `gorm:"foreignKey:FollowingID;references:ID;"`
	CreatedAt   time.Time
	Role        string
}
