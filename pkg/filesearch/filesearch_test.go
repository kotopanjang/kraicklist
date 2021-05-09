package filesearch_test

import (
	"testing"

	"challenge.haraj.com.sa/kraicklist/pkg/filesearch"
	"github.com/stretchr/testify/assert"
	// "challenge.haraj.com.sa/kraicklist/pkg/filesearch"
)

func TestLoad_Success(t *testing.T) {
	filePath := "../../data.gz"
	fs := filesearch.NewFinder(filePath)

	err := fs.Load()

	assert.NoError(t, err, "Load should not be error")
}

func TestLoad_Error_UnableToReadZip(t *testing.T) {
	filePath := "../../data"
	fs := filesearch.NewFinder(filePath)

	err := fs.Load()

	assert.Error(t, err, "Load should be error")
}

func TestLoad_Error_UnableToOpen(t *testing.T) {
	filePath := ""
	fs := filesearch.NewFinder(filePath)

	err := fs.Load()

	assert.Error(t, err, "Load should be error")
}

func TestSearch_Success(t *testing.T) {
	filePath := "../../data.gz"
	fs := filesearch.NewFinder(filePath)

	errLoad := fs.Load()
	datas := fs.Search("test")

	assert.NoError(t, errLoad, "Load should not be error")
	assert.Greater(t, len(datas), 0, "Should return some datas")

}
