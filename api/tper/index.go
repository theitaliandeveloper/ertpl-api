package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type Server struct {
	Name string
	URL  string
}

var servers = []Server{
	{Name: "Serverissimo", URL: "https://tperapi.serverissimo.com"},
	{Name: "DaniLab", URL: "https://tperapi.daninet.freeddns.org"},
	{Name: "Vichingo455", URL: "https://api.vichingo455.com/tper"},
}

func checkServer(server Server) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", server.URL, nil)
	if err != nil {
		return false
	}
	req.Header.Set("Range", "bytes=0-0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	for _, server := range servers {
		if checkServer(server) {
			w.Header().Set("Cache-Control", "public, s-maxage=30")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{
				"status": "ok",
				"server": server.Name,
				"url":    server.URL,
			})
			return
		}
	}

	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusServiceUnavailable)
	json.NewEncoder(w).Encode(map[string]string{"status": "offline"})
}
