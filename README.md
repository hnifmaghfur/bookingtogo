# BookingToGo

BookingToGo adalah project integrasi backend Go (REST API) dan frontend Laravel (PHP) yang saling terhubung melalui API, dengan database terpisah untuk masing-masing service. Semua service dapat dijalankan sekaligus menggunakan Docker Compose.

## Struktur Direktori

```
bookingtogo/
├── backend/         # Source code backend Go
├── frontend/        # Source code frontend Laravel
├── docker-compose.yaml # Compose file untuk menjalankan semua service
└── README.md        # Dokumentasi global
```

## Arsitektur
- **Backend:** Go + PostgreSQL
- **Frontend:** Laravel (PHP) + (opsional: nginx/php-fpm) + (opsional: MySQL)
- **Komunikasi:** Frontend mengakses backend melalui HTTP API
- **Docker Compose:** Menyediakan semua service dalam satu perintah

## Menjalankan Project
1. **Clone repository**
2. **Copy dan sesuaikan env**
   - Backend: `cp backend/.env.example backend/.env`
   - Frontend: `cp frontend/.env.example frontend/.env` atau gunakan `.env.docker`
3. **Jalankan semua service:**
   ```bash
   docker compose up -d
   ```
4. **Akses aplikasi:**
   - Backend API: http://localhost:8080
   - Frontend: http://localhost:8000

## Dokumentasi Lain
- Lihat `frontend/README.md` untuk detail frontend
- Lihat `backend/README.md` untuk detail backend
