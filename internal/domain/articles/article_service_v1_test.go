package articles

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetArticleByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockArticleRepository(ctrl)
	repo.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, nil).MaxTimes(2)

	service := NewArticleService(repo)
	var tests = []struct {
		name        string
		expectedErr error
		givenID     string
	}{
		{
			"invalid id",
			ErrArticleIDNotProvided,
			"",
		},
		{
			"valid id",
			nil,
			"1",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := service.GetArticleByID(context.Background(), tt.givenID)
			assert.Equal(t, tt.expectedErr, actual)
		})
	}
}

func TestCreateArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var tests = []struct {
		name         string
		expectedErr  error
		givenArticle *Article
		givenRepo    *MockArticleRepository
	}{
		{
			"Article already exists",
			ErrArticleAlreadyExists,
			&Article{
				ID:    "1",
				Title: "title",
			},
			setupArticleRepoReturningArticle(ctrl),
		},
		{
			"Creating new article with success",
			nil,
			&Article{
				ID:    "1",
				Title: "title",
			},
			setupSuccessMockArticleRepo(ctrl),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			service := NewArticleService(tt.givenRepo)
			actual := service.CreateArticle(context.Background(), tt.givenArticle)
			assert.Equal(t, tt.expectedErr, actual)
		})
	}
}

func setupSuccessMockArticleRepo(ctrl *gomock.Controller) *MockArticleRepository {
	repo := NewMockArticleRepository(ctrl)
	repo.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, ErrArticleNotFound).MaxTimes(1)
	repo.EXPECT().Store(gomock.Any(), gomock.Any()).Return(nil).MaxTimes(1)
	return repo
}

func setupArticleRepoReturningArticle(ctrl *gomock.Controller) *MockArticleRepository {
	repo := NewMockArticleRepository(ctrl)
	repo.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(new(Article), nil).MaxTimes(1)
	return repo
}
