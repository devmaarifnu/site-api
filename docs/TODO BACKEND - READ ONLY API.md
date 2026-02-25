# TODO BACKEND - LP Ma'arif NU Website API (Read-Only)

## 📋 Overview
Dokumen ini berisi requirement lengkap untuk **Read-Only Backend API** yang dikhususkan untuk **passing data dari database ke website**. API ini hanya menyediakan endpoint GET untuk membaca data, tanpa operasi Create, Update, atau Delete.

---

## 🎯 Tech Stack

### Backend Framework
- **Golang** - Programming language
- **Gin Framework** - HTTP web framework
- **GORM** - ORM library untuk MySQL
- **MySQL** - Database

### Libraries & Tools
```go
github.com/gin-gonic/gin              // Web framework
github.com/gin-contrib/cors           // CORS middleware
gorm.io/gorm                          // ORM
gorm.io/driver/mysql                  // MySQL driver
github.com/spf13/viper                // Configuration management (YAML)
github.com/go-playground/validator/v10 // Validation
```

---

## 🎯 Base URL
```
Production: https://api.lpmaarifnu.or.id/v1
Development: http://localhost:8080/api/v1
```

---

## 📌 API Characteristics

### Read-Only Features:
- ✅ **GET only** - Semua endpoint hanya GET
- ✅ **No Authentication** - Public API untuk website
- ✅ **CORS Enabled** - Allow frontend access
- ✅ **Caching** - Redis/In-memory cache untuk performance
- ✅ **Rate Limiting** - 100 requests per minute per IP
- ✅ **Response Compression** - Gzip compression

### No Write Operations:
- ❌ No POST, PUT, DELETE, PATCH
- ❌ No Authentication required
- ❌ No Admin panel (separate admin app will handle writes)

---

## 🗂️ API Endpoints (Read-Only)

### 1. BERITA (News Articles)

#### 1.1 Get All News Articles
```http
GET /api/v1/news
```

**Query Parameters:**
- `category` (optional): Filter by category slug (nasional, daerah, program)
- `page` (optional, default: 1): Page number
- `limit` (optional, default: 10): Items per page
- `search` (optional): Search by title or content
- `sort` (optional, default: -published_at): Sort field
- `tags` (optional): Filter by tags (comma separated)
- `featured` (optional): Filter featured articles (true/false)

**Example Request:**
```bash
curl https://api.lpmaarifnu.or.id/v1/news?category=nasional&page=1&limit=10
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "News articles retrieved successfully",
  "data": {
    "articles": [
      {
        "id": 1,
        "title": "Peluncuran Program Beasiswa 2024",
        "slug": "peluncuran-program-beasiswa-2024",
        "excerpt": "LP Ma'arif NU meluncurkan program beasiswa...",
        "content": "<p>Full HTML content...</p>",
        "image": "https://cdn.lpmaarifnu.or.id/images/news/image.jpg",
        "published_at": "2024-12-15T10:30:00Z",
        "category": {
          "id": 1,
          "name": "Nasional",
          "slug": "nasional"
        },
        "author": {
          "id": 1,
          "name": "Admin LP Ma'arif"
        },
        "tags": [
          {
            "id": 1,
            "name": "beasiswa",
            "slug": "beasiswa"
          }
        ],
        "views": 1250,
        "is_featured": true,
        "created_at": "2024-12-15T08:00:00Z",
        "updated_at": "2024-12-15T10:30:00Z"
      }
    ],
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

**Response Error (400):**
```json
{
  "success": false,
  "message": "Invalid query parameters",
  "errors": {
    "page": "must be a positive integer"
  }
}
```

**Response Error (500):**
```json
{
  "success": false,
  "message": "Internal server error",
  "error": "Database connection failed"
}
```

#### 1.2 Get Single News Article by Slug
```http
GET /api/v1/news/:slug
```

**Path Parameters:**
- `slug` (required): Article slug

**Example Request:**
```bash
curl https://api.lpmaarifnu.or.id/v1/news/peluncuran-program-beasiswa-2024
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "News article retrieved successfully",
  "data": {
    "id": 1,
    "title": "Peluncuran Program Beasiswa 2024",
    "slug": "peluncuran-program-beasiswa-2024",
    "excerpt": "LP Ma'arif NU meluncurkan...",
    "content": "<p>Full HTML content...</p>",
    "image": "https://cdn.lpmaarifnu.or.id/images/news/image.jpg",
    "published_at": "2024-12-15T10:30:00Z",
    "category": {
      "id": 1,
      "name": "Nasional",
      "slug": "nasional"
    },
    "author": {
      "id": 1,
      "name": "Admin LP Ma'arif",
      "avatar": "https://cdn.lpmaarifnu.or.id/images/avatars/admin.jpg"
    },
    "tags": [
      {
        "id": 1,
        "name": "beasiswa",
        "slug": "beasiswa"
      }
    ],
    "views": 1250,
    "is_featured": true,
    "meta": {
      "title": "Peluncuran Program Beasiswa 2024",
      "description": "LP Ma'arif NU meluncurkan...",
      "keywords": "beasiswa, pendidikan, ma'arif"
    },
    "related_articles": [
      {
        "id": 2,
        "title": "Program Beasiswa Tahun Lalu",
        "slug": "program-beasiswa-tahun-lalu",
        "image": "https://cdn.lpmaarifnu.or.id/images/news/image2.jpg",
        "published_at": "2023-12-10T09:00:00Z"
      }
    ],
    "created_at": "2024-12-15T08:00:00Z",
    "updated_at": "2024-12-15T10:30:00Z"
  }
}
```

**Response Error (404):**
```json
{
  "success": false,
  "message": "News article not found",
  "error": "Article with slug 'invalid-slug' does not exist"
}
```

#### 1.3 Get Featured News
```http
GET /api/v1/news/featured
```

**Query Parameters:**
- `limit` (optional, default: 5): Number of featured articles

**Response Success (200):**
```json
{
  "success": true,
  "message": "Featured news retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "Peluncuran Program Beasiswa 2024",
      "slug": "peluncuran-program-beasiswa-2024",
      "excerpt": "LP Ma'arif NU meluncurkan...",
      "image": "https://cdn.lpmaarifnu.or.id/images/news/image.jpg",
      "published_at": "2024-12-15T10:30:00Z",
      "category": {
        "id": 1,
        "name": "Nasional",
        "slug": "nasional"
      },
      "views": 1250
    }
  ]
}
```

---

### 2. OPINI (Opinion Articles)

#### 2.1 Get All Opinion Articles
```http
GET /api/v1/opinions
```

**Query Parameters:**
- `page` (optional, default: 1): Page number
- `limit` (optional, default: 10): Items per page
- `search` (optional): Search by title
- `tags` (optional): Filter by tags

**Response Success (200):**
```json
{
  "success": true,
  "message": "Opinion articles retrieved successfully",
  "data": {
    "articles": [
      {
        "id": 1,
        "title": "Pendidikan Karakter di Era Digital",
        "slug": "pendidikan-karakter-era-digital",
        "excerpt": "Pentingnya menanamkan nilai-nilai...",
        "content": "<p>Full HTML content...</p>",
        "image": "https://cdn.lpmaarifnu.or.id/images/opinions/image.jpg",
        "published_at": "2024-12-12T14:00:00Z",
        "author": {
          "name": "Prof. Dr. Ahmad Syafii",
          "title": "Pakar Pendidikan Islam",
          "image": "https://cdn.lpmaarifnu.or.id/images/authors/author1.jpg",
          "bio": "Pakar pendidikan Islam dengan pengalaman..."
        },
        "tags": [
          {
            "id": 5,
            "name": "pendidikan",
            "slug": "pendidikan"
          }
        ],
        "views": 890,
        "created_at": "2024-12-12T12:00:00Z",
        "updated_at": "2024-12-12T14:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 5,
      "total_items": 48,
      "items_per_page": 10,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

#### 2.2 Get Single Opinion Article by Slug
```http
GET /api/v1/opinions/:slug
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Opinion article retrieved successfully",
  "data": {
    "id": 1,
    "title": "Pendidikan Karakter di Era Digital",
    "slug": "pendidikan-karakter-era-digital",
    "excerpt": "Pentingnya menanamkan nilai-nilai...",
    "content": "<p>Full HTML content...</p>",
    "image": "https://cdn.lpmaarifnu.or.id/images/opinions/image.jpg",
    "published_at": "2024-12-12T14:00:00Z",
    "author": {
      "name": "Prof. Dr. Ahmad Syafii",
      "title": "Pakar Pendidikan Islam",
      "image": "https://cdn.lpmaarifnu.or.id/images/authors/author1.jpg",
      "bio": "Pakar pendidikan Islam..."
    },
    "tags": [
      {
        "id": 5,
        "name": "pendidikan",
        "slug": "pendidikan"
      }
    ],
    "views": 890,
    "meta": {
      "title": "Pendidikan Karakter di Era Digital",
      "description": "Pentingnya menanamkan...",
      "keywords": "pendidikan, karakter, digital"
    },
    "created_at": "2024-12-12T12:00:00Z",
    "updated_at": "2024-12-12T14:00:00Z"
  }
}
```

---

### 3. DOKUMEN (Documents)

#### 3.1 Get All Documents
```http
GET /api/v1/documents
```

**Query Parameters:**
- `category` (optional): Filter by category (Pedoman, Kurikulum, Regulasi, Panduan, Formulir)
- `search` (optional): Search by title
- `page` (optional, default: 1): Page number
- `limit` (optional, default: 20): Items per page
- `sort` (optional, default: -created_at): Sort field

**Response Success (200):**
```json
{
  "success": true,
  "message": "Documents retrieved successfully",
  "data": {
    "documents": [
      {
        "id": 1,
        "title": "Pedoman Penyelenggaraan Pendidikan Ma'arif",
        "description": "Panduan lengkap penyelenggaraan...",
        "category": {
          "id": 5,
          "name": "Pedoman",
          "slug": "pedoman"
        },
        "file_name": "pedoman-pendidikan-maarif.pdf",
        "file_type": "PDF",
        "file_size": 2621440,
        "file_size_formatted": "2.5 MB",
        "download_url": "https://cdn.lpmaarifnu.or.id/documents/pedoman-pendidikan-maarif.pdf",
        "download_count": 1534,
        "uploaded_at": "2024-11-15T10:00:00Z",
        "created_at": "2024-11-15T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 8,
      "total_items": 156,
      "items_per_page": 20,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

#### 3.2 Get Single Document
```http
GET /api/v1/documents/:id
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Document retrieved successfully",
  "data": {
    "id": 1,
    "title": "Pedoman Penyelenggaraan Pendidikan Ma'arif",
    "description": "Panduan lengkap penyelenggaraan pendidikan...",
    "category": {
      "id": 5,
      "name": "Pedoman",
      "slug": "pedoman"
    },
    "file_name": "pedoman-pendidikan-maarif.pdf",
    "file_type": "PDF",
    "file_size": 2621440,
    "file_size_formatted": "2.5 MB",
    "mime_type": "application/pdf",
    "download_url": "https://cdn.lpmaarifnu.or.id/documents/pedoman-pendidikan-maarif.pdf",
    "download_count": 1534,
    "uploaded_by": {
      "id": 1,
      "name": "Admin LP Ma'arif"
    },
    "uploaded_at": "2024-11-15T10:00:00Z",
    "created_at": "2024-11-15T10:00:00Z",
    "updated_at": "2024-11-15T10:00:00Z"
  }
}
```

---

### 4. HERO SLIDER (Home Page)

#### 4.1 Get All Active Hero Slides
```http
GET /api/v1/hero-slides
```

**Query Parameters:**
- `active` (optional, default: true): Get only active slides

**Response Success (200):**
```json
{
  "success": true,
  "message": "Hero slides retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "Membangun Pendidikan Islam Berkualitas",
      "description": "LP Ma'arif NU berkomitmen mengembangkan...",
      "image": "https://cdn.lpmaarifnu.or.id/images/hero/slide1.jpg",
      "cta": {
        "label": "Pelajari Lebih Lanjut",
        "href": "/tentang/visi-misi",
        "secondary": {
          "label": "Hubungi Kami",
          "href": "/kontak"
        }
      },
      "order": 1,
      "is_active": true,
      "created_at": "2024-11-01T08:00:00Z"
    }
  ]
}
```

---

### 5. ORGANISASI (Organization)

#### 5.1 Get Organization Structure
```http
GET /api/v1/organization/structure
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Organization structure retrieved successfully",
  "data": {
    "ketua": {
      "id": 1,
      "name": "Prof. Dr. KH. Said Aqil Siradj, MA",
      "position": "Ketua Umum",
      "photo": "https://cdn.lpmaarifnu.or.id/images/org/ketua.jpg",
      "bio": "Ketua Umum LP Ma'arif NU...",
      "period": "2024-2029"
    },
    "wakil": [
      {
        "id": 2,
        "name": "Dr. H. Ahmad Lutfi, M.Pd",
        "position": "Wakil Ketua I",
        "photo": "https://cdn.lpmaarifnu.or.id/images/org/wakil1.jpg",
        "bio": "Wakil Ketua I LP Ma'arif NU...",
        "period": "2024-2029"
      }
    ],
    "sekretaris": {
      "id": 4,
      "name": "Dr. Hj. Siti Aisyah, M.Pd",
      "position": "Sekretaris Umum",
      "photo": "https://cdn.lpmaarifnu.or.id/images/org/sekretaris.jpg",
      "bio": "Sekretaris Umum LP Ma'arif NU...",
      "period": "2024-2029"
    },
    "bendahara": {
      "id": 5,
      "name": "H. Abdul Rahman, SE, M.Ak",
      "position": "Bendahara Umum",
      "photo": "https://cdn.lpmaarifnu.or.id/images/org/bendahara.jpg",
      "bio": "Bendahara Umum LP Ma'arif NU...",
      "period": "2024-2029"
    },
    "departments": [
      {
        "id": 1,
        "name": "Bidang Pendidikan Dasar",
        "head": "Dr. Muhammad Yusuf, M.Pd",
        "description": "Mengelola pengembangan pendidikan tingkat SD/MI"
      }
    ]
  }
}
```

#### 5.2 Get All Board Members
```http
GET /api/v1/organization/board-members
```

**Query Parameters:**
- `period` (optional): Filter by period (e.g., 2024-2029)
- `active` (optional, default: true): Get only active members

**Response Success (200):**
```json
{
  "success": true,
  "message": "Board members retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Prof. Dr. KH. Said Aqil Siradj, MA",
      "position": "Ketua Umum",
      "photo": "https://cdn.lpmaarifnu.or.id/images/org/member1.jpg",
      "bio": "Ketua Umum LP Ma'arif NU periode 2024-2029",
      "email": "ketua@lpmaarifnu.or.id",
      "phone": "021-1234567",
      "period": "2024-2029",
      "order": 1
    }
  ]
}
```

---

### 6. CONTENT PAGES (Static Content)

#### 6.1 Get Page Content by Slug
```http
GET /api/v1/pages/:slug
```

**Available Slugs:**
- `visi-misi`
- `sejarah`
- `program-strategis`
- `pramuka`

**Example Request:**
```bash
curl https://api.lpmaarifnu.or.id/v1/pages/visi-misi
```

**Response Success (200) - Visi Misi:**
```json
{
  "success": true,
  "message": "Page content retrieved successfully",
  "data": {
    "slug": "visi-misi",
    "title": "Visi & Misi",
    "template": "visi-misi",
    "content": {
      "visi": "Menjadi lembaga pendidikan Islam yang unggul...",
      "misi": [
        "Menyelenggarakan pendidikan berkualitas...",
        "Mengembangkan kurikulum yang integratif...",
        "Memberdayakan sumber daya manusia..."
      ],
      "nilai_nilai": [
        {
          "icon": "integrity",
          "title": "Integritas",
          "description": "Menjunjung tinggi kejujuran..."
        }
      ]
    },
    "meta": {
      "title": "Visi & Misi LP Ma'arif NU",
      "description": "Visi dan Misi LP Ma'arif NU...",
      "keywords": "visi, misi, ma'arif, nu"
    },
    "is_active": true,
    "updated_at": "2024-12-01T10:00:00Z"
  }
}
```

**Response Success (200) - Sejarah:**
```json
{
  "success": true,
  "message": "Page content retrieved successfully",
  "data": {
    "slug": "sejarah",
    "title": "Sejarah LP Ma'arif NU",
    "template": "sejarah",
    "content": {
      "introduction": "<p>LP Ma'arif NU merupakan...</p>",
      "timeline": [
        {
          "year": "1916",
          "title": "Berdirinya Nahdlatul Ulama",
          "description": "Nahdlatul Ulama (NU) didirikan..."
        }
      ]
    },
    "meta": {
      "title": "Sejarah LP Ma'arif NU",
      "description": "Perjalanan panjang LP Ma'arif NU...",
      "keywords": "sejarah, ma'arif, nu"
    },
    "is_active": true,
    "updated_at": "2024-12-01T10:00:00Z"
  }
}
```

---

### 7. CATEGORIES & TAGS

#### 7.1 Get All Categories
```http
GET /api/v1/categories
```

**Query Parameters:**
- `type` (optional): Filter by type (news, opinion, document)

**Response Success (200):**
```json
{
  "success": true,
  "message": "Categories retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Nasional",
      "slug": "nasional",
      "description": "Berita tingkat nasional",
      "type": "news",
      "icon": "flag",
      "color": "#0066CC",
      "article_count": 45
    }
  ]
}
```

#### 7.2 Get Single Category
```http
GET /api/v1/categories/:slug
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Category retrieved successfully",
  "data": {
    "id": 1,
    "name": "Nasional",
    "slug": "nasional",
    "description": "Berita tingkat nasional",
    "type": "news",
    "icon": "flag",
    "color": "#0066CC",
    "article_count": 45,
    "latest_articles": [
      {
        "id": 1,
        "title": "Peluncuran Program Beasiswa 2024",
        "slug": "peluncuran-program-beasiswa-2024",
        "image": "https://cdn.lpmaarifnu.or.id/images/news/image.jpg",
        "published_at": "2024-12-15T10:30:00Z"
      }
    ]
  }
}
```

#### 7.3 Get All Tags
```http
GET /api/v1/tags
```

**Query Parameters:**
- `limit` (optional): Limit number of tags
- `popular` (optional): Get most popular tags

**Response Success (200):**
```json
{
  "success": true,
  "message": "Tags retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "beasiswa",
      "slug": "beasiswa",
      "article_count": 23
    }
  ]
}
```

---

### 8. SETTINGS

#### 8.1 Get Public Settings
```http
GET /api/v1/settings
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Settings retrieved successfully",
  "data": {
    "site_name": "LP Ma'arif NU",
    "site_description": "Lembaga Pendidikan Ma'arif Nahdlatul Ulama",
    "logo": "https://cdn.lpmaarifnu.or.id/images/logo.png",
    "favicon": "https://cdn.lpmaarifnu.or.id/images/favicon.ico",
    "contact": {
      "email": "info@lpmaarifnu.or.id",
      "phone": "021-1234567",
      "address": "Jakarta, Indonesia"
    },
    "social_media": {
      "facebook": "https://facebook.com/lpmaarifnu",
      "twitter": "https://twitter.com/lpmaarifnu",
      "instagram": "https://instagram.com/lpmaarifnu",
      "youtube": "https://youtube.com/@lpmaarifnu"
    },
    "statistics": {
      "total_articles": 95,
      "total_documents": 156
    }
  }
}
```

---

### 9. SEARCH (Global Search)

#### 9.1 Global Search
```http
GET /api/v1/search
```

**Query Parameters:**
- `q` (required): Search query
- `type` (optional): Filter by type (news, opinion, document, all)
- `page` (optional, default: 1): Page number
- `limit` (optional, default: 20): Items per page

**Example Request:**
```bash
curl "https://api.lpmaarifnu.or.id/v1/search?q=pendidikan&type=all&page=1&limit=20"
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Search results retrieved successfully",
  "data": {
    "query": "pendidikan",
    "total_results": 87,
    "results": [
      {
        "type": "news",
        "id": 1,
        "title": "Program Pendidikan Terpadu",
        "slug": "program-pendidikan-terpadu",
        "excerpt": "Mengintegrasikan nilai-nilai keislaman...",
        "image": "https://cdn.lpmaarifnu.or.id/images/news/image.jpg",
        "published_at": "2024-12-10T10:00:00Z",
        "relevance_score": 0.95
      },
      {
        "type": "opinion",
        "id": 5,
        "title": "Pendidikan Karakter di Era Digital",
        "slug": "pendidikan-karakter-era-digital",
        "excerpt": "Pentingnya menanamkan nilai-nilai...",
        "image": "https://cdn.lpmaarifnu.or.id/images/opinions/image.jpg",
        "published_at": "2024-12-12T14:00:00Z",
        "relevance_score": 0.89
      },
      {
        "type": "document",
        "id": 3,
        "title": "Pedoman Pendidikan Ma'arif",
        "file_type": "PDF",
        "file_size_formatted": "3.2 MB",
        "relevance_score": 0.78
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 5,
      "total_items": 87,
      "items_per_page": 20,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

---

### 10. ANALYTICS (Public Stats)

#### 10.1 Get Public Statistics
```http
GET /api/v1/analytics/stats
```

**Response Success (200):**
```json
{
  "success": true,
  "message": "Statistics retrieved successfully",
  "data": {
    "total_articles": 95,
    "total_opinions": 48,
    "total_documents": 156,
    "total_views": 125430,
    "total_downloads": 8756,
    "popular_articles": [
      {
        "id": 1,
        "title": "Peluncuran Program Beasiswa 2024",
        "slug": "peluncuran-program-beasiswa-2024",
        "views": 1250,
        "image": "https://cdn.lpmaarifnu.or.id/images/news/image.jpg"
      }
    ],
    "popular_documents": [
      {
        "id": 1,
        "title": "Pedoman Penyelenggaraan Pendidikan",
        "downloads": 1534,
        "file_type": "PDF"
      }
    ],
    "latest_update": "2024-12-15T10:30:00Z"
  }
}
```

---

## 🗂️ Golang Project Structure

```
lpmaarifnu-site-api/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration (YAML)
│   ├── database/
│   │   └── database.go            # Database connection
│   ├── models/
│   │   ├── news.go                # News model
│   │   ├── opinion.go             # Opinion model
│   │   ├── document.go            # Document model
│   │   ├── hero_slide.go          # Hero slide model
│   │   ├── organization.go        # Organization model
│   │   ├── page.go                # Page model
│   │   ├── category.go            # Category model
│   │   ├── tag.go                 # Tag model
│   │   └── setting.go             # Setting model
│   ├── handlers/
│   │   ├── news_handler.go        # News handlers
│   │   ├── opinion_handler.go     # Opinion handlers
│   │   ├── document_handler.go    # Document handlers
│   │   ├── hero_handler.go        # Hero slide handlers
│   │   ├── org_handler.go         # Organization handlers
│   │   ├── page_handler.go        # Page handlers
│   │   ├── category_handler.go    # Category handlers
│   │   ├── search_handler.go      # Search handlers
│   │   └── setting_handler.go     # Setting handlers
│   ├── repositories/
│   │   ├── news_repository.go     # News repository
│   │   ├── opinion_repository.go  # Opinion repository
│   │   └── ...                    # Other repositories
│   ├── services/
│   │   ├── news_service.go        # News business logic
│   │   ├── opinion_service.go     # Opinion business logic
│   │   └── ...                    # Other services
│   ├── middleware/
│   │   ├── cors.go                # CORS middleware
│   │   ├── logger.go              # Logger middleware
│   │   ├── rate_limit.go          # Rate limiter
│   │   └── cache.go               # Cache middleware
│   ├── utils/
│   │   ├── response.go            # Response helpers
│   │   ├── pagination.go          # Pagination helpers
│   │   └── validator.go           # Validation helpers
│   └── routes/
│       └── routes.go              # API routes
├── pkg/
│   ├── cache/
│   │   └── redis.go               # Redis cache (optional)
│   └── logger/
│       └── logger.go              # Logger configuration
├── migrations/
│   └── ...                        # Database migrations
├── config.yaml                    # Configuration file
├── config.example.yaml            # Configuration example
├── .gitignore
├── go.mod
├── go.sum
├── Makefile                       # Build commands
└── README.md
```

---

## 🔧 Installation & Setup

### Prerequisites
```bash
- Go 1.21+
- MySQL 8.0+
- Redis (optional, for caching)
```

### Configuration File (config.yaml)
```yaml
# Application Configuration
app:
  name: "LP Maarif API"
  env: "development"  # development, staging, production
  port: 8080
  version: "1.0.0"

# Database Configuration
database:
  host: "localhost"
  port: 3306
  user: "root"
  password: "password"
  name: "lp_maarif_nu"
  max_idle_conns: 10
  max_open_conns: 100
  conn_max_lifetime: 3600  # seconds

# CORS Configuration
cors:
  allowed_origins:
    - "http://localhost:3000"
    - "https://lpmaarifnu.or.id"
  allowed_methods:
    - "GET"
    - "OPTIONS"
  allowed_headers:
    - "Content-Type"
    - "Authorization"

# Cache Configuration (Optional)
cache:
  enabled: true
  type: "redis"  # redis or memory
  redis:
    host: "localhost"
    port: 6379
    password: ""
    db: 0
  ttl: 3600  # seconds

# Rate Limiting Configuration
rate_limit:
  enabled: true
  requests: 100  # max requests
  window: 60     # time window in seconds

# CDN/Storage Configuration
storage:
  cdn_url: "https://cdn.lpmaarifnu.or.id"

# Logging Configuration
logging:
  level: "info"  # debug, info, warn, error
  format: "json" # json or text
```

### Installation Steps

1. **Clone Repository**
```bash
git clone https://github.com/lpmaarifnu/api.git
cd lpmaarifnu-site-api
```

2. **Install Dependencies**
```bash
go mod download
```

3. **Setup Configuration File**
```bash
cp config.example.yaml config.yaml
# Edit config.yaml with your configuration
```

4. **Setup MySQL Database**
```bash
# Create database
mysql -u root -p -e "CREATE DATABASE lp_maarif_nu;"

# Import database schema
mysql -u root -p lp_maarif_nu < database_schema.sql
```

5. **Setup Redis (Optional)**
```bash
# Install Redis
# Ubuntu/Debian:
sudo apt install redis-server

# macOS:
brew install redis

# Start Redis
redis-server
```

6. **Run Application**
```bash
# Development mode
go run cmd/api/main.go

# Build for production
go build -o bin/api cmd/api/main.go

# Run production binary
./bin/api
```

### Configuration Loading Example

Create `internal/config/config.go`:

```go
package config

import (
    "github.com/spf13/viper"
    "log"
)

type Config struct {
    App      AppConfig      `mapstructure:"app"`
    Database DatabaseConfig `mapstructure:"database"`
    CORS     CORSConfig     `mapstructure:"cors"`
    Cache    CacheConfig    `mapstructure:"cache"`
    RateLimit RateLimitConfig `mapstructure:"rate_limit"`
    Storage  StorageConfig  `mapstructure:"storage"`
    Logging  LoggingConfig  `mapstructure:"logging"`
}

type AppConfig struct {
    Name    string `mapstructure:"name"`
    Env     string `mapstructure:"env"`
    Port    int    `mapstructure:"port"`
    Version string `mapstructure:"version"`
}

type DatabaseConfig struct {
    Host            string `mapstructure:"host"`
    Port            int    `mapstructure:"port"`
    User            string `mapstructure:"user"`
    Password        string `mapstructure:"password"`
    Name            string `mapstructure:"name"`
    MaxIdleConns    int    `mapstructure:"max_idle_conns"`
    MaxOpenConns    int    `mapstructure:"max_open_conns"`
    ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

type CORSConfig struct {
    AllowedOrigins []string `mapstructure:"allowed_origins"`
    AllowedMethods []string `mapstructure:"allowed_methods"`
    AllowedHeaders []string `mapstructure:"allowed_headers"`
}

type CacheConfig struct {
    Enabled bool        `mapstructure:"enabled"`
    Type    string      `mapstructure:"type"`
    Redis   RedisConfig `mapstructure:"redis"`
    TTL     int         `mapstructure:"ttl"`
}

type RedisConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Password string `mapstructure:"password"`
    DB       int    `mapstructure:"db"`
}

type RateLimitConfig struct {
    Enabled  bool `mapstructure:"enabled"`
    Requests int  `mapstructure:"requests"`
    Window   int  `mapstructure:"window"`
}

type StorageConfig struct {
    CDNURL string `mapstructure:"cdn_url"`
}

type LoggingConfig struct {
    Level  string `mapstructure:"level"`
    Format string `mapstructure:"format"`
}

func Load() (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AddConfigPath("./config")

    if err := viper.ReadInConfig(); err != nil {
        log.Printf("Error reading config file: %v", err)
        return nil, err
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        log.Printf("Error unmarshaling config: %v", err)
        return nil, err
    }

    return &config, nil
}
```

### Usage in main.go

```go
package main

import (
    "fmt"
    "log"

    "lpmaarifnu-site-api/internal/config"
    "github.com/gin-gonic/gin"
)

func main() {
    // Load configuration
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    // Initialize Gin
    if cfg.App.Env == "production" {
        gin.SetMode(gin.ReleaseMode)
    }

    router := gin.Default()

    // Start server
    addr := fmt.Sprintf(":%d", cfg.App.Port)
    log.Printf("Starting %s on %s", cfg.App.Name, addr)

    if err := router.Run(addr); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
```

---

## 📊 Performance Optimization

### 1. Database Indexing
```sql
-- Already included in database_schema.sql
CREATE INDEX idx_news_published ON news_articles(status, published_at, is_featured);
CREATE INDEX idx_opinion_published ON opinion_articles(status, published_at, is_featured);
```

### 2. Response Caching
```go
// Cache news list for 5 minutes
cache.Set("news:page:1", newsData, 5*time.Minute)
```

### 3. Database Connection Pooling
```go
// config/database.go
db.SetMaxIdleConns(10)
db.SetMaxOpenConns(100)
db.SetConnMaxLifetime(time.Hour)
```

### 4. Gzip Compression
```go
// Already handled by Gin middleware
router.Use(gzip.Gzip(gzip.DefaultCompression))
```

## 📝 API Documentation

### Swagger/OpenAPI
Generate API documentation:
```bash
# Install swag
go install github.com/swaggo/swag/cmd/swag@latest

# Generate docs
swag init -g cmd/api/main.go
```

Access documentation:
```
http://localhost:8080/swagger/index.html
```


---

## 📞 Support

For questions or issues:
- GitHub Issues: https://github.com/lpmaarifnu/api/issues
- Email: dev@lpmaarifnu.or.id

---

**Last Updated:** 2025-01-11
**Version:** 1.0.0
**Stack:** Golang + Gin + MySQL
