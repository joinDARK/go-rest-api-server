package models

import (
	"time"
)

type User struct {
	ID          int64
	Email       string
	EncPassword string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}