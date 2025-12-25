# Go Fiber + PostgreSQL API

Fiber API yang membaca data `master_category_merchant`.

## Prerequisites
- Go 1.22+
- PostgreSQL reachable with an items table (id BIGSERIAL PRIMARY KEY, name TEXT)

## Setup
1. Copy .env.example to .env and adjust values.
2. Install modules:
   go mod tidy
3. Jalankan server:
   go run .

## Endpoint
- GET /master-category-merchants → mengembalikan daftar master_category_merchant.

## Notes
- Connection string dibangun dari env vars (APP_ADDR, DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSLMODE).
- Graceful shutdown menangani SIGINT/SIGTERM.

## Load Testing (k6)
- Script: [k6/master_category_merchant_test.js](k6/master_category_merchant_test.js)
- Instal k6 (Windows):
   - winget: `winget install Grafana.k6`
   - choco: `choco install k6`
   - atau unduh installer dari situs resmi k6
- Jalankan server:
   - `go run .`
- Jalankan tes:
   - `k6 run k6/master_category_merchant_test.js`
- VS Code task tersedia: "k6: master-category-merchants"
