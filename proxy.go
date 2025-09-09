package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	if strings.Contains(host, ":") {
		host = strings.Split(host, ":")[0]
	}

	parts := strings.Split(host, ".")
	if len(parts) < 2 {
		http.Error(w, "Invalid host", http.StatusBadRequest)
		return
	}
	student := parts[0]

	target := "http://" + student + ":8443"
	log.Printf("Proxying %s → %s", host, target)

	backend, err := url.Parse(target)
	if err != nil {
		http.Error(w, "Invalid target", http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(backend)
	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Proxy running on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
