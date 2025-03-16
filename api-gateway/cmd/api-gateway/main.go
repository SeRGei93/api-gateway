package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func createReverseProxy(target string, host string, prefixToStrip string) *httputil.ReverseProxy {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Ошибка создания прокси: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.Host = host

		req.URL.Path = strings.TrimPrefix(req.URL.Path, prefixToStrip)
		if req.URL.Path == "" {
			req.URL.Path = "/"
		}
	}

	return proxy
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Handle("/api/*", createReverseProxy("http://nginx", "api.loc", "/api"))
	r.Handle("/*", createReverseProxy("http://nginx", "frontend.loc", ""))

	addr := ":8080"
	log.Printf("🚀 API Gateway слушает на %s", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
