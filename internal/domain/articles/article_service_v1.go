package articles

import "context"

type ArticleServiceV1 struct {
	repo ArticleRepository
}

func NewArticleService(repo ArticleRepository) *ArticleServiceV1 {
	return &ArticleServiceV1{
		repo: repo,
	}
}

func (s ArticleServiceV1) CreateArticle(ctx context.Context, article *Article) error {
	// check if article already exists
	if _, err := s.GetArticleByID(ctx, article.ID); err == nil {
		return ErrArticleAlreadyExists
	}

	// create article
	if err := s.repo.Store(ctx, article); err != nil {
		return err
	}

	return nil
}

func (s ArticleServiceV1) UpdateArticle(ctx context.Context, article *Article) error {
	panic("not implemented") // TODO: Implement
}

func (s ArticleServiceV1) GetArticleByID(ctx context.Context, id string) (*Article, error) {
	if id == "" {
		return nil, ErrArticleIDNotProvided
	}

	return s.repo.FindByID(ctx, id)
}

func (s ArticleServiceV1) GetArticlesSamples(ctx context.Context, whence int64, pageSize int64) ([]*Article, error) {
	panic("not implemented") // TODO: Implement
}

func (s ArticleServiceV1) GetArticlesFromAuthor(ctx context.Context, authorID string, whence int64, pageSize int64) ([]*Article, error) {
	panic("not implemented") // TODO: Implement
}
