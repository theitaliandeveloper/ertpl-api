package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	Name        string
	URL         string
	APIEndpoint string
}

var servers = []Server{
	{Name: "Serverissimo", URL: "https://mezziapi.serverissimo.com/health", APIEndpoint: "/api"},
	{Name: "DaniLab", URL: "https://ertpl-db.daninet.freeddns.org/health", APIEndpoint: "/api"},
	{Name: "Vichingo455", URL: "https://api.vichingo455.com/ertpl/health", APIEndpoint: ""},
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

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false
	}

	var data struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return false
	}

	return data.Status == "ok"
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
				"url":    strings.Replace(server.URL, "/health", server.APIEndpoint, 1),
			})
			return
		}
	}

	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusServiceUnavailable)
	json.NewEncoder(w).Encode(map[string]string{"status": "offline"})
}
