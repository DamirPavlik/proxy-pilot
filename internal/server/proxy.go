package server

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func NewReverseProxy(target *url.URL) *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(target)
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy, targetURL *url.URL, endpointPrefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now().UTC()

		log.Printf("proxy: received request %s at %s", r.URL.String(), startTime)

		r.URL.Host = targetURL.Host
		r.URL.Scheme = targetURL.Scheme
		r.Host = targetURL.Host
		r.Header.Set("X-Forwarded-Host", r.Host)

		r.URL.Path = strings.TrimPrefix(r.URL.Path, endpointPrefix)

		log.Printf("proxy: forwarding request to %s", r.URL.String())

		proxy.ServeHTTP(w, r)

		log.Printf("proxy: completed request %s in %s", r.URL.String(), time.Since(startTime))
	}
}
