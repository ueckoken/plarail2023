package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	frontendURL, _ := url.Parse("http://localhost:5173")
	backendURL, _ := url.Parse("http://localhost:8080")

	mux := http.NewServeMux()
	mux.Handle("/", httputil.NewSingleHostReverseProxy(frontendURL))
	rev := httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Path = strings.Replace(r.URL.Path, "/api", "", 1)
			r.URL.Host = backendURL.Host
			r.URL.Scheme = backendURL.Scheme
			r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		},
	}
	mux.Handle("/api/", &rev)

	fmt.Println(http.ListenAndServe(":3031", mux))
}
