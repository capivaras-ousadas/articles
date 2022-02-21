package articles

import "time"

// Article is a domain model
type Article struct {
	ID          string
	Title       string
	Subtitle    string
	Content     []Block
	Banner      string
	Tags        []string
	Author      string
	Draft       bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt time.Time
}
