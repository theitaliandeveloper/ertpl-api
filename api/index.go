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
	"encoding/json"
	"net/http"
)

type PackageInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	pkg := PackageInfo{
		Name:        "ertpl-api",
		Version:     "2.0.0-B1",
		Description: "API for High Availability for ER-TPL website.",
		Homepage:    "https://ertpl.pages.dev",
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, s-maxage=300")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pkg)
}
