# TODO BACKEND EXTEND - READ ONLY API

## 📋 Overview
Dokumen ini berisi **Extension API Endpoints** yang perlu ditambahkan untuk melengkapi kebutuhan frontend. Semua endpoint ini adalah **Read-Only (GET only)** dan melanjutkan dari dokumentasi API utama.

**Note:**
- Data Satuan Pendidikan memiliki file dokumentasi terpisah: `TODO BACKEND - SATUAN PENDIDIKAN API.md`
- Home Features tidak diperlukan (data hardcoded di frontend)

---

## 🎯 New API Endpoints Required

### 1. MASTER DATA - PROVINSI & KABUPATEN

#### 1.1 Get All Provinsi
```http
GET /api/v1/master/provinsi
```

**Description:** Mendapatkan daftar provinsi untuk filter

**Query Parameters:**
- `with_count` (optional): Include jumlah satpen per provinsi

**Example Request:**
```bash
curl "https://api.lpmaarifnu.or.id/v1/master/provinsi?with_count=true"
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Provinsi retrieved successfully",
  "data": [
    {
      "id": 1,
      "nama": "DKI Jakarta",
      "kode": "31",
      "jumlah_satpen": 350
    },
    {
      "id": 2,
      "nama": "Jawa Barat",
      "kode": "32",
      "jumlah_satpen": 2100
    },
    {
      "id": 3,
      "nama": "Jawa Tengah",
      "kode": "33",
      "jumlah_satpen": 2500
    }
  ]
}
```

**Database Table:** `provinsi`

#### 1.2 Get Kabupaten/Kota by Provinsi
```http
GET /api/v1/master/provinsi/:id/kabupaten
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Kabupaten/kota retrieved successfully",
  "data": {
    "provinsi": {
      "id": 1,
      "nama": "DKI Jakarta",
      "kode": "31"
    },
    "kabupaten": [
      {
        "id": 1,
        "nama": "Jakarta Pusat",
        "kode": "3171",
        "jumlah_satpen": 75
      },
      {
        "id": 2,
        "nama": "Jakarta Utara",
        "kode": "3172",
        "jumlah_satpen": 68
      },
      {
        "id": 3,
        "nama": "Jakarta Barat",
        "kode": "3173",
        "jumlah_satpen": 82
      }
    ]
  }
}
```

**Database Table:** `kabupaten_kota`

---

### 2. PENGURUS ORGANISASI

#### 2.1 Get All Pengurus
```http
GET /api/v1/organization/pengurus
```

**Description:** Mendapatkan daftar pengurus organisasi LP Ma'arif NU

**Query Parameters:**
- `periode` (optional): Filter by periode (e.g., 2024-2029)
- `kategori` (optional): Filter by kategori (pimpinan_utama, bidang, sekretariat, bendahara)
- `active` (optional, default: true): Filter only active members

**Example Request:**
```bash
curl "https://api.lpmaarifnu.or.id/v1/organization/pengurus?periode=2024-2029"
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Pengurus retrieved successfully",
  "data": {
    "periode": "2024-2029",
    "pengurus": [
      {
        "id": 1,
        "nama": "Prof. Dr. KH. Said Aqil Siradj, MA",
        "jabatan": "Ketua Umum",
        "kategori": "pimpinan_utama",
        "foto": "https://cdn.lpmaarifnu.or.id/images/pengurus/ketua.jpg",
        "bio": "Ulama dan intelektual muslim Indonesia, Ketua Umum PBNU",
        "email": "ketua@lpmaarifnu.or.id",
        "phone": "021-3920677",
        "periode": {
          "mulai": 2024,
          "selesai": 2029
        },
        "order": 1,
        "is_active": true
      },
      {
        "id": 2,
        "nama": "Dr. H. Ahmad Lutfi, M.Pd",
        "jabatan": "Wakil Ketua I",
        "kategori": "pimpinan_utama",
        "foto": "https://cdn.lpmaarifnu.or.id/images/pengurus/wakil1.jpg",
        "bio": "Pakar pendidikan Islam dengan pengalaman lebih dari 20 tahun",
        "email": "wakil1@lpmaarifnu.or.id",
        "phone": "021-3920678",
        "periode": {
          "mulai": 2024,
          "selesai": 2029
        },
        "order": 2,
        "is_active": true
      }
    ]
  }
}
```

**Database Table:** `pengurus`

**Implementation Notes:**
- Grouped by `kategori` for easy rendering
- Sorted by `order` ASC
- Period validation for active members

---

### 3. EDITORIAL TEAM (REDAKSI)

#### 3.1 Get Editorial Team
```http
GET /api/v1/editorial/team
```

**Description:** Mendapatkan tim redaksi website dan publikasi

**Response Success (200):**
```json
{
  "success": true,
  "message": "Editorial team retrieved successfully",
  "data": {
    "pemimpin_redaksi": {
      "name": "Dr. H. Muhammad Fadhil, M.Pd",
      "position": "Pemimpin Redaksi",
      "photo": "https://cdn.lpmaarifnu.or.id/images/redaksi/pemred.jpg",
      "bio": "Pakar pendidikan Islam dengan pengalaman lebih dari 15 tahun di bidang jurnalistik pendidikan.",
      "email": "fadhil@lpmaarifnu.or.id",
      "phone": "021-12345678"
    },
    "wakil_pemimpin_redaksi": [
      {
        "name": "Dra. Hj. Nur Azizah, M.Si",
        "position": "Wakil Pemimpin Redaksi I",
        "photo": "https://cdn.lpmaarifnu.or.id/images/redaksi/wakil1.jpg",
        "bio": "Spesialis media dan komunikasi publik.",
        "email": "azizah@lpmaarifnu.or.id"
      },
      {
        "name": "H. Abdul Malik, S.Pd.I, M.M",
        "position": "Wakil Pemimpin Redaksi II",
        "photo": "https://cdn.lpmaarifnu.or.id/images/redaksi/wakil2.jpg",
        "bio": "Praktisi media dengan fokus pada jurnalisme pendidikan.",
        "email": "malik@lpmaarifnu.or.id"
      }
    ],
    "redaktur_pelaksana": {
      "name": "Ahmad Syarif, S.Sos, M.I.Kom",
      "position": "Redaktur Pelaksana",
      "photo": "https://cdn.lpmaarifnu.or.id/images/redaksi/redpel.jpg",
      "bio": "Koordinator harian tim redaksi dengan pengalaman di berbagai media nasional.",
      "email": "syarif@lpmaarifnu.or.id"
    },
    "dewan_redaksi": [
      {
        "name": "Prof. Dr. KH. Abdullah Shiddiq, MA",
        "institution": "UIN Syarif Hidayatullah Jakarta",
        "expertise": "Pendidikan Islam & Budaya",
        "photo": "https://cdn.lpmaarifnu.or.id/images/redaksi/dewan1.jpg"
      },
      {
        "name": "Dr. Hj. Fatimah Zahra, M.Pd",
        "institution": "Universitas Nahdlatul Ulama",
        "expertise": "Kurikulum & Pembelajaran",
        "photo": "https://cdn.lpmaarifnu.or.id/images/redaksi/dewan2.jpg"
      },
      {
        "name": "Dr. Muhammad Ridwan, M.A",
        "institution": "IAIN Surakarta",
        "expertise": "Media & Komunikasi Islam",
        "photo": "https://cdn.lpmaarifnu.or.id/images/redaksi/dewan3.jpg"
      },
      {
        "name": "Dr. Hj. Siti Nurjanah, M.Si",
        "institution": "UIN Maulana Malik Ibrahim Malang",
        "expertise": "Manajemen Pendidikan",
        "photo": "https://cdn.lpmaarifnu.or.id/images/redaksi/dewan4.jpg"
      }
    ],
    "tim_redaksi": [
      {
        "name": "Rizki Aulia Rahman, S.Pd",
        "position": "Editor Berita",
        "photo": "https://cdn.lpmaarifnu.or.id/images/redaksi/editor1.jpg"
      },
      {
        "name": "Dewi Kusuma Wardani, S.Sos",
        "position": "Editor Opini",
        "photo": "https://cdn.lpmaarifnu.or.id/images/redaksi/editor2.jpg"
      },
      {
        "name": "Faisal Akbar, S.Kom",
        "position": "Web Administrator",
        "photo": "https://cdn.lpmaarifnu.or.id/images/redaksi/webadmin.jpg"
      },
      {
        "name": "Rina Melati, S.Ds",
        "position": "Desainer Grafis",
        "photo": "https://cdn.lpmaarifnu.or.id/images/redaksi/designer.jpg"
      },
      {
        "name": "Hendra Gunawan, S.I.Kom",
        "position": "Reporter",
        "photo": "https://cdn.lpmaarifnu.or.id/images/redaksi/reporter.jpg"
      },
      {
        "name": "Laila Nur Hidayah, S.Pd",
        "position": "Content Writer",
        "photo": "https://cdn.lpmaarifnu.or.id/images/redaksi/writer.jpg"
      }
    ],
    "contact": {
      "email": "redaksi@lpmaarifnu.or.id",
      "phone": "021-12345678"
    }
  }
}
```

**Database Tables:**
- `editorial_team`
- `editorial_council`

**Implementation Notes:**
- Group by `role_type` for structured response
- Sorted by `order` ASC
- Contact info from settings or dedicated fields

---

### 4. CONTACT FORM SUBMISSION

#### 4.1 Submit Contact Message
```http
POST /api/v1/contact/submit
```

**Description:** Submit pesan dari contact form (WRITE OPERATION - Exception)

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "081234567890",
  "subject": "Pertanyaan tentang pendaftaran",
  "message": "Saya ingin menanyakan prosedur pendaftaran..."
}
```

**Validation Rules:**
- `name`: required, min 3 chars, max 255 chars
- `email`: required, valid email format
- `phone`: optional, valid phone format
- `subject`: required, min 5 chars, max 500 chars
- `message`: required, min 10 chars

**Response Success (201):**
```json
{
  "success": true,
  "message": "Pesan Anda telah terkirim. Kami akan segera menghubungi Anda.",
  "data": {
    "ticket_id": "CTK-2024-0001",
    "name": "John Doe",
    "email": "john@example.com",
    "subject": "Pertanyaan tentang pendaftaran",
    "status": "new",
    "created_at": "2024-12-15T10:30:00Z"
  }
}
```

**Response Error (400):**
```json
{
  "success": false,
  "message": "Validation failed",
  "errors": {
    "email": "Invalid email format",
    "message": "Message must be at least 10 characters"
  }
}
```

**Database Table:** `contact_messages`

**Implementation Notes:**
- Auto-generate unique `ticket_id` (CTK-YYYY-NNNN)
- Capture IP address and user agent
- Email notification to admin (optional)
- Rate limiting: max 5 messages per hour per IP

---

### 5. EVENT FLAYERS

#### 5.1 Get All Event Flayers
```http
GET /api/v1/events/flayers
```

**Description:** Mendapatkan daftar event/kegiatan yang aktif

**Query Parameters:**
- `limit` (optional, default: 10): Number of items
- `active` (optional, default: true): Filter only active events
- `upcoming` (optional): Filter upcoming events only

**Example Request:**
```bash
curl "https://api.lpmaarifnu.or.id/v1/events/flayers?limit=10&active=true"
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Event flayers retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "Seminar Nasional Pendidikan Islam 2024",
      "description": "Seminar nasional dengan tema 'Transformasi Pendidikan Islam di Era Digital'",
      "image": "https://cdn.lpmaarifnu.or.id/images/events/seminar2024.jpg",
      "event_date": "2024-03-15",
      "event_location": "Hotel Borobudur Jakarta",
      "registration_url": "https://forms.lpmaarifnu.or.id/seminar2024",
      "contact": {
        "person": "Panitia Seminar",
        "phone": "021-3920677",
        "email": "seminar@lpmaarifnu.or.id"
      },
      "display_period": {
        "start": "2024-01-01T00:00:00Z",
        "end": "2024-03-15T23:59:59Z"
      },
      "order": 1,
      "is_active": true
    },
    {
      "id": 2,
      "title": "Workshop Kurikulum Merdeka",
      "description": "Workshop implementasi Kurikulum Merdeka untuk guru-guru Ma'arif",
      "image": "https://cdn.lpmaarifnu.or.id/images/events/workshop.jpg",
      "event_date": "2024-02-20",
      "event_location": "Gedung LP Ma'arif NU Jakarta",
      "registration_url": "https://forms.lpmaarifnu.or.id/workshop-kurikulum",
      "contact": {
        "person": "Tim Kurikulum",
        "phone": "021-3920678",
        "email": "kurikulum@lpmaarifnu.or.id"
      },
      "display_period": {
        "start": "2024-01-05T00:00:00Z",
        "end": "2024-02-20T23:59:59Z"
      },
      "order": 2,
      "is_active": true
    },
    {
      "id": 3,
      "title": "Jambore Pramuka Ma'arif Nasional",
      "description": "Jambore Pramuka tingkat nasional untuk siswa-siswa Ma'arif",
      "image": "https://cdn.lpmaarifnu.or.id/images/events/jambore.jpg",
      "event_date": "2024-04-10",
      "event_location": "Cibubur, Jakarta Timur",
      "registration_url": "https://forms.lpmaarifnu.or.id/jambore2024",
      "contact": {
        "person": "Kwartir Pramuka Ma'arif",
        "phone": "021-8093789",
        "email": "pramuka@lpmaarifnu.or.id"
      },
      "display_period": {
        "start": "2024-01-10T00:00:00Z",
        "end": "2024-04-10T23:59:59Z"
      },
      "order": 3,
      "is_active": true
    }
  ]
}
```

**Database Table:** `event_flayers`

**Implementation Notes:**
- Filter by `is_active = true`
- Filter by display date range (current date between start_display_date and end_display_date)
- Sort by `order` ASC, then `event_date` DESC
- For `upcoming=true`: filter `event_date >= CURRENT_DATE`

---

## 🗂️ Database Schema Updates

### Required Tables:
1. ✅ `provinsi` - Already created in database_schema.sql
2. ✅ `kabupaten_kota` - Already created
3. ✅ `pengurus` - Already created
4. ✅ `editorial_team` - Already created
5. ✅ `editorial_council` - Already created
6. ✅ `contact_messages` - Already created
7. ✅ `event_flayers` - Already created

**All tables are ready!** See `database_schema.sql` for complete schema.

---

## 🔧 Implementation Priority

### High Priority (Must Have):
1. ✅ **GET /api/v1/master/provinsi** - Diperlukan untuk filter Data Satpen
2. ✅ **POST /api/v1/contact/submit** - Contact form functionality

### Medium Priority (Should Have):
3. ✅ **GET /api/v1/organization/pengurus** - Pengurus page
4. ✅ **GET /api/v1/editorial/team** - Redaksi page
5. ✅ **GET /api/v1/events/flayers** - Homepage events

### Low Priority (Nice to Have):
6. ✅ **GET /api/v1/master/provinsi/:id/kabupaten** - Nested filter

---

## 📊 Performance Considerations

### 1. Caching Strategy
```go
// Cache provinsi list (rarely changes)
cache.Set("master:provinsi", provinsiData, 24*time.Hour)

// Cache pengurus (changes infrequently)
cache.Set("pengurus:2024-2029", pengurusData, 12*time.Hour)

// Cache editorial team (rarely changes)
cache.Set("editorial:team", teamData, 12*time.Hour)

// Cache event flayers
cache.Set("events:flayers:active", eventsData, 15*time.Minute)
```

### 2. Database Optimization
```sql
-- Indexes already created in database_schema.sql
-- Pengurus
CREATE INDEX idx_pengurus_kategori ON pengurus(kategori);
CREATE INDEX idx_pengurus_periode ON pengurus(periode_mulai, periode_selesai);

-- Editorial Team
CREATE INDEX idx_editorial_role_type ON editorial_team(role_type);

-- Event Flayers
CREATE INDEX idx_event_dates ON event_flayers(start_display_date, end_display_date);
```

---

## 🔒 Security Considerations

### Rate Limiting
```yaml
# config.yaml
rate_limit:
  # Contact form - stricter limit
  contact_submit:
    requests: 5
    window: 3600  # 5 requests per hour

  # General endpoints
  general:
    requests: 100
    window: 60    # 100 requests per minute
```

### Input Validation
```go
// Contact form validation
type ContactRequest struct {
    Name    string `json:"name" binding:"required,min=3,max=255"`
    Email   string `json:"email" binding:"required,email"`
    Phone   string `json:"phone" binding:"omitempty,e164"`
    Subject string `json:"subject" binding:"required,min=5,max=500"`
    Message string `json:"message" binding:"required,min=10"`
}
```

### CORS Configuration
```yaml
# config.yaml
cors:
  allowed_origins:
    - "http://localhost:3000"
    - "https://lpmaarifnu.or.id"
    - "https://www.lpmaarifnu.or.id"
  allowed_methods:
    - "GET"
    - "POST"    # Only for contact form
    - "OPTIONS"
```

---

## 📝 Testing Endpoints

### Sample cURL Commands:

```bash
# 1. Get Provinsi List
curl "https://api.lpmaarifnu.or.id/v1/master/provinsi?with_count=true"

# 2. Get Kabupaten by Provinsi
curl "https://api.lpmaarifnu.or.id/v1/master/provinsi/1/kabupaten"

# 3. Get Pengurus
curl "https://api.lpmaarifnu.or.id/v1/organization/pengurus?periode=2024-2029"

# 4. Get Editorial Team
curl https://api.lpmaarifnu.or.id/v1/editorial/team

# 5. Submit Contact Form
curl -X POST https://api.lpmaarifnu.or.id/v1/contact/submit \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "phone": "081234567890",
    "subject": "Test Subject",
    "message": "This is a test message"
  }'

# 6. Get Event Flayers
curl "https://api.lpmaarifnu.or.id/v1/events/flayers?limit=10&active=true"
```

---

## 🎯 Integration with Frontend

### Example: Pengurus Page
```javascript
// src/app/tentang/susunan-pengurus/page.js
const response = await fetch(
  `${process.env.NEXT_PUBLIC_API_URL}/api/v1/organization/pengurus?periode=2024-2029`
);
const data = await response.json();
```

### Example: Contact Form
```javascript
// src/components/contact/ContactForm.jsx
const handleSubmit = async (formData) => {
  const response = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/api/v1/contact/submit`,
    {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(formData)
    }
  );
  return response.json();
};
```

---

## 📞 Support & Documentation

**Related Files:**
- Main API Docs: `TODO BACKEND - READ ONLY API.md`
- Satuan Pendidikan API: `TODO BACKEND - SATUAN PENDIDIKAN API.md`
- Database Schema: `database_schema.sql`
- Database Seeder: `database_seeder.sql`
- Frontend Requirements: `TODO EXTEND BACKEND.md`

**For Questions:**
- GitHub Issues: https://github.com/lpmaarifnu/api/issues
- Email: dev@lpmaarifnu.or.id

---

**Last Updated:** 2025-01-14
**Version:** 2.0.0
**Extension for:** LP Ma'arif NU Read-Only API
