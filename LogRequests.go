package commontools

import (
	"fmt"
	"net/http"
)

func LogRequests(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header.Get("x-forwarded-for"), r.Method, r.RequestURI,
			"\n ⤷", Cookie2Str(r))
		handler(w, r)
	}
}
