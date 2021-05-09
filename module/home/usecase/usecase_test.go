package usecase_test

import (
	"context"
	"errors"
	"testing"

	"challenge.haraj.com.sa/kraicklist/model"
	repoMocks "challenge.haraj.com.sa/kraicklist/module/home/mocks"
	"challenge.haraj.com.sa/kraicklist/module/home/usecase"
	"challenge.haraj.com.sa/kraicklist/pkg/wrapper"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSearchFromFile_Success(t *testing.T) {
	repo := new(repoMocks.HomeRepo)
	repo.On("Search", mock.Anything, mock.AnythingOfType("string")).Return([]model.Record{}, nil)

	uc := usecase.NewSearchUsecase(logrus.New(), repo)
	result := uc.SearchFromFile(context.TODO(), "Test")

	assert.Equal(t, nil, result.Err, "Error should be null")
	assert.Equal(t, false, result.IsError, "IsError should be FALSE")

	repo.AssertExpectations(t)

}

func TestSearchFromFile_Error_InternalServer(t *testing.T) {
	err := wrapper.ErrInternalServer

	repo := new(repoMocks.HomeRepo)
	repo.On("Search", mock.Anything, mock.AnythingOfType("string")).Return([]model.Record{}, wrapper.ErrInternalServer)

	uc := usecase.NewSearchUsecase(logrus.New(), repo)
	result := uc.SearchFromFile(context.TODO(), "Test")

	assert.Equal(t, err, result.Err, "Error should Unexpected Error")
	assert.Equal(t, true, result.IsError, "IsError should be TRUE")

	repo.AssertExpectations(t)

}

func TestSearchFromFile_Error_InvalidRequest(t *testing.T) {
	err := errors.New(string(wrapper.StatUnexpectedError))

	repo := new(repoMocks.HomeRepo)
	// repo.On("Search", mock.Anything, mock.AnythingOfType("string")).Return([]model.Record{}, wrapper.ErrInvalidRequest)

	uc := usecase.NewSearchUsecase(logrus.New(), repo)
	result := uc.SearchFromFile(context.TODO(), "")

	assert.Equal(t, err, result.Err, "Error should Unexpected Error")
	assert.Equal(t, true, result.IsError, "IsError should be TRUE")

	repo.AssertExpectations(t)

}
