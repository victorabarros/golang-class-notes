package mem

import (
	"github.com/ricardoerikson/sgg/pkg/model"
)

// inMemoryDatabase is the in-memory database for articles
type database struct {
	storage map[string]model.Article
}
