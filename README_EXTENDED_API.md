# Extended API Features - LP Ma'arif NU

## ✅ Fitur API Tambahan

Berikut adalah API tambahan yang telah dibuat untuk melengkapi kebutuhan frontend:

### 1. **Pengurus Organisasi** ✓
**Endpoint:** `GET /api/v1/organization/pengurus`

- Filter by periode (e.g., "2024-2029")
- Filter by kategori (pimpinan_utama, bidang, sekretariat, bendahara)
- Filter by active status
- Response dengan format terstruktur per periode

**Sample Data:** 12 pengurus dari database_seeder.sql

---

### 2. **Editorial Team (Tim Redaksi)** ✓
**Endpoint:** `GET /api/v1/editorial/team`

- Grouped by role type
- Pemimpin Redaksi
- Wakil Pemimpin Redaksi
- Redaktur Pelaksana
- Tim Redaksi
- Dewan Redaksi

**Sample Data:** 14 anggota redaksi dari database_seeder.sql

---

### 3. **Contact Form Submission** ✓
**Endpoint:** `POST /api/v1/contact/submit`

- Validation lengkap (name, email, phone, subject, message)
- Auto-generate ticket ID (CTK-YYYY-NNNN)
- Capture IP address & user agent
- Rate limiting: 5 requests per hour per IP

---

### 4. **Event Flayers** ✓
**Endpoint:** `GET /api/v1/events/flayers`

- Filter by limit, active status, upcoming events
- Auto-filter by display date range
- Sorted by order & event date

**Sample Data:** 3 events dari database_seeder.sql

---

## 📁 File Structure

```
lpmaarifnu-site-api/
├── internal/
│   ├── models/
│   │   ├── pengurus.go        # Pengurus model
│   │   ├── editorial.go       # Editorial team models
│   │   ├── contact.go         # Contact message model
│   │   └── event.go           # Event flayer model
│   │
│   ├── repositories/
│   │   ├── pengurus_repository.go
│   │   ├── editorial_repository.go
│   │   ├── contact_repository.go
│   │   └── event_repository.go
│   │
│   ├── services/
│   │   ├── pengurus_service.go
│   │   ├── editorial_service.go
│   │   ├── contact_service.go
│   │   └── event_service.go
│   │
│   └── handlers/
│       ├── pengurus_handler.go
│       ├── editorial_handler.go
│       ├── contact_handler.go
│       └── event_handler.go
│
├── LP_Maarif_NU_API.postman_collection.json
└── README_EXTENDED_API.md
```

---

## 🚀 Quick Start

### 1. Build & Run
```bash
# Build
go build -o api.exe cmd/api/main.go

# Run
go run cmd/api/main.go
```

### 2. Test Endpoints

```bash
# Health check
curl http://localhost:8080/health

# Pengurus
curl "http://localhost:8080/api/v1/organization/pengurus?periode=2024-2029"

# Editorial team
curl "http://localhost:8080/api/v1/editorial/team"

# Contact form
curl -X POST "http://localhost:8080/api/v1/contact/submit" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "subject": "Test Subject",
    "message": "This is a test message"
  }'

# Events
curl "http://localhost:8080/api/v1/events/flayers?limit=10"
```

---

## 📮 Postman Collection

Import file `LP_Maarif_NU_API.postman_collection.json` ke Postman untuk testing lengkap:

**Included Endpoints:**
- ✅ Health Check
- ✅ News (Get All, Featured, By Slug)
- ✅ Opinions (Get All, By Slug)
- ✅ Documents (Get All, By ID)
- ✅ Hero Slides
- ✅ Organization (Structure, Board Members, **Pengurus**)
- ✅ Pages
- ✅ Categories
- ✅ Settings
- ✅ **Editorial Team** (NEW)
- ✅ **Contact Form** (NEW)
- ✅ **Event Flayers** (NEW)

**Steps:**
1. Open Postman
2. Click **Import**
3. Select `LP_Maarif_NU_API.postman_collection.json`
4. Set environment variable `base_url` = `http://localhost:8080`
5. Test all endpoints!

---

## 📊 Sample Data

All endpoints sudah memiliki sample data dari `database_seeder.sql`:

| Entity | Records | Description |
|--------|---------|-------------|
| Pengurus | 12 | 3 Pimpinan Utama, 1 Sekretariat, 1 Bendahara, 7 Bidang |
| Editorial Team | 10 | Tim redaksi lengkap |
| Editorial Council | 4 | Dewan redaksi dari berbagai universitas |
| Event Flayers | 3 | Seminar, Workshop, Jambore |
| Contact Messages | 4 | Sample pesan kontak |

---

## 🎯 API Response Format

### Success Response
```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... }
}
```

### Error Response
```json
{
  "success": false,
  "message": "Error message",
  "error": "Detailed error description"
}
```

---

## 🔧 Features

### 1. Clean Architecture
- Handler → Service → Repository → Database
- Consistent patterns
- Easy to extend

### 2. Validation
- Request validation using struct tags
- User-friendly error messages
- Type-safe filtering

### 3. Performance
- Efficient database queries
- Proper indexing
- Minimal N+1 queries

### 4. Security
- Rate limiting on contact form
- IP address & user agent logging
- Input sanitization

---

## 📝 Notes

### Contact Form
- **Ticket ID Format:** CTK-2024-NNNN
- **Rate Limit:** 5 requests per hour per IP
- **Validation:** All fields checked before submission

### Editorial Team
- **Grouped Response:** By role type for easy frontend rendering
- **Contact Info:** Includes general redaksi contact

### Event Flayers
- **Auto-filtering:** By display date range
- **Sorting:** By order then event date
- **Upcoming Filter:** Available for future events

---

## 🔗 Related Files

- **Postman Collection:** `LP_Maarif_NU_API.postman_collection.json`
- **Database Schema:** `database_schema.sql`
- **Database Seeder:** `database_seeder.sql`
- **Main Config:** `config/config.yaml`

---

## 🎉 Summary

**Total New Endpoints:** 4 endpoints
- `GET /api/v1/organization/pengurus`
- `GET /api/v1/editorial/team`
- `POST /api/v1/contact/submit`
- `GET /api/v1/events/flayers`

**Total Files Created:** 16 files (4 models, 4 repositories, 4 services, 4 handlers)

**Build Status:** ✅ Compiled successfully

**Ready for:** Production deployment

---

**Version:** 1.0.0
**Last Updated:** 2025-01-14
**Status:** ✅ Complete & Tested
