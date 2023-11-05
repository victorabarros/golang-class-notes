package http_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/moshenahmias/failure"
	"github.com/ricardoerikson/sgg/internal/cfg"
	"github.com/ricardoerikson/sgg/internal/data/mem"
	"github.com/ricardoerikson/sgg/internal/server"
	"github.com/ricardoerikson/sgg/internal/service"
	"github.com/ricardoerikson/sgg/internal/transport/entities"
	httpt "github.com/ricardoerikson/sgg/internal/transport/http"
	"github.com/ricardoerikson/sgg/internal/utils"
	"github.com/ricardoerikson/sgg/pkg/codes"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type articleSuite struct {
	suite.Suite
	handler              http.Handler
	futureExpirationDate string
	pastExpirationDate   string
}

func (s *articleSuite) SetupTest() {
	config := new(cfg.Configuration)
	if err := cleanenv.ReadEnv(config); err != nil {
		log.Error().Err(err).Msg("could not read environment variables")
		return
	}

	repo := mem.NewArticleRepository()
	service := service.NewArticleService(repo)
	handler := httpt.NewArticleHandler(service, config)

	newServer := server.NewHTTPServer(context.TODO(), handler, config)
	svr := newServer.(*server.HTTPServer)
	s.handler = svr.API.Handler
	s.futureExpirationDate = time.Now().Add(time.Hour * 24 * 7).Format("2006-01-02")
	s.pastExpirationDate = time.Now().Add(-time.Hour * 24 * 7).Format("2006-01-02")
}

func (s *articleSuite) TestCreateArticleWithInvalidData() {

	testCases := []struct {
		body          string
		expectedError error
	}{
		{body: fmt.Sprintf(`{"title":"","description":"d","expirationDate":"%s"}`, s.futureExpirationDate), expectedError: codes.ErrTitleIsRequired},
		{body: fmt.Sprintf(`{"title":"a","description":"","expirationDate":"%s"}`, s.futureExpirationDate), expectedError: codes.ErrDescriptionIsRequired},
		{body: fmt.Sprintf(`{"title":"a","description":"%s","expirationDate":"%s"}`, utils.GenerateString(4001), s.futureExpirationDate), expectedError: codes.ErrDescriptionIsTooLong},
		{body: `{"title":"a","description":"d","expirationDate":""}`, expectedError: codes.ErrExpirationDateIsRequired},
		{body: `{"title":"a","description":"d","expirationDate":""}`, expectedError: codes.ErrExpirationDateIsRequired},
		{body: fmt.Sprintf(`{"title":"a","description":"d","expirationDate":"%s"}`, s.pastExpirationDate), expectedError: codes.ErrInvalidExpirationDate},
	}

	for _, tc := range testCases {
		s.T().Run(failure.Message(tc.expectedError), func(t *testing.T) {
			assert := assert.New(t)
			b := []byte(tc.body)

			r := httptest.NewRequest(http.MethodPost, "/api/v1/articles", bytes.NewReader(b))
			w := httptest.NewRecorder()

			s.handler.ServeHTTP(w, r)

			m := map[string]string{}
			err := readResponse(w, &m)
			assert.Nil(err)
			assert.Equal(failure.Message(tc.expectedError), m["message"])
			field, err := failure.Field(tc.expectedError, "status")
			status, ok := field.(int)
			assert.True(ok)
			assert.Equal(status, w.Code)
		})
	}

}

func (s *articleSuite) TestCreateArticleWithoutImages() {
	b := []byte(fmt.Sprintf(`{"title":"a","description":"desc","expirationDate":"%s"}`, s.futureExpirationDate))

	r := httptest.NewRequest(http.MethodPost, "/api/v1/articles", bytes.NewReader(b))
	w := httptest.NewRecorder()

	s.handler.ServeHTTP(w, r)
	s.Assert().Equal(http.StatusCreated, w.Code)

	article := new(entities.ArticleCreatedResponse)
	err := readResponse(w, article)
	s.Assert().Nil(err)
	s.Assert().Equal("a", article.Title)
	s.Assert().Equal("desc", article.Description)
	s.Assert().Equal(s.futureExpirationDate, article.ExpirationDate)
	s.Assert().Len(article.Images, 0)
}

func (s *articleSuite) TestCreateArticleWithExpiredDate() {
	b := []byte(fmt.Sprintf(`{"title":"a","description":"desc","expirationDate":"%s"}`, s.futureExpirationDate))

	r := httptest.NewRequest(http.MethodPost, "/api/v1/articles", bytes.NewReader(b))
	w := httptest.NewRecorder()

	s.handler.ServeHTTP(w, r)
	s.Assert().Equal(http.StatusCreated, w.Code)

	article := new(entities.ArticleCreatedResponse)
	err := readResponse(w, article)
	s.Assert().Nil(err)
	s.Assert().Equal("a", article.Title)
	s.Assert().Equal("desc", article.Description)
	s.Assert().Equal(s.futureExpirationDate, article.ExpirationDate)
	s.Assert().Len(article.Images, 0)
}

func (s *articleSuite) TestCreateArticleWithSmallImage() {
	b := []byte(fmt.Sprintf(`{"title":"a","description":"desc","expirationDate":"%s","images":[{"path":"%s"}]}`, s.futureExpirationDate, smallImage))

	r := httptest.NewRequest(http.MethodPost, "/api/v1/articles", bytes.NewReader(b))
	w := httptest.NewRecorder()

	s.handler.ServeHTTP(w, r)
	s.Assert().Equal(http.StatusCreated, w.Code)

	article := new(entities.ArticleCreatedResponse)
	err := readResponse(w, article)
	s.Assert().Nil(err)
	s.Assert().Equal("a", article.Title)
	s.Assert().Equal("desc", article.Description)
	s.Assert().Equal(s.futureExpirationDate, article.ExpirationDate)
	s.Assert().Len(article.Images, 1)
	s.Assert().Equal(smallImage, article.Images[0].Path)
}

func (s *articleSuite) TestCreateArticleWithLargeImage() {
	b := []byte(fmt.Sprintf(`{"title":"a","description":"desc","expirationDate":"%s","images":[{"path":"%s"}]}`, s.futureExpirationDate, largeImage))

	r := httptest.NewRequest(http.MethodPost, "/api/v1/articles", bytes.NewReader(b))
	w := httptest.NewRecorder()

	s.handler.ServeHTTP(w, r)
	s.Assert().Equal(http.StatusBadRequest, w.Code)

	response := map[string]string{}
	err := readResponse(w, &response)
	s.Assert().Nil(err)
	s.Assert().Equal(failure.Message(codes.ErrImageSizeExceeded), response["message"])
}

func (s *articleSuite) TestCreateArticleInvalidImage() {
	b := []byte(fmt.Sprintf(`{"title":"a","description":"desc","expirationDate":"%s","images":[{"path":"%s"}]}`, s.futureExpirationDate, invalidImage))

	r := httptest.NewRequest(http.MethodPost, "/api/v1/articles", bytes.NewReader(b))
	w := httptest.NewRecorder()

	s.handler.ServeHTTP(w, r)
	s.Assert().Equal(http.StatusBadRequest, w.Code)

	response := map[string]string{}
	err := readResponse(w, &response)
	s.Assert().Nil(err)
	s.Assert().Equal(failure.Message(codes.ErrInvalidImageType), response["message"])
}

func (s *articleSuite) TestCreateArticleExceededImageNumber() {
	b := []byte(fmt.Sprintf(`
  {
    "title":"a",
    "description":"desc",
    "expirationDate":"%s",
    "images":[ {"path":"%s"}, {"path":"%s"}, {"path":"%s"}, {"path":"%s"}
    ]
  }`, s.futureExpirationDate, invalidImage, invalidImage, invalidImage, invalidImage))

	r := httptest.NewRequest(http.MethodPost, "/api/v1/articles", bytes.NewReader(b))
	w := httptest.NewRecorder()

	s.handler.ServeHTTP(w, r)
	s.Assert().Equal(http.StatusBadRequest, w.Code)

	response := map[string]string{}
	err := readResponse(w, &response)
	s.Assert().Nil(err)
	s.Assert().Equal(failure.Message(codes.ErrNumberOfImagesExceeded), response["message"])
}

func (s *articleSuite) TestListArticles() {
	r := httptest.NewRequest(http.MethodGet, "/api/v1/articles", nil)
	w := httptest.NewRecorder()
	s.handler.ServeHTTP(w, r)

	s.Assert().Equal(http.StatusOK, w.Code)

	response := []entities.ArticleCreatedResponse{}
	err := readResponse(w, &response)
	s.Assert().Nil(err)
	s.Assert().Len(response, 0)

	b := []byte(fmt.Sprintf(`{"title":"a","description":"desc","expirationDate":"%s"}`, s.futureExpirationDate))

	for i := 1; i <= 4; i++ {
		r := httptest.NewRequest(http.MethodPost, "/api/v1/articles", bytes.NewReader(b))
		w := httptest.NewRecorder()

		s.handler.ServeHTTP(w, r)
		s.Assert().Equal(http.StatusCreated, w.Code)

		r = httptest.NewRequest(http.MethodGet, "/api/v1/articles", nil)
		w = httptest.NewRecorder()
		s.handler.ServeHTTP(w, r)
		s.Assert().Equal(http.StatusOK, w.Code)

		response := []entities.ArticleCreatedResponse{}
		err := readResponse(w, &response)
		s.Assert().Nil(err)
		s.Assert().Len(response, i)
	}
}

func TestArticleHTTP(t *testing.T) {
	suite.Run(t, new(articleSuite))
}
