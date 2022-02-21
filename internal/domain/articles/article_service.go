package articles

import "context"

// ArticleService is an interface for article service
type ArticleService interface {
	CreateArticle(ctx context.Context, article *Article) error
	UpdateArticle(ctx context.Context, article *Article) error
	GetArticleByID(ctx context.Context, id string) (*Article, error)
	GetArticlesSamples(ctx context.Context, whence, pageSize int64) ([]*Article, error)
	GetArticlesFromAuthor(ctx context.Context, authorID string, whence, pageSize int64) ([]*Article, error)
}
