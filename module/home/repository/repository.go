package repository

import (
	"context"
	"strings"

	"challenge.haraj.com.sa/kraicklist/model"
	"challenge.haraj.com.sa/kraicklist/module/home"
	"challenge.haraj.com.sa/kraicklist/pkg/filesearch"
	"challenge.haraj.com.sa/kraicklist/pkg/wrapper"
	"github.com/sirupsen/logrus"
)

type SearchRepository struct {
	logger     *logrus.Logger
	filesearch filesearch.Filesearch
}

// NewSearchRepository is a constructor.
func NewSearchRepository(logger *logrus.Logger, fileSearch filesearch.Filesearch) home.HomeRepo {
	return &SearchRepository{logger, fileSearch}
}

func (s SearchRepository) Search(ctx context.Context, stringToSearch string) (searchResult []model.Record, err error) {
	if strings.TrimSpace(stringToSearch) == "" {
		err = wrapper.ErrInvalidRequest
		return
	}

	searchResult = s.filesearch.Search(stringToSearch)

	return
}
