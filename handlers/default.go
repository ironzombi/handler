package handlers

import (
	"io"
	"io/ioutil"
	"net/http"
	"text/template"
)

var t = template.Must(template.New("hello").Parse("Hello!!! {{.}}!"))

func DefaultHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func(r io.ReadCloser) {
				_, _ = io.Copy(ioutil.Discard, r)
				_ = r.Close()
			}(r.Body)

			var b []byte

			switch r.Method {
			case http.MethodPut:
				b = []byte("Put")
			case http.MethodGet:
				b = []byte("Friend")
			case http.MethodPost:
				var err error
				b, err = ioutil.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}

			case http.MethodDelete:
				b = []byte("Delete")
			default:
				// rfc requires allow header
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			_ = t.Execute(w, string(b))
		},
	)
}
