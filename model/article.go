package model

// Article models a blog article
type Article struct {
	Title       string
	Content     string
	Tags        []string
	CreatedDate string
	UpdateDate  string
}
