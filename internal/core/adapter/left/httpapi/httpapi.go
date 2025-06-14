package httpapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ksckaan1/connect4backend/internal/core/ports"
)

type HttpAPI struct {
	router *chi.Mux
	port   string
}

func New() *HttpAPI {
	return &HttpAPI{
		router: chi.NewRouter(),
		port:   ":3000",
	}
}

func (h *HttpAPI) WithPort(p string) *HttpAPI {
	h.port = p
	return h
}

func (h *HttpAPI) Mount(m ports.HttpAPIModulePort) {
	h.router.Mount(m.App())
}

func (h *HttpAPI) Start() error {
	return http.ListenAndServe(h.port, h.router)
}
