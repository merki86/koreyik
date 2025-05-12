# /config директориясы

Осы директория жобаның конфигурация (пішіндеме) файлдары сақтайды.

# /config directory

This directory contains configuration files of the project.

# /config direktoriyası

Osı direktoriya jobanıñ konfigurasıya (picindeme) fayıldarı saqtaydı.

# Usage
Create `local.yaml` or `prod.yaml` in this directory and fill the fields below:
```yaml
env: "local" # local/prod
version: "v1.0.1"

server:
  address: "localhost:8080"
  timeout: 4s
  idle_timeout: 30s
  shutdown_timeout: 10s

storage:
  server: "localhost"
  database: "koreyik"
  port: <port here>
  username: "<username here>"
  password: STORAGE_PASSWORD
```
Uppercased values are environment variables. Ensure you have set up them in `.env` file. Use `.env.example` as sample form.