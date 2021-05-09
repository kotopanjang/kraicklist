package main

import (
	"fmt"
	"os"

	"challenge.haraj.com.sa/kraicklist/pkg/filesearch"
	"challenge.haraj.com.sa/kraicklist/server"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	// get env
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	pathFileSearch := os.Getenv("PATHSEARCHFILE")
	if pathFileSearch == "" {
		fmt.Println("PATHSEARCHFILE is not set! Now using default path!")
		pathFileSearch = "data.gz"
	}

	// set logger
	logger := logrus.New()

	// set router
	router := mux.NewRouter()

	// initialize searcher
	ss := filesearch.NewFinder(pathFileSearch)
	err := ss.Load()
	if err != nil {
		logger.Fatal(err)
	}

	s := server.HTTPServer{
		Port:     port,
		Logger:   logger,
		Router:   router,
		Searcher: ss,
	}
	s.Start()

}
