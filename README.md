# ERTPL API
API for the ERTPL website to check if the servers are up and running.

## Details
This API is hosted on Vercel and is called by the ER-TPL website to check if the servers are up and running.
If the first server is down, it uses the second one, if the second one is also down, it uses the third one, and so on.
If all the servers are down it returns offline to the frontend, letting the site know that all servers are down.

## License
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
