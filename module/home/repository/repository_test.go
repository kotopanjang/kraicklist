package repository_test

import (
	"context"
	"fmt"
	"testing"

	"challenge.haraj.com.sa/kraicklist/model"
	"challenge.haraj.com.sa/kraicklist/module/home/repository"
	fileSearchMock "challenge.haraj.com.sa/kraicklist/pkg/filesearch/mocks"
	"challenge.haraj.com.sa/kraicklist/pkg/wrapper"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSearch_Success(t *testing.T) {
	fs := new(fileSearchMock.Filesearch)
	fs.On("Search", mock.AnythingOfType("string")).Return([]model.Record{}, nil)

	repo := repository.NewSearchRepository(logrus.New(), fs)
	result, err := repo.Search(context.TODO(), "Test")

	assert.NoError(t, err, "SHOULD NOT be error")
	assert.GreaterOrEqual(t, 0, len(result), "Result should be 0 or more than 0")

	fs.AssertExpectations(t)
}

func TestSearch_Success_No_Data(t *testing.T) {
	fs := new(fileSearchMock.Filesearch)
	fs.On("Search", mock.AnythingOfType("string")).Return([]model.Record{}, nil)

	repo := repository.NewSearchRepository(logrus.New(), fs)
	result, err := repo.Search(context.TODO(), "Test")

	assert.NoError(t, err, "SHOULD NOT be error")
	assert.Equal(t, 0, len(result), "Result should be 0 or more than 0")

	fs.AssertExpectations(t)
}

func TestSearch_ErrorInvalidRequest(t *testing.T) {
	fs := new(fileSearchMock.Filesearch)
	// fs.On("Search", mock.AnythingOfType("string")).Return([]model.Record{}, wrapper.ErrInvalidRequest)

	repo := repository.NewSearchRepository(logrus.New(), fs)
	result, err := repo.Search(context.TODO(), "")

	assert.Error(t, err, "should be error")
	assert.Equal(t, wrapper.ErrInvalidRequest, err, "Should be Invalid Request Error")
	assert.Equal(t, 0, len(result), "Return should be 0")

	fs.AssertExpectations(t)
}

func TestSearch_ErrorInternalServerError(t *testing.T) {
	fs := new(fileSearchMock.Filesearch)
	fs.On("Search", mock.AnythingOfType("string")).Return([]model.Record{}, wrapper.ErrInternalServer)

	repo := repository.NewSearchRepository(logrus.New(), fs)
	result, err := repo.Search(context.TODO(), "Test")

	fmt.Println("result >> ", result)
	fmt.Println("err >> ", err)

	assert.Error(t, err, "should be error")
	assert.Equal(t, wrapper.ErrInternalServer, err, "Should be Internal Server Error")
	assert.Equal(t, 0, len(result), "Return should be 0")

	fs.AssertExpectations(t)
}
