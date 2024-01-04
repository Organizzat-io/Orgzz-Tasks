package models

import "time"

type Task struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	DeadLine time.Time
}
