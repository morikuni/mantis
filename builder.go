package mantis

import (
	"net/http"
)

type Builder struct {
	Filter  Filter
	Adapter Adapter
}

func (b *Builder) service(s Service) Service {
	if b.Filter == nil {
		return s
	}
	return ServiceFunc(func(r *http.Request) (Response, error) {
		return b.Filter.Filter(r, s)
	})
}

func (b *Builder) Use(fs ...Filter) {
	f := ComposeFilters(fs...)
	if b.Filter != nil {
		f = ComposeFilters(b.Filter, f)
	}
	b.Filter = f
}

func (b *Builder) UseFunc(funcs ...func(r *http.Request, s Service) (Response, error)) {
	fs := make([]Filter, len(funcs))
	for i, f := range funcs {
		fs[i] = FilterFunc(f)
	}
	b.Use(fs...)
}

func (b *Builder) Serve(s Service) http.Handler {
	adapter := b.Adapter
	if adapter == nil {
		adapter = DefaultAdapter
	}
	return adapter.Adapt(b.service(s))
}

func (b *Builder) ServeFunc(f func(r *http.Request) (Response, error)) http.Handler {
	return b.Serve(ServiceFunc(f))
}

var DefaultBuilder *Builder = &Builder{}

func Use(fs ...Filter) {
	DefaultBuilder.Use(fs...)
}

func UseFunc(funcs ...func(r *http.Request, s Service) (Response, error)) {
	DefaultBuilder.UseFunc(funcs...)
}

func Serve(s Service) http.Handler {
	return DefaultBuilder.Serve(s)
}

func ServeFunc(f func(r *http.Request) (Response, error)) http.Handler {
	return DefaultBuilder.ServeFunc(f)
}
