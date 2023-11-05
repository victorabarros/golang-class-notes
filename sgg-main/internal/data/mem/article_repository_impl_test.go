package mem_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/ricardoerikson/sgg/internal/data/mem"
	"github.com/ricardoerikson/sgg/pkg/codes"
	"github.com/ricardoerikson/sgg/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type memTestSuite struct {
	suite.Suite
}

// TestAddArticles tests if the articles are being added to the database.
func (s *memTestSuite) TestAddArticles() {
	ctx := context.TODO()

	repo := mem.NewArticleRepository()

	articles := []model.Article{
		{Title: "T1", Description: "D1"},
		{Title: "T2", Description: "D2"},
		{Title: "T3", Description: "D3"},
	}

	count, err := repo.Count(ctx)
	s.Assert().Nil(err)
	s.Assert().Equal(0, count)

	for i, a := range articles {
		s.T().Run(a.Title, func(t *testing.T) {
			assert := assert.New(t)
			newArticle, err := repo.Create(ctx, a)
			assert.NotNil(newArticle)
			assert.Nil(err)

			count, err := repo.Count(ctx)
			assert.Nil(err)
			assert.Equal(i+1, count)
		})
	}

	list, err := repo.List(ctx, "")
	s.Assert().Nil(err)
	s.Assert().Len(list, 3)
}

func (s *memTestSuite) TestRetrieveArticleByID() {
	testCases := []struct {
		article             model.Article
		expectedTitle       string
		expectedDescription string
	}{
		{
			article:             model.Article{Title: "T1", Description: "D1"},
			expectedTitle:       "T1",
			expectedDescription: "D1",
		},
		{
			article:             model.Article{Title: "T2", Description: "D2"},
			expectedTitle:       "T2",
			expectedDescription: "D2",
		},
		{
			article:             model.Article{Title: "T3", Description: "D3"},
			expectedTitle:       "T3",
			expectedDescription: "D3",
		},
	}

	ctx := context.TODO()
	repo := mem.NewArticleRepository()

	for _, tc := range testCases {
		s.T().Run(tc.expectedTitle, func(t *testing.T) {
			assert := assert.New(t)
			a, err := repo.Create(ctx, tc.article)
			assert.Nil(err)
			assert.Equal(tc.expectedTitle, a.Title)
			assert.Equal(tc.expectedDescription, a.Description)
		})
	}

	list, err := repo.List(ctx, "")
	s.Assert().Nil(err)
	s.Assert().Len(list, 3)
}

func (s *memTestSuite) TestAssertItWontChangeExternally() {
	ctx := context.TODO()
	repo := mem.NewArticleRepository()

	article := model.Article{
		Title:       "inserted title",
		Description: "a description",
	}

	newArticle, err := repo.Create(ctx, article)
	s.Assert().Nil(err)
	newArticle.Title = "changed title"

	retrievedArticle, err := repo.Retrieve(ctx, newArticle.ID)
	s.Assert().Nil(err)
	s.Assert().NotEqual(retrievedArticle.Title, newArticle.Title)
	s.Assert().Equal("changed title", newArticle.Title)
	s.Assert().Equal("inserted title", retrievedArticle.Title)
}

func (s *memTestSuite) TestUpdateArticle() {
	ctx := context.TODO()
	repo := mem.NewArticleRepository()
	article := model.Article{
		Title:       "inserted title",
		Description: "a description",
	}

	newArticle, err := repo.Create(ctx, article)
	s.Assert().Nil(err)

	newArticle.Title = "changed title"
	updatedArticle, err := repo.Update(ctx, *newArticle)
	s.Assert().Nil(err)

	retrievedArticle, err := repo.Retrieve(ctx, newArticle.ID)
	s.Assert().Nil(err)
	s.Assert().Equal(newArticle.ID.String(), updatedArticle.ID.String())
	s.Assert().Equal("changed title", retrievedArticle.Title)
	list, err := repo.List(ctx, "")
	s.Assert().Nil(err)
	s.Assert().Len(list, 1)
}

// TestDeleteArticle asserts if articles are being deleted
func (s *memTestSuite) TestDeleteArticle() {
	ctx := context.TODO()
	repo := mem.NewArticleRepository()
	article1 := model.Article{
		Title:       "first title",
		Description: "a description",
	}
	article2 := model.Article{
		Title:       "second title",
		Description: "a description",
	}
	count, err := repo.Count(ctx)
	s.Assert().Nil(err)
	s.Assert().Equal(0, count)

	newArticle1, err := repo.Create(ctx, article1)
	count, err = repo.Count(ctx)
	s.Assert().Nil(err)
	s.Assert().Equal(1, count)

	newArticle2, err := repo.Create(ctx, article2)
	count, err = repo.Count(ctx)
	s.Assert().Nil(err)
	s.Assert().Equal(2, count)

	err = repo.Delete(ctx, newArticle1.ID)
	s.Assert().Nil(err)
	count, err = repo.Count(ctx)
	s.Assert().Nil(err)
	s.Assert().Equal(1, count)

	remainingArticle, err := repo.Retrieve(ctx, newArticle2.ID)
	s.Assert().Nil(err)
	s.Assert().Equal("second title", remainingArticle.Title)

	err = repo.Delete(ctx, newArticle2.ID)
	s.Assert().Nil(err)
	count, err = repo.Count(ctx)
	s.Assert().Nil(err)
	s.Assert().Equal(0, count)
}

// TestTryToRetrieveANonExistingArticle asserts that an error is returned if
// we try to retrieve a non-existing article
func (s *memTestSuite) TestTryToRetrieveANonExistingArticle() {
	ctx := context.TODO()
	repo := mem.NewArticleRepository()

	article, err := repo.Retrieve(ctx, uuid.New())
	s.Assert().Nil(article)
	s.Assert().Error(err)
	s.Assert().True(errors.Is(err, codes.ErrArticleNotFound))
}

// TestTryToUpdateANonExistingArticle asserts that an error is returned if
// we try to update a non-existing article
func (s *memTestSuite) TestTryToUpdateANonExistingArticle() {
	ctx := context.TODO()
	repo := mem.NewArticleRepository()

	count, err := repo.Count(ctx)
	s.Assert().Nil(err)
	s.Assert().Equal(0, count)

	article, err := repo.Update(ctx, model.Article{})
	s.Assert().Nil(article)
	s.Assert().Error(err)
	s.Assert().True(errors.Is(err, codes.ErrArticleNotFound))

	// check that the number of records is unchanged
	count, err = repo.Count(ctx)
	s.Assert().Nil(err)
	s.Assert().Equal(0, count)
}

func TestInMemoryDBTestSuite(t *testing.T) {
	suite.Run(t, new(memTestSuite))
}
