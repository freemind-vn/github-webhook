package service

import (
	"errors"
	"net/http"

	"freemind.com/webhook/helper"
	"freemind.com/webhook/service/index"
)

type Service struct {
	http.Handler
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		index.Handle(w, r)
	default:
		helper.WriteHttpError(w, http.StatusNotFound, errors.New("not found"))
	}
}
