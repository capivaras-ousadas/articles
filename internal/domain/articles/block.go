package articles

import "time"

// Block represents a block of text or image.
type Block struct {
	ID        string
	Type      string
	Data      string
	Index     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
