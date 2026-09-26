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

package main

import (
	"fmt"
	"net/http"
	"os"
	"ertpl-api/api"
	"ertpl-api/utils"
)

func main() {
	mux := http.NewServeMux()

	// Vercel routes traffic according to vercel.json rewrites, 
	// which rewrites everything to /api/*
	mux.HandleFunc("/api/index", utils.WithCORS(api.IndexHandler))
	mux.HandleFunc("/api/cdn", utils.WithCORS(api.CDNHandler))
	mux.HandleFunc("/api/mezzi", utils.WithCORS(api.MezziHandler))
	mux.HandleFunc("/api/seta", utils.WithCORS(api.SetaHandler))
	mux.HandleFunc("/api/start", utils.WithCORS(api.StartHandler))
	mux.HandleFunc("/api/startsopp", utils.WithCORS(api.StartsoppHandler))
	mux.HandleFunc("/api/tper", utils.WithCORS(api.TperHandler))
	
	// Fallback for root /api
	mux.HandleFunc("/api", utils.WithCORS(api.IndexHandler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
