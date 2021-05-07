package handler

import (
	"errors"
	"net/http"

	"challenge.haraj.com.sa/kraicklist/module/home"
	"challenge.haraj.com.sa/kraicklist/pkg/wrapper"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	basePath string = "/search"
)

type HomeHttpHandler struct {
	Logger  *logrus.Logger
	Usecase home.HomeUsecase
}

func NewHomeHttpHandler(logger *logrus.Logger, usecase home.HomeUsecase, router *mux.Router) {

	handler := HomeHttpHandler{
		Logger:  logger,
		Usecase: usecase,
	}

	router.HandleFunc(
		basePath,
		handler.SearchString,
	)

}

func (h HomeHttpHandler) SearchString(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	wp := wrapper.Result{}

	//get param
	q := r.URL.Query().Get("q")
	if len(q) == 0 {

		w.Write([]byte("missing search query in query params"))
		// err := errors.New("missing search query in query params")

		wp.IsError = true
		wp.Err = errors.New("missing search query in query params")
		wrapper.ResponseError(w, &wp)
		return
	}

	result := h.Usecase.SearchFromFile(ctx, q)
	wp.Data = result.Data
	wp.IsError = false

	wrapper.ResponseSuccess(w, 200, &wp, "")
	return

}
