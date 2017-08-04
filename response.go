package mantis

import (
	"net/http"
)

var (
	Nil Response
)

type Response struct {
	Code   int
	Header http.Header
	Body   Body
}

func (r Response) WriteTo(w http.ResponseWriter) error {
	for k, vs := range r.Header {
		w.Header()[k] = vs
	}
	if r.Code != 0 {
		w.WriteHeader(r.Code)
	}
	if r.Body == nil {
		return nil
	}
	return r.Body.WriteTo(w)
}

func TextOK(text string) Response {
	return Text(http.StatusOK, text)
}

func Text(code int, text string) Response {
	header := http.Header{}
	header.Set("content-type", "text/plain")
	return Response{
		Code:   code,
		Header: header,
		Body:   TextBody(text),
	}
}

func JSONOK(value interface{}) Response {
	return JSON(http.StatusOK, value)
}

func JSON(code int, value interface{}) Response {
	header := http.Header{}
	header.Set("content-type", "application/json")
	return Response{
		Code:   code,
		Header: header,
		Body:   JSONBody{value},
	}
}
