package storage

import (
	"errors"
	"github.com/mrChex/urlshorter/backend/model"
)

var ErrNotFound = errors.New("not found")

type Storage interface {
	PutLink(url string) (model.Link, error)
	GetLinkByID(id int64) (model.Link, error)
}
