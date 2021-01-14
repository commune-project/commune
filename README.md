# commune - ActivityPub backend server of Commune Project.

Commune is an ActivityPub-enabled forum software in its very very very early stage of development.

DO NOT use it in a production environment, NO function is done.

Dependencies:
* Go (tested with go1.15.6 linux/amd64)
* PostgreSQL (tested with 11.10-1.pgdg100+1)

# Development

You must prepare a PostgreSQL database for it.

Copy .env.sample to .env, and configure DATABASE_URL, COMMUNE_LOCAL_DOMAINS and so on.

```bash
~/go/bin/godotenv go run cli/communectl/communectl.go migrate

bash main.sh
```

And you can access the unfinished page on http://localhost:8000. You may need to set a proper reverse proxy before it…

# Run Unit Tests
#TODO
```bash
bash test_all.sh
```

License

Copyright (C) 2021 Misaka 0x4e21

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see https://www.gnu.org/licenses/.