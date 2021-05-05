package filesearch

import (
	. "challenge.haraj.com.sa/kraicklist/model"
)

// Filesearch is behaviour of Search8+50
type Filesearch interface {
	Load() error
	Search(query string) ([]Record, error)
}
