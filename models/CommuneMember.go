package models

import (
	"time"
)

type CommuneMember struct {
	CommuneID int `gorm:"primaryKey"`
	Commune   Commune
	AccountID int `gorm:"primaryKey"`
	Account   Account
	CreatedAt time.Time
	Role      string
}
