# API Documentation - LP Ma'arif NU

## Base Information

- **Base URL Development:** `http://localhost:8080/api/v1`
- **Base URL Production:** `https://api.lpmaarifnu.or.id/v1`
- **Response Format:** JSON
- **API Type:** Mostly Read-Only (GET), with Contact Form (POST)
- **Authentication:** None (Public API)

## Common Response Format

### Success Response
```json
{
  "success": true,
  "message": "Success message here",
  "data": { }
}
```

### Error Response
```json
{
  "success": false,
  "message": "Error message here",
  "error": "Error details"
}
```

### Paginated Response
```json
{
  "success": true,
  "message": "Success message",
  "data": {
    "items": [],
    "pagination": {
      "current_page": 1,
      "total_pages": 10,
      "total_items": 95,
      "items_per_page": 10,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

## Endpoints

### 1. News Articles

#### GET /news
Get all news articles with pagination and filters

**Query Parameters:**
- `page` (optional) - Page number (default: 1)
- `limit` (optional) - Items per page (default: 10, max: 100)
- `category` (optional) - Filter by category slug
- `search` (optional) - Search in title and excerpt
- `featured` (optional) - Filter featured articles (true/false)
- `sort` (optional) - Sort field (default: -published_at)

**Example:**
```bash
GET /api/v1/news?page=1&limit=10&category=nasional&featured=true
```

#### GET /news/featured
Get featured news articles

**Query Parameters:**
- `limit` (optional) - Number of items (default: 5)

**Example:**
```bash
GET /api/v1/news/featured?limit=5
```

#### GET /news/:slug
Get single news article by slug

**Example:**
```bash
GET /api/v1/news/peluncuran-program-beasiswa-2024
```

**Response includes:**
- Full article details
- Related articles (if category exists)
- Meta information for SEO

---

### 2. Opinion Articles

#### GET /opinions
Get all opinion articles with pagination

**Query Parameters:**
- `page` (optional) - Page number
- `limit` (optional) - Items per page
- `search` (optional) - Search in title

**Example:**
```bash
GET /api/v1/opinions?page=1&limit=10
```

#### GET /opinions/:slug
Get single opinion article by slug

**Example:**
```bash
GET /api/v1/opinions/pendidikan-karakter-era-digital
```

---

### 3. Documents

#### GET /documents
Get all documents with pagination and filters

**Query Parameters:**
- `page` (optional) - Page number
- `limit` (optional) - Items per page (default: 20)
- `category` (optional) - Filter by category slug
- `search` (optional) - Search in title and description
- `sort` (optional) - Sort field (default: -created_at)

**Example:**
```bash
GET /api/v1/documents?category=pedoman&page=1&limit=20
```

#### GET /documents/:id
Get single document by ID

**Example:**
```bash
GET /api/v1/documents/1
```

**Note:** Accessing this endpoint will increment the download counter

---

### 4. Hero Slides

#### GET /hero-slides
Get all active hero slides for homepage

**Example:**
```bash
GET /api/v1/hero-slides
```

**Response includes:**
- Title, description, image
- CTA buttons (primary and secondary)
- Ordered by order_number

---

### 5. Organization

#### GET /organization/structure
Get complete organization structure

**Example:**
```bash
GET /api/v1/organization/structure
```

**Response includes:**
- Ketua Umum
- Wakil Ketua (array)
- Sekretaris Umum
- Bendahara Umum
- Departments (array)

#### GET /organization/board-members
Get all board members with filters

**Query Parameters:**
- `period` (optional) - Filter by period (e.g., "2024-2029")
- `active` (optional) - Filter active members (true/false)

**Example:**
```bash
GET /api/v1/organization/board-members?period=2024-2029&active=true
```

#### GET /organization/pengurus
Get organization pengurus list with detailed information

**Query Parameters:**
- `periode` (optional) - Filter by periode (e.g., "2024-2029")
- `kategori` (optional) - Filter by kategori (pimpinan_utama, bidang, sekretariat, bendahara)
- `active` (optional) - Filter active members only (default: true)

**Example:**
```bash
GET /api/v1/organization/pengurus?periode=2024-2029&kategori=pimpinan_utama
```

**Response Example:**
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
        "foto": "https://ui-avatars.com/api/?name=Said+Aqil+Siradj",
        "bio": "Ulama dan intelektual muslim Indonesia",
        "email": "ketua@lpmaarifnu.or.id",
        "phone": "021-3920677",
        "periode": {
          "mulai": 2024,
          "selesai": 2029
        },
        "order": 1,
        "is_active": true
      }
    ]
  }
}
```

---

### 6. Pages

#### GET /pages/:slug
Get page content by slug

**Available slugs:**
- `visi-misi`
- `sejarah`
- `program-strategis`
- `pramuka`

**Example:**
```bash
GET /api/v1/pages/visi-misi
```

**Response includes:**
- Page title and content
- Template type
- Meta information

---

### 7. Categories

#### GET /categories
Get all categories

**Query Parameters:**
- `type` (optional) - Filter by type (news, opinion, document)

**Example:**
```bash
GET /api/v1/categories?type=news
```

#### GET /categories/:slug
Get single category with latest articles

**Example:**
```bash
GET /api/v1/categories/nasional
```

**Response includes:**
- Category details
- Latest 5 articles in this category

---

### 8. Settings

#### GET /settings
Get public settings

**Example:**
```bash
GET /api/v1/settings
```

**Response includes:**
- Site name and description
- Contact information
- Social media links
- Statistics

---

### 9. Editorial Team

#### GET /editorial/team
Get complete editorial team structure

**Example:**
```bash
GET /api/v1/editorial/team
```

**Response includes:**
- Pemimpin Redaksi
- Wakil Pemimpin Redaksi (array)
- Redaktur Pelaksana
- Tim Redaksi (array)
- Dewan Redaksi (array)
- Contact information

**Response Example:**
```json
{
  "success": true,
  "message": "Editorial team retrieved successfully",
  "data": {
    "pemimpin_redaksi": {
      "name": "Dr. H. Muhammad Fadhil, M.Pd",
      "position": "Pemimpin Redaksi",
      "photo": "https://ui-avatars.com/api/?name=Muhammad+Fadhil",
      "bio": "Pakar pendidikan Islam dengan pengalaman lebih dari 15 tahun",
      "email": "fadhil@lpmaarifnu.or.id",
      "phone": "021-12345678"
    },
    "wakil_pemimpin_redaksi": [
      {
        "name": "Dra. Hj. Nur Azizah, M.Si",
        "position": "Wakil Pemimpin Redaksi I",
        "photo": "https://ui-avatars.com/api/?name=Nur+Azizah",
        "bio": "Spesialis media dan komunikasi publik",
        "email": "azizah@lpmaarifnu.or.id"
      }
    ],
    "redaktur_pelaksana": {
      "name": "Ahmad Syarif, S.Sos, M.I.Kom",
      "position": "Redaktur Pelaksana",
      "photo": "https://ui-avatars.com/api/?name=Ahmad+Syarif",
      "bio": "Koordinator harian tim redaksi",
      "email": "syarif@lpmaarifnu.or.id"
    },
    "dewan_redaksi": [
      {
        "name": "Prof. Dr. KH. Abdullah Shiddiq, MA",
        "institution": "UIN Syarif Hidayatullah Jakarta",
        "expertise": "Pendidikan Islam & Budaya",
        "photo": "https://ui-avatars.com/api/?name=Abdullah+Shiddiq"
      }
    ],
    "tim_redaksi": [
      {
        "name": "Rizki Aulia Rahman, S.Pd",
        "position": "Editor Berita",
        "photo": "https://ui-avatars.com/api/?name=Rizki+Rahman"
      }
    ],
    "contact": {
      "email": "redaksi@lpmaarifnu.or.id",
      "phone": "021-12345678"
    }
  }
}
```

---

### 10. Contact Form

#### POST /contact/submit
Submit contact form message

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "081234567890",
  "subject": "Pertanyaan tentang pendaftaran",
  "message": "Saya ingin menanyakan prosedur pendaftaran untuk siswa baru"
}
```

**Validation Rules:**
- `name` - Required, min 3 chars, max 255 chars
- `email` - Required, valid email format
- `phone` - Optional, max 20 chars
- `subject` - Required, min 5 chars, max 500 chars
- `message` - Required, min 10 chars

**Example:**
```bash
curl -X POST http://localhost:8080/api/v1/contact/submit \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "081234567890",
    "subject": "Pertanyaan tentang pendaftaran",
    "message": "Saya ingin menanyakan prosedur pendaftaran untuk siswa baru tahun 2024/2025. Mohon informasinya."
  }'
```

**Success Response (201):**
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
    "created_at": "2025-01-14T10:30:00Z"
  }
}
```

**Error Response (400):**
```json
{
  "success": false,
  "message": "Validation failed",
  "error": "Key: 'ContactRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"
}
```

**Rate Limit:** 5 requests per hour per IP address

---

### 11. Event Flayers

#### GET /events/flayers
Get event flayers/banners for display

**Query Parameters:**
- `limit` (optional) - Number of items (default: 10, max: 100)
- `active` (optional) - Filter active events only (default: true)
- `upcoming` (optional) - Filter upcoming events only (true/false)

**Example:**
```bash
GET /api/v1/events/flayers?limit=10&active=true&upcoming=true
```

**Response Example:**
```json
{
  "success": true,
  "message": "Event flayers retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "Seminar Nasional Pendidikan Islam 2024",
      "description": "Seminar nasional dengan tema 'Transformasi Pendidikan Islam di Era Digital'",
      "image": "https://images.unsplash.com/photo-1540575467063-178a50c2df87?w=800",
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
    }
  ]
}
```

**Note:** Events are automatically filtered by display date range (between start_display_date and end_display_date)

---

### 12. Health Check

#### GET /health
Check API health status

**Example:**
```bash
GET /health
```

**Response:**
```json
{
  "status": "OK",
  "message": "LP Maarif API is running"
}
```

## Rate Limiting

### General Endpoints
- **Limit:** 100 requests per minute per IP
- **Applies to:** All GET endpoints

### Contact Form
- **Limit:** 5 requests per hour per IP
- **Applies to:** POST /contact/submit

**Response when exceeded:**
```json
{
  "success": false,
  "message": "Rate limit exceeded. Please try again later."
}
```

## CORS

API supports CORS for allowed origins configured in config.yaml

Default allowed origins:
- `http://localhost:3000`
- `http://localhost:5173`
- `https://lpmaarifnu.or.id`

## Error Codes

- **200** - Success
- **400** - Bad Request (invalid parameters)
- **404** - Not Found (resource doesn't exist)
- **429** - Too Many Requests (rate limit exceeded)
- **500** - Internal Server Error

## Example Usage

### JavaScript (Fetch)
```javascript
// Get all news
fetch('http://localhost:8080/api/v1/news?page=1&limit=10')
  .then(response => response.json())
  .then(data => console.log(data));

// Get single news
fetch('http://localhost:8080/api/v1/news/artikel-slug')
  .then(response => response.json())
  .then(data => console.log(data));

// Get pengurus
fetch('http://localhost:8080/api/v1/organization/pengurus?periode=2024-2029')
  .then(response => response.json())
  .then(data => console.log(data));

// Get editorial team
fetch('http://localhost:8080/api/v1/editorial/team')
  .then(response => response.json())
  .then(data => console.log(data));

// Submit contact form
fetch('http://localhost:8080/api/v1/contact/submit', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    name: 'John Doe',
    email: 'john@example.com',
    phone: '081234567890',
    subject: 'Pertanyaan tentang pendaftaran',
    message: 'Saya ingin menanyakan prosedur pendaftaran'
  })
})
  .then(response => response.json())
  .then(data => console.log(data));

// Get event flayers
fetch('http://localhost:8080/api/v1/events/flayers?limit=10&upcoming=true')
  .then(response => response.json())
  .then(data => console.log(data));
```

### cURL
```bash
# Get categories
curl http://localhost:8080/api/v1/categories

# Get news with filters
curl "http://localhost:8080/api/v1/news?category=nasional&page=1&limit=10"

# Get settings
curl http://localhost:8080/api/v1/settings

# Get pengurus
curl "http://localhost:8080/api/v1/organization/pengurus?periode=2024-2029&kategori=pimpinan_utama"

# Get editorial team
curl http://localhost:8080/api/v1/editorial/team

# Submit contact form
curl -X POST http://localhost:8080/api/v1/contact/submit \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "081234567890",
    "subject": "Pertanyaan tentang pendaftaran",
    "message": "Saya ingin menanyakan prosedur pendaftaran untuk siswa baru"
  }'

# Get event flayers
curl "http://localhost:8080/api/v1/events/flayers?limit=10&active=true&upcoming=true"
```

### Python (requests)
```python
import requests

# Get all news
response = requests.get('http://localhost:8080/api/v1/news', params={
    'page': 1,
    'limit': 10,
    'category': 'nasional'
})
data = response.json()
print(data)

# Get pengurus
response = requests.get('http://localhost:8080/api/v1/organization/pengurus', params={
    'periode': '2024-2029',
    'kategori': 'pimpinan_utama'
})
data = response.json()
print(data)

# Get editorial team
response = requests.get('http://localhost:8080/api/v1/editorial/team')
data = response.json()
print(data)

# Submit contact form
response = requests.post('http://localhost:8080/api/v1/contact/submit', json={
    'name': 'John Doe',
    'email': 'john@example.com',
    'phone': '081234567890',
    'subject': 'Pertanyaan tentang pendaftaran',
    'message': 'Saya ingin menanyakan prosedur pendaftaran'
})
data = response.json()
print(data)

# Get event flayers
response = requests.get('http://localhost:8080/api/v1/events/flayers', params={
    'limit': 10,
    'active': 'true',
    'upcoming': 'true'
})
data = response.json()
print(data)
```

### Next.js (React)
```javascript
// app/pengurus/page.js
export default async function PengurusPage() {
  const response = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/api/v1/organization/pengurus?periode=2024-2029`,
    { cache: 'no-store' }
  );
  const { data } = await response.json();

  return (
    <div>
      <h1>Pengurus {data.periode}</h1>
      {data.pengurus.map(p => (
        <div key={p.id}>
          <h2>{p.nama}</h2>
          <p>{p.jabatan}</p>
        </div>
      ))}
    </div>
  );
}

// components/ContactForm.jsx
'use client';
import { useState } from 'react';

export default function ContactForm() {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    phone: '',
    subject: '',
    message: ''
  });
  const [status, setStatus] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/api/v1/contact/submit`,
        {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(formData)
        }
      );

      const result = await response.json();

      if (result.success) {
        setStatus(`Pesan terkirim! Ticket ID: ${result.data.ticket_id}`);
        setFormData({ name: '', email: '', phone: '', subject: '', message: '' });
      } else {
        setStatus(`Error: ${result.message}`);
      }
    } catch (error) {
      setStatus('Gagal mengirim pesan');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={formData.name}
        onChange={(e) => setFormData({...formData, name: e.target.value})}
        placeholder="Nama"
        required
      />
      <input
        type="email"
        value={formData.email}
        onChange={(e) => setFormData({...formData, email: e.target.value})}
        placeholder="Email"
        required
      />
      <input
        type="tel"
        value={formData.phone}
        onChange={(e) => setFormData({...formData, phone: e.target.value})}
        placeholder="Telepon (opsional)"
      />
      <input
        type="text"
        value={formData.subject}
        onChange={(e) => setFormData({...formData, subject: e.target.value})}
        placeholder="Subjek"
        required
      />
      <textarea
        value={formData.message}
        onChange={(e) => setFormData({...formData, message: e.target.value})}
        placeholder="Pesan"
        required
      />
      <button type="submit">Kirim</button>
      {status && <p>{status}</p>}
    </form>
  );
}

// app/events/page.js
export default async function EventsPage() {
  const response = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/api/v1/events/flayers?limit=10&upcoming=true`,
    { next: { revalidate: 900 } } // Cache 15 minutes
  );
  const { data } = await response.json();

  return (
    <div>
      <h1>Event & Kegiatan</h1>
      {data.map(event => (
        <div key={event.id}>
          <img src={event.image} alt={event.title} />
          <h2>{event.title}</h2>
          <p>{event.description}</p>
          <p>Tanggal: {event.event_date}</p>
          <p>Lokasi: {event.event_location}</p>
          {event.registration_url && (
            <a href={event.registration_url}>Daftar</a>
          )}
        </div>
      ))}
    </div>
  );
}
```

## Best Practices

1. **Always handle errors** - Check `success` field in response
2. **Use pagination** - Don't request too many items at once
3. **Cache responses** - Use client-side caching for better performance (especially for editorial team and settings)
4. **Respect rate limits** - Implement retry logic with exponential backoff
5. **Handle 404s gracefully** - Show user-friendly error messages
6. **Contact Form** - Validate on client-side before submission to reduce failed requests
7. **Event Flayers** - Use `upcoming=true` filter for future events only
8. **Pengurus** - Cache response as it changes infrequently

## Support

For API support or questions:
- GitHub Issues: https://github.com/lpmaarifnu/api/issues
- Email: dev@lpmaarifnu.or.id
- Documentation: See `TODO BACKEND - READ ONLY API.md` for detailed specs
