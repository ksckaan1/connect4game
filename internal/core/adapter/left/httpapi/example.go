package httpapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ExampleHttpAPI struct {
	router  *chi.Mux
	pattern string
}

func NewExampleHttpAPI() *ExampleHttpAPI {
	return &ExampleHttpAPI{
		router:  chi.NewRouter(),
		pattern: "/example",
	}
}

func (h *ExampleHttpAPI) WithPattern(p string) *ExampleHttpAPI {
	h.pattern = p
	return h
}

func (h *ExampleHttpAPI) App() (string, http.Handler) {
	h.router.Get("/", h.root)
	return h.pattern, h.router
}

func (h *ExampleHttpAPI) root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
