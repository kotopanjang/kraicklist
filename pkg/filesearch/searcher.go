package filesearch

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	. "challenge.haraj.com.sa/kraicklist/model"
)

type Finder struct {
	records  []Record
	FilePath string
}

func NewFinder(filepath string) *Finder {
	return &Finder{[]Record{}, filepath}
}

func (s *Finder) Load() error {
	// open file
	file, err := os.Open(s.FilePath)
	if err != nil {
		return fmt.Errorf("unable to open source file due: %v", err)
	}
	defer file.Close()
	// read as gzip
	reader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("unable to initialize gzip reader due: %v", err)
	}
	// read the reader using scanner to contstruct records
	var records []Record
	cs := bufio.NewScanner(reader)
	for cs.Scan() {
		var r Record
		err = json.Unmarshal(cs.Bytes(), &r)
		if err != nil {
			continue
		}
		records = append(records, r)
	}
	s.records = records

	return nil
}

func (s *Finder) Search(query string) []Record {
	result := []Record{}
	for _, record := range s.records {
		if strings.Contains(strings.ToLower(record.Title), strings.ToLower(query)) || strings.Contains(strings.ToLower(record.Content), strings.ToLower(query)) {
			result = append(result, record)
		}
	}
	return result
}
