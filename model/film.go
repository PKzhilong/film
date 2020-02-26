package model

type Film struct {
	Name string
	Cover string
	Type int
	Language string
	ShowTime string
	Length int
	Area string
	TypeName string


	Director string
	Writer string
	Actors string

	Plot string
	Content string

	DownloadUrls []DownloadUrl
	HtmlOnlines []HtmlOnline
}

type DownloadUrl struct {
	DownType string
	Title string
	Url string
	Password string
}

type HtmlOnline struct {
	PlayApp string
	Items []DirItem
}

type DirItem struct {
	Name string
	Url string
}

