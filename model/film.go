package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Film struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	Cover    string
	Type     int
	CategoryID int
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
	gorm.Model
	DownType string
	FilmID uint
	Title    string
	Url      string
	Password string
}

type HtmlOnline struct {
	ID   int `gorm:"primary_key"`
	FilmID uint
	ParentUrl string
	Name string
	PlayUrl string
}
