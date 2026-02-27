# LP Ma'arif NU Website API

Read-Only Backend API untuk website LP Ma'arif NU yang dibangun dengan Golang, Gin Framework, dan MySQL.

> 📚 **Looking for specific documentation?** See [INDEX.md](INDEX.md) for complete documentation index.

## 📋 Tech Stack

- **Golang** - Programming language
- **Gin Framework** - HTTP web framework
- **GORM** - ORM library untuk MySQL
- **MySQL** - Database
- **Viper** - Configuration management

## 🚀 Quick Start

### Prerequisites

- Go 1.21 atau lebih tinggi
- MySQL 8.0 atau lebih tinggi
- Make (optional)

### Installation

1. Clone repository ini
```bash
git clone https://github.com/lpmaarifnu/api.git
cd lpmaarifnu-site-api
```

2. Install dependencies
```bash
go mod download
# atau menggunakan make
make install
```

3. Setup configuration file
```bash
cp config.example.yaml config.yaml
# Edit config.yaml sesuai dengan environment Anda
```

4. Setup database
```bash
# Buat database
mysql -u root -p -e "CREATE DATABASE lp_maarif_nu;"

# Import schema
mysql -u root -p lp_maarif_nu < database_schema.sql

# atau menggunakan make
make setup-db
```

5. Run aplikasi
```bash
go run cmd/api/main.go
# atau menggunakan make
make run
```

API akan berjalan di `http://localhost:8080`

## 📁 Project Structure

```
lpmaarifnu-site-api/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration loader
│   ├── database/
│   │   └── database.go            # Database connection
│   ├── models/                    # Database models
│   ├── handlers/                  # HTTP handlers
│   ├── repositories/              # Data access layer
│   ├── services/                  # Business logic layer
│   ├── middleware/                # Middleware (CORS, logger, rate limit)
│   ├── utils/                     # Utility functions
│   └── routes/
│       └── routes.go              # API routes
├── config.yaml                    # Configuration file
├── config.example.yaml            # Configuration example
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 🔧 Configuration

Edit file `config.yaml` untuk konfigurasi aplikasi:

```yaml
app:
  name: "LP Maarif API"
  env: "development"
  port: 8080
  version: "1.0.0"

database:
  host: "localhost"
  port: 3306
  user: "root"
  password: "password"
  name: "lp_maarif_nu"

cors:
  allowed_origins:
    - "http://localhost:3000"
  allowed_methods:
    - "GET"
    - "OPTIONS"

rate_limit:
  enabled: true
  requests: 100
  window: 60
```

## 📚 API Documentation

### Base URL
- Development: `http://localhost:8080/api/v1`
- Production: `https://api.lpmaarifnu.or.id/v1`

### Available Endpoints

#### News
- `GET /api/v1/news` - Get all news articles
- `GET /api/v1/news/featured` - Get featured news
- `GET /api/v1/news/:slug` - Get single news article

#### Opinions
- `GET /api/v1/opinions` - Get all opinion articles
- `GET /api/v1/opinions/:slug` - Get single opinion article

#### Documents
- `GET /api/v1/documents` - Get all documents
- `GET /api/v1/documents/:id` - Get single document

#### Hero Slides
- `GET /api/v1/hero-slides` - Get active hero slides

#### Organization
- `GET /api/v1/organization/structure` - Get organization structure
- `GET /api/v1/organization/board-members` - Get board members

#### Pages
- `GET /api/v1/pages/:slug` - Get page content

#### Categories
- `GET /api/v1/categories` - Get all categories
- `GET /api/v1/categories/:slug` - Get single category

#### Settings
- `GET /api/v1/settings` - Get public settings

#### Health Check
- `GET /health` - Health check endpoint

Untuk dokumentasi lengkap, lihat file `TODO BACKEND - READ ONLY API.md`

## 🧪 Testing with Postman

### Import Collection

1. Buka Postman
2. Import `postman_collection.json` (38 requests)
3. Import `postman_environment.json` (Development environment)
4. Select environment "LP Ma'arif NU - Development"
5. Run any request!

### Quick Test

```bash
# 1. Start API
make run

# 2. In Postman, run:
#    - Health Check
#    - Get All Categories
#    - Get All News
```

Lihat [POSTMAN_GUIDE.md](POSTMAN_GUIDE.md) untuk panduan lengkap.

## 🛠️ Development

### Run in development mode
```bash
make run
```

### Build aplikasi
```bash
make build
```

### Run tests
```bash
make test
```

### Format code
```bash
make fmt
```

## 🔐 Security Features

- CORS middleware
- Rate limiting (100 requests per minute per IP)
- Input validation
- SQL injection protection (via GORM)

## 📝 License

Copyright © 2025 LP Ma'arif NU

## 📞 Support

Untuk pertanyaan atau issues:
- GitHub Issues: https://github.com/lpmaarifnu/api/issues
- Email: dev@lpmaarifnu.or.id
