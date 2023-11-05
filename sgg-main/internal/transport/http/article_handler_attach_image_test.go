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
	"github.com/ricardoerikson/sgg/pkg/codes"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
)

type imageSuite struct {
	suite.Suite
	handler              http.Handler
	futureExpirationDate string
	pastExpirationDate   string
}

func (s *imageSuite) SetupSuite() {
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

func (s *imageSuite) TestAttachImageToArticleSucceed() {
	b := []byte(fmt.Sprintf(`
  {
    "title":"a",
    "description":"desc",
    "expirationDate":"%s",
    "images":[ {"path":"%s"}, {"path":"%s"}]
  }`, s.futureExpirationDate, smallImage, smallImage))

	r := httptest.NewRequest(http.MethodPost, "/api/v1/articles", bytes.NewReader(b))
	w := httptest.NewRecorder()

	s.handler.ServeHTTP(w, r)
	s.Assert().Equal(http.StatusCreated, w.Code)

	response := new(entities.ArticleCreatedResponse)
	err := readResponse(w, &response)
	s.Assert().Nil(err)

	b = []byte(fmt.Sprintf(`{"path":"%s"}`, smallImage))

	r = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/articles/%s/images", response.ID.String()), bytes.NewReader(b))
	w = httptest.NewRecorder()
	s.handler.ServeHTTP(w, r)

	updated := new(entities.ArticleCreatedResponse)
	s.Assert().Equal(http.StatusOK, w.Code)
	err = readResponse(w, &updated)
	s.Assert().Nil(err)
	s.Assert().Len(updated.Images, 3)
	s.Assert().Equal("a", updated.Title)
	s.Assert().Equal("desc", updated.Description)
	s.Assert().Equal(s.futureExpirationDate, updated.ExpirationDate)
}

func (s *imageSuite) TestAttachImage_ErrorWithSizeExceeded() {
	b := []byte(fmt.Sprintf(`
  {
    "title":"a",
    "description":"desc",
    "expirationDate":"%s"
  }`, s.futureExpirationDate))

	r := httptest.NewRequest(http.MethodPost, "/api/v1/articles", bytes.NewReader(b))
	w := httptest.NewRecorder()

	s.handler.ServeHTTP(w, r)
	s.Assert().Equal(http.StatusCreated, w.Code)

	response := new(entities.ArticleCreatedResponse)
	err := readResponse(w, &response)
	s.Assert().Nil(err)

	b = []byte(fmt.Sprintf(`{"path":"%s"}`, largeImage))

	r = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/articles/%s/images", response.ID.String()), bytes.NewReader(b))
	w = httptest.NewRecorder()
	s.handler.ServeHTTP(w, r)

	resp := map[string]string{}
	s.Assert().Equal(http.StatusBadRequest, w.Code)
	err = readResponse(w, &resp)
	s.Assert().Nil(err)
	s.Assert().Equal(failure.Message(codes.ErrImageSizeExceeded), resp["message"])
}

func (s *imageSuite) TestAttachImageToArticleFail() {
	b := []byte(fmt.Sprintf(`
  {
    "title":"a",
    "description":"desc",
    "expirationDate":"%s",
    "images":[{"path":"%s"},{"path":"%s"},{"path":"%s"}]
  }`, s.futureExpirationDate, smallImage, smallImage, smallImage))

	r := httptest.NewRequest(http.MethodPost, "/api/v1/articles", bytes.NewReader(b))
	w := httptest.NewRecorder()

	s.handler.ServeHTTP(w, r)
	s.Assert().Equal(http.StatusCreated, w.Code)

	response := new(entities.ArticleCreatedResponse)
	err := readResponse(w, &response)
	s.Assert().Nil(err)

	b = []byte(fmt.Sprintf(`{"path":"%s"}`, smallImage))

	r = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/articles/%s/images", response.ID.String()), bytes.NewReader(b))
	w = httptest.NewRecorder()
	s.handler.ServeHTTP(w, r)

	resp := map[string]string{}
	s.Assert().Equal(http.StatusBadRequest, w.Code)
	err = readResponse(w, &resp)
	s.Assert().Nil(err)
	s.Assert().Equal(failure.Message(codes.ErrNumberOfImagesExceeded), resp["message"])
}

func TestArticleAttachImageHTTP(t *testing.T) {
	suite.Run(t, new(imageSuite))
}
