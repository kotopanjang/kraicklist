package home

import (
	"context"

	"challenge.haraj.com.sa/kraicklist/model"
	"challenge.haraj.com.sa/kraicklist/pkg/wrapper"
)

type HomeRepo interface {
	Search(ctx context.Context, stringToSearch string) (searchResult []model.Record, err error)
}

type HomeUsecase interface {
	HandleSearchFromFile(ctx context.Context, whatToSearch string) (wrap wrapper.Result)
}
