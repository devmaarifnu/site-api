# Setup Guide - LP Ma'arif NU Backend API

## Panduan Lengkap Setup Project

### 1. Prerequisites

Pastikan sistem Anda sudah terinstall:
- **Go 1.21+** - [Download](https://golang.org/dl/)
- **MySQL 8.0+** - [Download](https://dev.mysql.com/downloads/)
- **Git** - [Download](https://git-scm.com/)
- **Make** (optional, untuk Windows bisa pakai chocolatey: `choco install make`)

### 2. Clone & Setup Project

```bash
# Clone repository
git clone https://github.com/lpmaarifnu/api.git
cd lpmaarifnu-site-api

# Install dependencies
go mod download
go mod tidy
```

### 3. Database Setup

#### A. Buat Database MySQL

```bash
# Login ke MySQL
mysql -u root -p

# Buat database
CREATE DATABASE lp_maarif_nu CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# Exit
exit;
```

#### B. Import Schema

```bash
# Import database schema
mysql -u root -p lp_maarif_nu < database_schema.sql
```

Database schema akan membuat:
- 14 tabel utama
- Default admin user (email: admin@lpmaarifnu.or.id, password: password)
- Default categories untuk news dan documents
- Default organization positions
- Default settings

### 4. Configuration

```bash
# Copy example config
cp config.example.yaml config.yaml

# Edit config.yaml
nano config.yaml  # atau gunakan text editor favorit Anda
```

**Minimal configuration yang harus diubah:**

```yaml
database:
  host: "localhost"
  port: 3306
  user: "root"              # Sesuaikan dengan user MySQL Anda
  password: "your_password"  # Sesuaikan dengan password MySQL Anda
  name: "lp_maarif_nu"
```

### 5. Run Application

#### Development Mode

```bash
# Menggunakan go run
go run cmd/api/main.go

# Atau menggunakan make
make run
```

#### Production Build

```bash
# Build binary
make build

# Run binary
./bin/api
```

### 6. Verify Installation

Buka browser dan akses:

```
http://localhost:8080/health
```

Response yang diharapkan:
```json
{
  "status": "OK",
  "message": "LP Maarif API is running"
}
```

Test endpoint lain:
```bash
# Get all categories
curl http://localhost:8080/api/v1/categories

# Get settings
curl http://localhost:8080/api/v1/settings
```

### 7. Troubleshooting

#### Database Connection Error

```
Error: Failed to connect to database
```

**Solution:**
- Pastikan MySQL service berjalan
- Cek credentials di config.yaml
- Pastikan database `lp_maarif_nu` sudah dibuat
- Test koneksi: `mysql -u root -p lp_maarif_nu`

#### Port Already in Use

```
Error: bind: address already in use
```

**Solution:**
- Ubah port di config.yaml (default: 8080)
- Atau kill process yang menggunakan port tersebut

#### Import Errors

```
Error: package not found
```

**Solution:**
```bash
go mod download
go mod tidy
```

### 8. Development Tools (Optional)

#### Air - Hot Reload

Install Air untuk auto-reload saat development:

```bash
# Install Air
go install github.com/air-verse/air@latest

# Run dengan Air
air
```

#### Database GUI

Untuk mempermudah management database, gunakan salah satu tools berikut:
- **MySQL Workbench** (Official)
- **phpMyAdmin** (Web-based)
- **DBeaver** (Cross-platform)
- **TablePlus** (Modern GUI)

### 9. Testing

```bash
# Run tests
make test

# Run tests with coverage
go test -v -cover ./...
```

### 10. Deployment

#### Build untuk Production

```bash
# Build
make build

# Atau build untuk Linux server
GOOS=linux GOARCH=amd64 go build -o bin/api-linux cmd/api/main.go
```

#### Environment Variables

Untuk production, set environment production di config.yaml:

```yaml
app:
  env: "production"  # Change from development
```

#### Systemd Service (Linux)

Buat file `/etc/systemd/system/lpmaarifnu-site-api.service`:

```ini
[Unit]
Description=LP Maarif API Service
After=network.target mysql.service

[Service]
Type=simple
User=www-data
WorkingDirectory=/var/www/lpmaarifnu-site-api
ExecStart=/var/www/lpmaarifnu-site-api/bin/api
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

Enable dan start service:
```bash
sudo systemctl enable lpmaarifnu-site-api
sudo systemctl start lpmaarifnu-site-api
sudo systemctl status lpmaarifnu-site-api
```

### 11. Next Steps

1. **Populate Data** - Tambahkan data berita, dokumen, dll melalui admin panel
2. **Configure CORS** - Update allowed_origins di config.yaml sesuai domain frontend
3. **Setup CDN** - Upload images/documents ke CDN dan update cdn_url di config
4. **Enable Caching** - Setup Redis untuk caching (optional tapi recommended)
5. **Monitor Logs** - Setup log monitoring untuk production

### 12. Useful Commands

```bash
# Format code
make fmt

# Run linter (jika sudah install golangci-lint)
make lint

# Clean build files
make clean

# Build for different platforms
make build-linux
make build-windows
```

## Support

Jika mengalami masalah:
1. Cek file log di terminal
2. Pastikan semua prerequisites terinstall
3. Buat issue di GitHub: https://github.com/lpmaarifnu/api/issues
4. Email: dev@lpmaarifnu.or.id
