package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type IPResponse struct {
	IP string `json:"ip"`
}

func getIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")

	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}

	if ip == "" {
		ip = r.RemoteAddr
	}

	return ip
}

func ipHandlerJSON(w http.ResponseWriter, r *http.Request) {
	ip := getIP(r)

	response := IPResponse{IP: ip}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ipHandlerText(w http.ResponseWriter, r *http.Request) {
	ip := getIP(r)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "%s", ip)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/ip/json", ipHandlerJSON)
	http.HandleFunc("/ip/html", ipHandlerText)
	http.HandleFunc("/ip/text", ipHandlerText)

	fmt.Printf("Server started at :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
