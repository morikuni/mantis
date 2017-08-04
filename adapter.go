package mantis

import (
	"log"
	"net/http"
)

type Adapter interface {
	Adapt(s Service) http.Handler
}

type defaultAdapter struct{}

func (da defaultAdapter) Adapt(s Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := s.Serve(r)
		if err != nil {
			log.Println("error while serving:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}
		if err := res.WriteTo(w); err != nil {
			log.Println("error while writing response:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}
	})
}

var DefaultAdapter Adapter = defaultAdapter{}

func Adapt(s Service) http.Handler {
	return DefaultAdapter.Adapt(s)
}
