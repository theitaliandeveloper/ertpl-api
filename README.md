# ERTPL API
API for the ERTPL website to check if the servers are up and running.

## Details
This API is hosted on Vercel and is called by the ER-TPL website to check if the servers are up and running.
If the first server is down, it uses the second one, if the second one is also down, it uses the third one, and so on.
If all the servers are down it returns offline to the frontend, letting the site know that all servers are down.

##### Copyright (C) 2025 EmmeV Code
