# mantis
HTTP server tool kit with composable middleware

## Example

```go
package main

import (
	"errors"
	"net/http"

	"github.com/morikuni/mantis"
)

func A(r *http.Request) (mantis.Response, error) {
	return mantis.JSONOK(map[string]interface{}{
		"name": "alice",
		"age":  18,
	}), nil
}

func B(r *http.Request) (mantis.Response, error) {
	return mantis.Nil, errors.New("error")
}

func ErrorFilter(r *http.Request, s mantis.Service) (mantis.Response, error) {
	res, err := s.Serve(r)
	if err != nil {
		return mantis.Text(http.StatusTeapot, "teapot!!!"), nil
	}
	return res, nil
}

func main() {
	m := &mantis.Builder{}
	m.UseFunc(ErrorFilter)

	http.Handle("/a", m.ServeFunc(A))
	http.Handle("/b", m.ServeFunc(B))

	http.ListenAndServe(":8080", nil)
}
```
