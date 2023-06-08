package domain

const (
	AliasAlreadyExists   = "custom alias already exists"
	ShortenedURLNotFound = "shortened URL not found"
)

type Url struct {
	Id        string
	Alias     string
	Original  string
	Shortened string
	Visits    int64
}

func (Url) TableName() string {
	return "public.url"
}
