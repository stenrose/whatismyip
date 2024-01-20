package main

import (
	"fmt"
	"net"
	"net/http"
)

func getIP(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-Forwarded-For")

	if ip == "" {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err == nil {
			ip = host
		}
	}

	fmt.Fprintf(w, "%s\n", ip)
}

func main() {
	http.HandleFunc("/", getIP)
	http.ListenAndServe(":8080", nil)
}
