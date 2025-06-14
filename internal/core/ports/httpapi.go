package ports

import "net/http"

type HttpAPIModulePort interface {
	App() (string, http.Handler)
}
