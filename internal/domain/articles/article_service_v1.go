package articles

import "context"

type ArticleServiceV1 struct {
	repo ArticleRepository
}

func (s ArticleServiceV1) CreateArticle(ctx context.Context, article *Article) error {
	panic("not implemented") // TODO: Implement
}

func (s ArticleServiceV1) UpdateArticle(ctx context.Context, article *Article) error {
	panic("not implemented") // TODO: Implement
}

func (s ArticleServiceV1) GetArticleByID(ctx context.Context, id string) (*Article, error) {
	panic("not implemented") // TODO: Implement
}

func (s ArticleServiceV1) GetArticlesSamples(ctx context.Context, whence int64, pageSize int64) ([]*Article, error) {
	panic("not implemented") // TODO: Implement
}

func (s ArticleServiceV1) GetArticlesFromAuthor(ctx context.Context, authorID string, whence int64, pageSize int64) ([]*Article, error) {
	panic("not implemented") // TODO: Implement
}
