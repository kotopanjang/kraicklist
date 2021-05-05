package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"challenge.haraj.com.sa/kraicklist/pkg/filesearch"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	homeHandler "challenge.haraj.com.sa/kraicklist/module/home/handler"
	homeRepo "challenge.haraj.com.sa/kraicklist/module/home/repository"
	homeUsecase "challenge.haraj.com.sa/kraicklist/module/home/usecase"
)

type HTTPServer struct {
	Port       string
	Logger     *logrus.Logger
	ExitSignal chan os.Signal
	Router     *mux.Router
	Searcher   filesearch.Filesearch
}

func (s *HTTPServer) Start() {
	defer close(s.ExitSignal)

	s.composeHandlerAPI()
	s.composeHandlerView()

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", s.Port),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 60,
		Handler:      s.Router,
	}

	go func() {
		s.Logger.Infof("Application is running on port %s", s.Port)
		s.Logger.Error(httpServer.ListenAndServe())
	}()

	<-s.ExitSignal

	httpServer.Shutdown(context.Background())
}

func (s *HTTPServer) composeHandlerView() {
	//set handler index
	fs := http.FileServer(http.Dir("./module/home/view"))
	s.Router.Handle("/", fs)
}

func (s *HTTPServer) composeHandlerAPI() {
	// initiate Repo
	RepoHome := homeRepo.NewSearchRepository(s.Logger, s.Searcher)

	//initiate usecase
	usecaseHome := homeUsecase.NewSearchUsecase(s.Logger, RepoHome)

	//initiate http handlers
	homeHandler.NewHomeHttpHandler(s.Logger, usecaseHome, s.Router)
}
