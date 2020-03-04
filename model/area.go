package model

import "time"

type Area struct {
	ID int `gorm:"primary_key"`
	Name string

	DeletedAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}