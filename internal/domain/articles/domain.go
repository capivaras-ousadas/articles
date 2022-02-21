package articles

import (
	"context"
	"time"
)

// ArticleBuilder build a new article validating its fields
type ArticleBuilder struct {
	article *Article
}

func (b ArticleBuilder) ID(id string) ArticleBuilder {
	b.article.ID = id
	return b
}

func (b ArticleBuilder) Title(title string) ArticleBuilder {
	b.article.Title = title
	return b
}

func (b ArticleBuilder) Subtitle(subtitle string) ArticleBuilder {
	b.article.Subtitle = subtitle
	return b
}

func (b ArticleBuilder) Content(content []Block) ArticleBuilder {
	b.article.Content = content
	return b
}

func (b ArticleBuilder) Banner(banner string) ArticleBuilder {
	b.article.Banner = banner
	return b
}

func (b ArticleBuilder) Tags(tags []string) ArticleBuilder {
	b.article.Tags = tags
	return b
}

func (b ArticleBuilder) Author(author string) ArticleBuilder {
	b.article.Author = author
	return b
}

func (b ArticleBuilder) Draft(draft bool) ArticleBuilder {
	b.article.Draft = draft
	return b
}

func (b ArticleBuilder) CreatedAt(createdAt time.Time) ArticleBuilder {
	b.article.CreatedAt = createdAt
	return b
}

func (b ArticleBuilder) UpdatedAt(updatedAt time.Time) ArticleBuilder {
	b.article.UpdatedAt = updatedAt
	return b
}

func (b ArticleBuilder) PublishedAt(publishedAt time.Time) ArticleBuilder {
	b.article.PublishedAt = publishedAt
	return b
}

func (b ArticleBuilder) Build(ctx context.Context) (*Article, error) {
	if b.article.Title == "" {
		return nil, ErrArticleTitleRequired
	}

	if b.article.Author == "" {
		return nil, ErrArticleAuthorRequired
	}

	if b.article.CreatedAt.IsZero() {
		b.article.CreatedAt = time.Now()
	}

	if b.article.UpdatedAt.IsZero() {
		b.article.UpdatedAt = time.Now()
	}
	return b.article, nil
}
