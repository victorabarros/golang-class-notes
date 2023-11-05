package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ricardoerikson/sgg/internal/data/mem"
	"github.com/ricardoerikson/sgg/internal/service"
	"github.com/ricardoerikson/sgg/pkg/codes"
	"github.com/ricardoerikson/sgg/pkg/model"
	"github.com/stretchr/testify/suite"
)

type serviceSuite struct {
	suite.Suite
	service service.ArticleService
}

func (s *serviceSuite) SetupTest() {
	repo := mem.NewArticleRepository()
	s.service = service.NewArticleService(repo)
}

func (s *serviceSuite) TestAddArticle() {
	ctx := context.TODO()
	article := model.Article{
		Title: "article 1",
	}

	newArticle, err := s.service.CreateArticle(ctx, article)
	s.Assert().Nil(err)
	s.Assert().NotNil(newArticle)
	s.Assert().Equal("article 1", newArticle.Title)

	list, err := s.service.ListArticles(ctx, "")
	s.Assert().Nil(err)
	s.Assert().Len(list, 1)
}

func (s *serviceSuite) TestAttachSingleImageToArticle() {
	ctx := context.TODO()
	article := model.Article{
		Title: "article 1",
	}

	newArticle, err := s.service.CreateArticle(ctx, article)
	s.Assert().Nil(err)

	newImage := model.Image{Path: "http://www.example.com/image.jpg"}
	_, err = s.service.AttachImage(ctx, newArticle.ID, newImage)
	s.Assert().Nil(err)

	// use the id returned when the article was created
	retrievedArticle, err := s.service.RetrieveArticleByID(ctx, newArticle.ID)
	s.Assert().Nil(err)

	list, err := s.service.ListArticles(ctx, "")
	s.Assert().Nil(err)
	s.Assert().Len(list, 1)

	s.Assert().Equal(1, len(retrievedArticle.Images))
	s.Assert().Equal("http://www.example.com/image.jpg", retrievedArticle.Images[0].Path)
}

func (s *serviceSuite) TestValidateAmountOfImages() {
	article := model.Article{
		Title: "article with images",
	}
	images := []string{
		"http://www.example.com/image1.jpg",
		"http://www.example.com/image2.jpg",
		"http://www.example.com/image3.jpg",
	}

	ctx := context.TODO()
	newArticle, err := s.service.CreateArticle(ctx, article)
	s.Assert().Nil(err)

	for _, imagePath := range images {
		image := model.Image{
			Path: imagePath,
		}
		_, err = s.service.AttachImage(ctx, newArticle.ID, image)
		s.Assert().Nil(err)
	}

	newImage := model.Image{Path: "http://www.example.com/image4.jpg"}
	_, err = s.service.AttachImage(ctx, newArticle.ID, newImage)
	s.Assert().Error(err)
	s.Assert().True(errors.Is(err, codes.ErrNumberOfImagesExceeded))

	retrievedArticle, err := s.service.RetrieveArticleByID(ctx, newArticle.ID)
	s.Assert().Nil(err)
	s.Assert().Len(retrievedArticle.Images, 3)
}

func TestInMemoryDBTestSuite(t *testing.T) {
	suite.Run(t, new(serviceSuite))
}
