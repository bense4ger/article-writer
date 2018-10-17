package model

import (
	"fmt"
	"strings"
)

// Article models a blog article
type Article struct {
	ID          string
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Tags        []string `json:"tags"`
	CreatedDate string   `json:"createdDate"`
	UpdateDate  string   `json:"updateDate"`
}

// SetID creates the ID from the Title and CreatedDate
func (a *Article) SetID() {
	rpl := strings.NewReplacer(" ", "")
	t := rpl.Replace(a.Title)

	a.ID = fmt.Sprintf("%s_%s", t, a.CreatedDate)
}
