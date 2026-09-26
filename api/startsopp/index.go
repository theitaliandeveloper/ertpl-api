package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	Name string
	URL  string
}

var servers = []Server{
	{Name: "Serverissimo", URL: "https://corsesopp.startapi.serverissimo.com/versione"},
	{Name: "DaniLab", URL: "https://startsopp.daninet.freeddns.org/versione"},
	{Name: "Vichingo455", URL: "https://api.vichingo455.com/start-corsesopp.json/versione"},
}

func checkServer(server Server) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", server.URL, nil)
	if err != nil {
		return false
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return false
	}

	// Just checking if we can read/parse it
	var dummy interface{}
	if err := json.NewDecoder(resp.Body).Decode(&dummy); err != nil && err != io.EOF {
		return false
	}

	return true
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
				"url":    strings.Replace(server.URL, "/versione", "", 1),
			})
			return
		}
	}

	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusServiceUnavailable)
	json.NewEncoder(w).Encode(map[string]string{"status": "offline"})
}
