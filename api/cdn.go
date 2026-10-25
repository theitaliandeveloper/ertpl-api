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

type ServerCDN struct {
	Name string
	URL  string
}

var serversCDN = []ServerCDN{
	{Name: "DaniLab", URL: "https://ertpl-cdn.daninet.freeddns.org/health"},
	{Name: "Serverissimo", URL: "https://ertpl-cdn.serverissimo.com/health"},
	{Name: "Vichingo455", URL: "https://ertpl-content.vichingo455.com/health"},
}

func checkServerCDN(server ServerCDN) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2500*time.Millisecond)
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

	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

func CDNHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	for _, server := range serversCDN {
		if checkServerCDN(server) {
			w.Header().Set("Cache-Control", "public, s-maxage=30")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{
				"status": "ok",
				"server": server.Name,
				"url":    strings.Replace(server.URL, "/health", "", 1),
			})
			return
		}
	}

	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusServiceUnavailable)
	json.NewEncoder(w).Encode(map[string]string{"status": "offline"})
}
