package mantis

import (
	"encoding/json"
	"io"
)

type Body interface {
	WriteTo(w io.Writer) error
}

type TextBody string

func (t TextBody) WriteTo(w io.Writer) error {
	_, err := io.WriteString(w, string(t))
	return err
}

type JSONBody struct {
	Value interface{}
}

func (j JSONBody) WriteTo(w io.Writer) error {
	return json.NewEncoder(w).Encode(j.Value)
}
