package model

// Article models a blog article
type Article struct {
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Tags        []string `json:"tags"`
	CreatedDate string   `json:"createdDate"`
	UpdateDate  string   `json:"updateDate"`
}
