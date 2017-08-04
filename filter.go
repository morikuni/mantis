package mantis

import (
	"net/http"
)

type Filter interface {
	Filter(r *http.Request, s Service) (Response, error)
}

type FilterFunc func(r *http.Request, s Service) (Response, error)

func (f FilterFunc) Filter(r *http.Request, s Service) (Response, error) {
	return f(r, s)
}

func ComposeFilters(filters ...Filter) Filter {
	l := len(filters)
	switch l {
	case 0:
		return FilterFunc(func(r *http.Request, s Service) (Response, error) {
			return s.Serve(r)
		})
	case 1:
		return filters[0]
	default:
		head := filters[0]
		tail := ComposeFilters(filters[1:]...)
		return FilterFunc(func(r *http.Request, s Service) (Response, error) {
			return head.Filter(r, ServiceFunc(func(r *http.Request) (Response, error) {
				return tail.Filter(r, s)
			}))
		})
	}
}
