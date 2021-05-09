package usecase

import (
	"context"
	"errors"
	"strings"

	"challenge.haraj.com.sa/kraicklist/module/home"
	"challenge.haraj.com.sa/kraicklist/pkg/wrapper"
	"github.com/sirupsen/logrus"
)

type SearchUsecase struct {
	logger *logrus.Logger
	repo   home.HomeRepo
}

func NewSearchUsecase(log *logrus.Logger, repo home.HomeRepo) home.HomeUsecase {
	return &SearchUsecase{logger: log, repo: repo}
}

func (u SearchUsecase) SearchFromFile(ctx context.Context, stringToSearch string) (wrap wrapper.Result) {
	if strings.TrimSpace(stringToSearch) == "" {
		return wrapper.Result{IsError: true, Err: errors.New(string(wrapper.StatUnexpectedError)), Data: nil}
	}

	result, err := u.repo.Search(ctx, stringToSearch)
	if err != nil {
		return wrapper.Result{IsError: true, Err: err, Data: nil}
	}

	return wrapper.Result{IsError: false, Err: nil, Data: result}
}
