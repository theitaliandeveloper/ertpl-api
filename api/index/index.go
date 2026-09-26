package handler

import (
	"encoding/json"
	"net/http"
	"os"
)

type PackageInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	pkg := PackageInfo{
		Name:        "ertpl-api",
		Version:     "2.0.0-A1",
		Description: "API for High Availability for ER-TPL website.",
		Homepage:    "https://ertpl.pages.dev",
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, s-maxage=300")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pkg)
}
