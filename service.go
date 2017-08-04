package mantis

import (
	"net/http"
)

type Service interface {
	Serve(r *http.Request) (Response, error)
}

type ServiceFunc func(r *http.Request) (Response, error)

func (f ServiceFunc) Serve(r *http.Request) (Response, error) {
	return f(r)
}
