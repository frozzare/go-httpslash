package httpslash

import (
	"net/http"
	"strings"
)

// TrailingSlash will add a trailing slash to the url path and redirect
// to the new url with http status 302.
func TrailingSlash(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		len := len(r.URL.Path)
		if string(r.URL.Path[len-1]) != "/" && path.Ext(r.URL.Path) == "" {
			url := strings.Replace(r.URL.String(), r.URL.Path, r.URL.Path+"/", -1)
			http.Redirect(w, r, url, 302)
			return
		}
		h.ServeHTTP(w, r)
	})
}

// NoTrailingSlash will remove the trailing slash from the url path and redirect
// to the new url with http status 302.
func NoTrailingSlash(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		len := len(r.URL.Path)
		if len > 1 && string(r.URL.Path[len-1]) == "/" {
			url := strings.Replace(r.URL.String(), r.URL.Path, r.URL.Path[0:len-1], -1)
			http.Redirect(w, r, url, 302)
			return
		}
		h.ServeHTTP(w, r)
	})
}
