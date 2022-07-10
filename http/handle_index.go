package http

import "net/http"

func (s Service) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World"))
	}
}
