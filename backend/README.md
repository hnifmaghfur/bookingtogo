# Backend - BookingToGo

Backend BookingToGo adalah REST API berbasis Go yang telah menerapkan `clean architecture`, dan melayani data customer dan family list, terhubung ke database PostgreSQL.

## Fitur

- CRUD Customer (dengan primary key `cst_id`)
- Bulk insert Family List (semua anggota keluarga dengan `cst_id` yang sama)
- Validasi dan logging

## Struktur Penting

- `internal/domain/` : Model/domain entity
- `internal/repository/`: Query database (GORM)
- `internal/service/` : Bisnis logic
- `routes/` : HTTP handler & routing
- `main.go` : Entry point

## Setup Development

1. **Copy env:**
   ```bash
   cp .env.example .env
   ```
2. **Jalankan PostgreSQL** (atau gunakan docker-compose)
3. **Jalankan backend:**
   ```bash
   go run main.go
   ```

## Jalankan via Docker (Recomended)

Jalankan dari root backend project:

```bash
   docker compose up -d
```

## Endpoint Penting

- `GET /customers` : Mengambil data customer
- `GET /family-lists` : Mengambil data family list
- Endpoint lain: lihat `routes/routes.go`

## Dokumentasi Global

Lihat file [../README.md](../README.md) untuk arsitektur dan instruksi global.
