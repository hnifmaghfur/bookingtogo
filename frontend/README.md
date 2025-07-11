# Frontend BookingToGo

Frontend BookingToGo adalah aplikasi Laravel yang berfungsi sebagai antarmuka pengguna untuk manajemen customer dan family list, terintegrasi dengan backend Go melalui REST API.

## Fitur Utama
- **Form Customer:** Input data customer baru.
- **Bulk Family List:** Tambah banyak anggota keluarga sekaligus untuk satu customer.
- **Integrasi API:** Semua data dikirim/diambil dari backend Go.
- **Validasi & Logging:** Validasi form dan logging error respons API.

## Struktur Penting
- `app/Http/Controllers/CustomerController.php` — logika utama komunikasi ke backend Go
- `resources/views/` — halaman utama dan form
- `.env` / `.env.docker` — konfigurasi environment

## Setup Development
1. **Install dependency:**
   ```bash
   composer install
   npm install
   ```
2. **Copy dan edit env:**
   ```bash
   cp .env.example .env
   # atau gunakan .env.docker untuk docker
   ```
3. **Generate key:**
   ```bash
   php artisan key:generate
   ```
4. **Jalankan secara lokal:**
   ```bash
   php artisan serve
   ```

## Jalankan via Docker
1. Jalankan dari root project:
   ```bash
   docker compose up -d
   ```
2. Akses di http://localhost:8000

## Koneksi ke Backend API
- Pastikan variabel `API_BASE_URL` di `.env` mengarah ke backend (misal: `http://backend-app:8080` dalam docker, atau `http://localhost:8080` secara lokal).
- Semua request customer & family akan diteruskan ke backend Go.

## Dokumentasi Lain
Lihat file [../README.md](../README.md) untuk dokumentasi global dan [../backend/README.md](../backend/README.md) untuk detail backend.
