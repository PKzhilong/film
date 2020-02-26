package model

import (
	"time"
)

type Film struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	Cover    string
	Type     int
	Language string
	ShowTime string
	Length   int
	Area     string
	TypeName string

	Director string
	Writer   string
	Actors   string

	Plot    string
	Content string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	DownloadUrls []DownloadUrl
	HtmlOnlines  []HtmlOnline
}

type DownloadUrl struct {
	DownType string
	Title    string
	Url      string
	Password string
}

type HtmlOnline struct {
	PlayApp string
	Items   []DirItem
}

type DirItem struct {
	Name string
	Url  string
}
