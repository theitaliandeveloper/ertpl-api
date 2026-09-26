/*
ER-TPL High Availability API
Copyright (C) 2025-26 Vichingo455

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type ServerMezzi struct {
	Name        string
	URL         string
	APIEndpoint string
}

var serversMezzi = []ServerMezzi{
	{Name: "Serverissimo", URL: "https://mezziapi.serverissimo.com/health", APIEndpoint: "/api"},
	{Name: "DaniLab", URL: "https://ertpl-db.daninet.freeddns.org/health", APIEndpoint: "/api"},
	{Name: "Vichingo455", URL: "https://api.vichingo455.com/ertpl/health", APIEndpoint: ""},
}

func checkServerMezzi(server ServerMezzi) bool {
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

func MezziHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	for _, server := range serversMezzi {
		if checkServerMezzi(server) {
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
