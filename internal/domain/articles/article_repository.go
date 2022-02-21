package articles

import "context"

// ArticleRepository is an interface for article repository
type ArticleRepository interface {
	Save(ctx context.Context, article *Article) error
	Update(ctx context.Context, article *Article) error
	FindByID(ctx context.Context, id string) (*Article, error)
	FindSamples(ctx context.Context, whence, limit int64) ([]*Article, error)
	FindByAuthor(ctx context.Context, authorID string, whence, pageSize int64) ([]*Article, error)
}
