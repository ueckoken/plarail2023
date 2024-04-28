package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func main() {
	frontendURL, _ := url.Parse(os.Getenv("FRONTEND_URL"))
	backendURL, _ := url.Parse(os.Getenv("BACKEND_URL"))

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
