# API Examples - Real Response Data

Complete API examples with **real response data** from seeded database. Ready to copy-paste and use!

**Base URL Development:** `http://localhost:8080`

**Base URL Production:** `https://api.lpmaarifnu.or.id`

---

## 🚀 Quick Start

1. Import database seeder: `mysql -u root -p lp_maarif_nu < database_seeder.sql`
2. Start API: `go run cmd/api/main.go`
3. Test endpoints below - they will return real data!

---

## 📋 Table of Contents

1. [Health Check](#1-health-check)
2. [News Articles](#2-news-articles)
3. [Opinion Articles](#3-opinion-articles)
4. [Documents](#4-documents)
5. [Hero Slides](#5-hero-slides)
6. [Organization](#6-organization)
7. [Pages](#7-pages)
8. [Categories](#8-categories)
9. [Settings](#9-settings)
10. [Editorial Team](#10-editorial-team) ⭐ NEW
11. [Contact Form](#11-contact-form) ⭐ NEW
12. [Event Flayers](#12-event-flayers) ⭐ NEW

---

## 1. Health Check

### GET /health

**Request:**
```bash
curl http://localhost:8080/health
```

**Response:** `200 OK`
```json
{
  "status": "OK",
  "message": "LP Maarif API is running"
}
```

---

## 2. News Articles

### GET /api/v1/news - Get All News

**Request:**
```bash
curl "http://localhost:8080/api/v1/news?page=1&limit=5"
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "News articles retrieved successfully",
  "data": {
    "articles": [
      {
        "id": 2,
        "title": "Implementasi Kurikulum Merdeka di Sekolah Ma'arif Se-Indonesia",
        "slug": "implementasi-kurikulum-merdeka-sekolah-maarif",
        "excerpt": "Lebih dari 5000 sekolah Ma'arif di seluruh Indonesia telah berhasil mengimplementasikan Kurikulum Merdeka dengan pendampingan intensif dari LP Ma'arif NU.",
        "content": "<h2>Kurikulum Merdeka di Sekolah Ma'arif</h2><p>Dalam upaya meningkatkan kualitas pendidikan...</p>",
        "image": "https://images.unsplash.com/photo-1503676260728-1c00da094a0b?w=800",
        "published_at": "2025-01-06T00:00:00Z",
        "category": {
          "id": 1,
          "name": "Nasional",
          "slug": "nasional"
        },
        "author": {
          "id": 1,
          "name": "Super Admin",
          "avatar": null
        },
        "tags": [
          {
            "id": 1,
            "name": "Pendidikan",
            "slug": "pendidikan"
          },
          {
            "id": 3,
            "name": "Kurikulum",
            "slug": "kurikulum"
          },
          {
            "id": 12,
            "name": "Inovasi",
            "slug": "inovasi"
          }
        ],
        "views": 2340,
        "is_featured": true,
        "created_at": "2025-01-06T00:00:00Z",
        "updated_at": "2025-01-06T00:00:00Z"
      },
      {
        "id": 3,
        "title": "Pelatihan Guru Digital untuk Meningkatkan Kualitas Pembelajaran",
        "slug": "pelatihan-guru-digital-kualitas-pembelajaran",
        "excerpt": "LP Ma'arif NU menyelenggarakan pelatihan guru digital yang diikuti oleh 2000 guru dari berbagai daerah untuk meningkatkan kompetensi dalam pembelajaran berbasis teknologi.",
        "image": "https://images.unsplash.com/photo-1516321318423-f06f85e504b3?w=800",
        "published_at": "2025-01-04T00:00:00Z",
        "category": {
          "id": 2,
          "name": "Daerah",
          "slug": "daerah"
        },
        "author": {
          "id": 2,
          "name": "Editor Berita",
          "avatar": "https://ui-avatars.com/api/?name=Editor+Berita"
        },
        "tags": [
          {
            "id": 1,
            "name": "Pendidikan",
            "slug": "pendidikan"
          },
          {
            "id": 4,
            "name": "Pelatihan",
            "slug": "pelatihan"
          },
          {
            "id": 9,
            "name": "Teknologi",
            "slug": "teknologi"
          }
        ],
        "views": 1890,
        "is_featured": true,
        "created_at": "2025-01-04T00:00:00Z",
        "updated_at": "2025-01-04T00:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 2,
      "total_items": 7,
      "items_per_page": 5,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

---

### GET /api/v1/news/featured - Get Featured News

**Request:**
```bash
curl "http://localhost:8080/api/v1/news/featured?limit=3"
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Featured news retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "Peluncuran Program Beasiswa Unggulan 2024 untuk Siswa Berprestasi",
      "slug": "peluncuran-program-beasiswa-unggulan-2024",
      "excerpt": "LP Ma'arif NU meluncurkan program beasiswa unggulan tahun 2024 yang ditujukan untuk siswa berprestasi dari keluarga kurang mampu di seluruh Indonesia.",
      "image": "https://images.unsplash.com/photo-1523050854058-8df90110c9f1?w=800",
      "published_at": "2025-01-09T00:00:00Z",
      "category": {
        "id": 1,
        "name": "Nasional",
        "slug": "nasional"
      },
      "views": 1520,
      "tags": [
        {
          "id": 1,
          "name": "Pendidikan",
          "slug": "pendidikan"
        },
        {
          "id": 2,
          "name": "Beasiswa",
          "slug": "beasiswa"
        }
      ]
    },
    {
      "id": 2,
      "title": "Implementasi Kurikulum Merdeka di Sekolah Ma'arif Se-Indonesia",
      "slug": "implementasi-kurikulum-merdeka-sekolah-maarif",
      "excerpt": "Lebih dari 5000 sekolah Ma'arif di seluruh Indonesia telah berhasil mengimplementasikan Kurikulum Merdeka dengan pendampingan intensif dari LP Ma'arif NU.",
      "image": "https://images.unsplash.com/photo-1503676260728-1c00da094a0b?w=800",
      "published_at": "2025-01-06T00:00:00Z",
      "category": {
        "id": 1,
        "name": "Nasional",
        "slug": "nasional"
      },
      "views": 2340
    },
    {
      "id": 3,
      "title": "Pelatihan Guru Digital untuk Meningkatkan Kualitas Pembelajaran",
      "slug": "pelatihan-guru-digital-kualitas-pembelajaran",
      "excerpt": "LP Ma'arif NU menyelenggarakan pelatihan guru digital yang diikuti oleh 2000 guru dari berbagai daerah untuk meningkatkan kompetensi dalam pembelajaran berbasis teknologi.",
      "image": "https://images.unsplash.com/photo-1516321318423-f06f85e504b3?w=800",
      "published_at": "2025-01-04T00:00:00Z",
      "category": {
        "id": 2,
        "name": "Daerah",
        "slug": "daerah"
      },
      "views": 1890
    }
  ]
}
```

---

### GET /api/v1/news/:slug - Get Single News

**Request:**
```bash
curl "http://localhost:8080/api/v1/news/peluncuran-program-beasiswa-unggulan-2024"
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "News article retrieved successfully",
  "data": {
    "id": 1,
    "title": "Peluncuran Program Beasiswa Unggulan 2024 untuk Siswa Berprestasi",
    "slug": "peluncuran-program-beasiswa-unggulan-2024",
    "excerpt": "LP Ma'arif NU meluncurkan program beasiswa unggulan tahun 2024 yang ditujukan untuk siswa berprestasi dari keluarga kurang mampu di seluruh Indonesia.",
    "content": "<h2>Program Beasiswa Unggulan 2024</h2><p>LP Ma'arif NU dengan bangga mengumumkan peluncuran Program Beasiswa Unggulan 2024...</p><h3>Kriteria Penerima</h3><ul><li>Siswa aktif jenjang SMP/MTs dan SMA/MA</li><li>Memiliki prestasi akademik minimal rata-rata 8.0</li></ul>",
    "image": "https://images.unsplash.com/photo-1523050854058-8df90110c9f1?w=800",
    "published_at": "2025-01-09T00:00:00Z",
    "category": {
      "id": 1,
      "name": "Nasional",
      "slug": "nasional"
    },
    "author": {
      "id": 1,
      "name": "Super Admin",
      "avatar": null
    },
    "tags": [
      {
        "id": 1,
        "name": "Pendidikan",
        "slug": "pendidikan"
      },
      {
        "id": 2,
        "name": "Beasiswa",
        "slug": "beasiswa"
      },
      {
        "id": 8,
        "name": "Prestasi",
        "slug": "prestasi"
      }
    ],
    "views": 1521,
    "is_featured": true,
    "meta": {
      "title": "Program Beasiswa Unggulan 2024 LP Ma'arif NU",
      "description": "LP Ma'arif NU meluncurkan program beasiswa unggulan 2024 untuk siswa berprestasi. Daftar sekarang!",
      "keywords": "beasiswa, pendidikan, ma'arif nu, siswa berprestasi, 2024"
    },
    "related_articles": [
      {
        "id": 2,
        "title": "Implementasi Kurikulum Merdeka di Sekolah Ma'arif Se-Indonesia",
        "slug": "implementasi-kurikulum-merdeka-sekolah-maarif",
        "image": "https://images.unsplash.com/photo-1503676260728-1c00da094a0b?w=800",
        "published_at": "2025-01-06T00:00:00Z"
      }
    ],
    "created_at": "2025-01-09T00:00:00Z",
    "updated_at": "2025-01-09T00:00:00Z"
  }
}
```

---

### GET /api/v1/news?category=nasional - Filter by Category

**Request:**
```bash
curl "http://localhost:8080/api/v1/news?category=nasional&page=1&limit=3"
```

**Response:** `200 OK` - Returns only news with category "Nasional"

---

### GET /api/v1/news?search=pendidikan - Search News

**Request:**
```bash
curl "http://localhost:8080/api/v1/news?search=pendidikan&page=1&limit=5"
```

**Response:** `200 OK` - Returns news containing "pendidikan" in title or excerpt

---

## 3. Opinion Articles

### GET /api/v1/opinions - Get All Opinions

**Request:**
```bash
curl "http://localhost:8080/api/v1/opinions?page=1&limit=5"
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Opinion articles retrieved successfully",
  "data": {
    "articles": [
      {
        "id": 1,
        "title": "Pendidikan Karakter di Era Digital: Tantangan dan Solusi",
        "slug": "pendidikan-karakter-era-digital",
        "excerpt": "Pentingnya menanamkan nilai-nilai karakter kepada generasi muda di tengah gempuran teknologi digital yang semakin masif.",
        "content": "<h2>Pendidikan Karakter di Era Digital</h2><p>Era digital membawa perubahan besar...</p>",
        "image": "https://images.unsplash.com/photo-1509062522246-3755977927d7?w=800",
        "published_at": "2025-01-08T00:00:00Z",
        "author_name": "Prof. Dr. KH. Ahmad Syafi'i, MA",
        "author_title": "Pakar Pendidikan Islam",
        "author_image": "https://ui-avatars.com/api/?name=Ahmad+Syafii&background=4CAF50&color=fff&size=200",
        "author_bio": "Pakar pendidikan Islam dengan pengalaman lebih dari 25 tahun di bidang pendidikan. Aktif menulis dan memberikan ceramah tentang pendidikan karakter.",
        "tags": [
          {
            "id": 10,
            "name": "Karakter",
            "slug": "karakter"
          },
          {
            "id": 9,
            "name": "Teknologi",
            "slug": "teknologi"
          },
          {
            "id": 1,
            "name": "Pendidikan",
            "slug": "pendidikan"
          }
        ],
        "views": 890,
        "is_featured": true,
        "created_at": "2025-01-08T00:00:00Z",
        "updated_at": "2025-01-08T00:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 1,
      "total_items": 3,
      "items_per_page": 5,
      "has_next": false,
      "has_prev": false
    }
  }
}
```

---

### GET /api/v1/opinions/:slug - Get Single Opinion

**Request:**
```bash
curl "http://localhost:8080/api/v1/opinions/pendidikan-karakter-era-digital"
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Opinion article retrieved successfully",
  "data": {
    "id": 1,
    "title": "Pendidikan Karakter di Era Digital: Tantangan dan Solusi",
    "slug": "pendidikan-karakter-era-digital",
    "excerpt": "Pentingnya menanamkan nilai-nilai karakter kepada generasi muda di tengah gempuran teknologi digital yang semakin masif.",
    "content": "<h2>Pendidikan Karakter di Era Digital</h2><p>Era digital membawa perubahan besar dalam cara kita mendidik generasi muda...</p>",
    "image": "https://images.unsplash.com/photo-1509062522246-3755977927d7?w=800",
    "published_at": "2025-01-08T00:00:00Z",
    "author": {
      "name": "Prof. Dr. KH. Ahmad Syafi'i, MA",
      "title": "Pakar Pendidikan Islam",
      "image": "https://ui-avatars.com/api/?name=Ahmad+Syafii&background=4CAF50&color=fff&size=200",
      "bio": "Pakar pendidikan Islam dengan pengalaman lebih dari 25 tahun di bidang pendidikan. Aktif menulis dan memberikan ceramah tentang pendidikan karakter."
    },
    "tags": [
      {
        "id": 10,
        "name": "Karakter",
        "slug": "karakter"
      }
    ],
    "views": 891,
    "is_featured": true,
    "meta": {
      "title": "Pendidikan Karakter di Era Digital",
      "description": "Pentingnya pendidikan karakter di era digital dengan solusi praktis menghadapi tantangan teknologi.",
      "keywords": "pendidikan, karakter, digital, generasi muda"
    },
    "created_at": "2025-01-08T00:00:00Z",
    "updated_at": "2025-01-08T00:00:00Z"
  }
}
```

---

## 4. Documents

### GET /api/v1/documents - Get All Documents

**Request:**
```bash
curl "http://localhost:8080/api/v1/documents?page=1&limit=5"
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Documents retrieved successfully",
  "data": {
    "documents": [
      {
        "id": 1,
        "title": "Pedoman Penyelenggaraan Pendidikan Ma'arif NU 2024",
        "description": "Pedoman lengkap penyelenggaraan pendidikan di lingkungan Ma'arif NU yang mencakup standar operasional, kurikulum, dan tata kelola sekolah.",
        "category": {
          "id": 5,
          "name": "Pedoman",
          "slug": "pedoman"
        },
        "file_name": "pedoman-pendidikan-maarif-2024.pdf",
        "file_type": "PDF",
        "file_size": 2621440,
        "file_size_formatted": "2.5 MB",
        "download_url": "https://cdn.lpmaarifnu.or.id/documents/pedoman-pendidikan-maarif-2024.pdf",
        "download_count": 1534,
        "uploaded_at": "2025-01-11T00:00:00Z",
        "created_at": "2025-01-11T00:00:00Z"
      },
      {
        "id": 5,
        "title": "Panduan Implementasi Kurikulum Merdeka di Sekolah Ma'arif",
        "description": "Panduan praktis implementasi Kurikulum Merdeka dengan tetap mempertahankan nilai-nilai Ma'arif NU.",
        "category": {
          "id": 8,
          "name": "Panduan",
          "slug": "panduan"
        },
        "file_name": "panduan-kurikulum-merdeka.pdf",
        "file_type": "PDF",
        "file_size": 2097152,
        "file_size_formatted": "2.0 MB",
        "download_url": "https://cdn.lpmaarifnu.or.id/documents/panduan-kurikulum-merdeka.pdf",
        "download_count": 3456,
        "uploaded_at": "2025-01-11T00:00:00Z",
        "created_at": "2025-01-11T00:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "total_pages": 2,
      "total_items": 10,
      "items_per_page": 5,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

---

### GET /api/v1/documents/:id - Get Single Document

**Request:**
```bash
curl "http://localhost:8080/api/v1/documents/1"
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Document retrieved successfully",
  "data": {
    "id": 1,
    "title": "Pedoman Penyelenggaraan Pendidikan Ma'arif NU 2024",
    "description": "Pedoman lengkap penyelenggaraan pendidikan di lingkungan Ma'arif NU yang mencakup standar operasional, kurikulum, dan tata kelola sekolah.",
    "category": {
      "id": 5,
      "name": "Pedoman",
      "slug": "pedoman"
    },
    "file_name": "pedoman-pendidikan-maarif-2024.pdf",
    "file_type": "PDF",
    "file_size": 2621440,
    "file_size_formatted": "2.5 MB",
    "mime_type": "application/pdf",
    "download_url": "https://cdn.lpmaarifnu.or.id/documents/pedoman-pendidikan-maarif-2024.pdf",
    "download_count": 1535,
    "uploaded_at": "2025-01-11T00:00:00Z",
    "created_at": "2025-01-11T00:00:00Z",
    "updated_at": "2025-01-11T00:00:00Z"
  }
}
```

---

## 5. Hero Slides

### GET /api/v1/hero-slides - Get Active Hero Slides

**Request:**
```bash
curl "http://localhost:8080/api/v1/hero-slides"
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Hero slides retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "Membangun Pendidikan Islam Berkualitas dan Berkarakter",
      "description": "LP Ma'arif NU berkomitmen mengembangkan pendidikan Islam yang berkualitas dengan mengintegrasikan nilai-nilai keislaman, kearifan lokal, dan kemajuan teknologi untuk membentuk generasi yang berakhlak mulia dan berdaya saing tinggi.",
      "image": "https://images.unsplash.com/photo-1427504494785-3a9ca7044f45?w=1920",
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
      "created_at": "2024-12-12T00:00:00Z"
    },
    {
      "id": 2,
      "title": "Program Beasiswa Unggulan 2024",
      "description": "Raih kesempatan emas untuk mendapatkan beasiswa pendidikan penuh! Daftar sekarang dan wujudkan impian pendidikanmu bersama LP Ma'arif NU.",
      "image": "https://images.unsplash.com/photo-1523050854058-8df90110c9f1?w=1920",
      "cta": {
        "label": "Daftar Sekarang",
        "href": "/beasiswa",
        "secondary": {
          "label": "Info Lengkap",
          "href": "/berita/peluncuran-program-beasiswa-unggulan-2024"
        }
      },
      "order": 2,
      "is_active": true,
      "created_at": "2024-12-22T00:00:00Z"
    }
  ]
}
```

---

## 6. Organization

### GET /api/v1/organization/structure - Get Organization Structure

**Request:**
```bash
curl "http://localhost:8080/api/v1/organization/structure"
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Organization structure retrieved successfully",
  "data": {
    "ketua": {
      "id": 1,
      "name": "Prof. Dr. KH. Said Aqil Siradj, MA",
      "position": "Ketua Umum",
      "photo": "https://ui-avatars.com/api/?name=Said+Aqil+Siradj&background=1976D2&color=fff&size=400",
      "bio": "Ketua Umum LP Ma'arif NU periode 2024-2029. Beliau adalah ulama dan akademisi terkemuka yang memiliki dedikasi tinggi dalam pengembangan pendidikan Islam di Indonesia.",
      "period": "2024-2029",
      "email": "ketua@lpmaarifnu.or.id",
      "phone": "021-12345678"
    },
    "wakil": [
      {
        "id": 2,
        "name": "Dr. H. Ahmad Lutfi, M.Pd",
        "position": "Wakil Ketua I",
        "photo": "https://ui-avatars.com/api/?name=Ahmad+Lutfi&background=388E3C&color=fff&size=400",
        "bio": "Wakil Ketua I LP Ma'arif NU yang membidangi Pendidikan Dasar dan Menengah. Berpengalaman dalam manajemen pendidikan lebih dari 20 tahun.",
        "period": "2024-2029"
      }
    ],
    "sekretaris": {
      "id": 4,
      "name": "Dr. Hj. Siti Aisyah, M.Pd",
      "position": "Sekretaris Umum",
      "photo": "https://ui-avatars.com/api/?name=Siti+Aisyah&background=C62828&color=fff&size=400",
      "bio": "Sekretaris Umum LP Ma'arif NU yang bertanggung jawab atas administrasi dan koordinasi kegiatan organisasi.",
      "period": "2024-2029"
    },
    "bendahara": {
      "id": 5,
      "name": "H. Abdul Rahman, SE, M.Ak",
      "position": "Bendahara Umum",
      "photo": "https://ui-avatars.com/api/?name=Abdul+Rahman&background=7B1FA2&color=fff&size=400",
      "bio": "Bendahara Umum yang mengelola keuangan dan aset LP Ma'arif NU dengan prinsip transparansi dan akuntabilitas.",
      "period": "2024-2029"
    },
    "departments": [
      {
        "id": 1,
        "name": "Bidang Pendidikan Dasar",
        "head_name": null,
        "description": "Mengelola pengembangan pendidikan tingkat SD/MI"
      }
    ]
  }
}
```

---

### GET /api/v1/organization/board-members - Get Board Members

**Request:**
```bash
curl "http://localhost:8080/api/v1/organization/board-members"
```

**Response:** `200 OK` - Returns all active board members with their positions

---

### GET /api/v1/organization/pengurus - Get Pengurus List

**Request:**
```bash
curl "http://localhost:8080/api/v1/organization/pengurus?periode=2024-2029"
```

**Response:** `200 OK`
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
        "bio": "Ulama dan intelektual muslim Indonesia, Ketua Umum PBNU periode 2021-2026. Beliau memiliki visi untuk memajukan pendidikan Ma'arif NU di seluruh Indonesia.",
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
        "nama": "Dr. H. Muktafi, M.Pd.I",
        "jabatan": "Wakil Ketua I",
        "kategori": "pimpinan_utama",
        "foto": "https://ui-avatars.com/api/?name=Muktafi",
        "bio": "Pakar pendidikan Islam dengan pengalaman lebih dari 20 tahun di bidang manajemen pendidikan.",
        "email": "wakil1@lpmaarifnu.or.id",
        "phone": "021-3920678",
        "periode": {
          "mulai": 2024,
          "selesai": 2029
        },
        "order": 2,
        "is_active": true
      },
      {
        "id": 3,
        "nama": "Prof. Dr. Hj. Masriyah, M.Pd",
        "jabatan": "Wakil Ketua II",
        "kategori": "pimpinan_utama",
        "foto": "https://ui-avatars.com/api/?name=Masriyah",
        "bio": "Guru Besar Universitas Negeri Surabaya, ahli dalam bidang pengembangan kurikulum pendidikan.",
        "email": "wakil2@lpmaarifnu.or.id",
        "phone": "021-3920679",
        "periode": {
          "mulai": 2024,
          "selesai": 2029
        },
        "order": 3,
        "is_active": true
      }
    ]
  }
}
```

---

### GET /api/v1/organization/pengurus?kategori=pimpinan_utama - Filter by Kategori

**Request:**
```bash
curl "http://localhost:8080/api/v1/organization/pengurus?periode=2024-2029&kategori=pimpinan_utama"
```

**Response:** `200 OK` - Returns only Pimpinan Utama (Ketua & Wakil Ketua)

---

### GET /api/v1/organization/pengurus?kategori=bidang - Get Department Heads

**Request:**
```bash
curl "http://localhost:8080/api/v1/organization/pengurus?kategori=bidang"
```

**Response:** `200 OK` - Returns all department heads (Kepala Bidang)

---

## 7. Pages

### GET /api/v1/pages/visi-misi - Get Visi Misi Page

**Request:**
```bash
curl "http://localhost:8080/api/v1/pages/visi-misi"
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Page content retrieved successfully",
  "data": {
    "slug": "visi-misi",
    "title": "Visi & Misi",
    "template": "visi-misi",
    "content": {
      "visi": "Menjadi lembaga pendidikan Islam yang unggul, modern, dan berkarakter untuk mewujudkan generasi yang berakhlak mulia, cerdas, dan berdaya saing global.",
      "misi": [
        "Menyelenggarakan pendidikan berkualitas yang berbasis nilai-nilai keislaman",
        "Mengembangkan kurikulum yang integratif antara ilmu agama dan ilmu umum",
        "Memberdayakan sumber daya manusia yang profesional dan berakhlak mulia",
        "Membangun kerjasama dengan berbagai pihak untuk kemajuan pendidikan",
        "Mengimplementasikan teknologi dalam pembelajaran"
      ],
      "nilai_nilai": [
        {
          "icon": "integrity",
          "title": "Integritas",
          "description": "Menjunjung tinggi kejujuran dan tanggung jawab dalam setiap aspek"
        },
        {
          "icon": "professionalism",
          "title": "Profesionalisme",
          "description": "Menjalankan tugas dengan kompeten dan penuh dedikasi"
        }
      ]
    },
    "meta": {
      "title": null,
      "description": null,
      "keywords": null
    },
    "is_active": true,
    "updated_at": "2025-01-11T00:00:00Z"
  }
}
```

---

### GET /api/v1/pages/sejarah - Get Sejarah Page

**Request:**
```bash
curl "http://localhost:8080/api/v1/pages/sejarah"
```

**Response:** `200 OK` - Returns history timeline from 1916 to 2024

---

## 8. Categories

### GET /api/v1/categories - Get All Categories

**Request:**
```bash
curl "http://localhost:8080/api/v1/categories"
```

**Response:** `200 OK`
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
      "icon": null,
      "color": null,
      "is_active": true,
      "order_number": 1,
      "created_at": "2025-01-11T00:00:00Z",
      "updated_at": "2025-01-11T00:00:00Z"
    },
    {
      "id": 2,
      "name": "Daerah",
      "slug": "daerah",
      "description": "Berita tingkat daerah",
      "type": "news",
      "icon": null,
      "color": null,
      "is_active": true,
      "order_number": 2
    }
  ]
}
```

---

### GET /api/v1/categories?type=news - Get News Categories

**Request:**
```bash
curl "http://localhost:8080/api/v1/categories?type=news"
```

**Response:** `200 OK` - Returns only categories with type "news"

---

### GET /api/v1/categories/:slug - Get Single Category

**Request:**
```bash
curl "http://localhost:8080/api/v1/categories/nasional"
```

**Response:** `200 OK`
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
    "icon": null,
    "color": null,
    "latest_articles": [
      {
        "id": 1,
        "title": "Peluncuran Program Beasiswa Unggulan 2024 untuk Siswa Berprestasi",
        "slug": "peluncuran-program-beasiswa-unggulan-2024",
        "excerpt": "LP Ma'arif NU meluncurkan program beasiswa unggulan tahun 2024...",
        "image": "https://images.unsplash.com/photo-1523050854058-8df90110c9f1?w=800",
        "published_at": "2025-01-09T00:00:00Z"
      }
    ]
  }
}
```

---

## 9. Settings

### GET /api/v1/settings - Get Public Settings

**Request:**
```bash
curl "http://localhost:8080/api/v1/settings"
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Settings retrieved successfully",
  "data": {
    "site_name": "LP Ma'arif NU",
    "site_description": "Lembaga Pendidikan Ma'arif Nahdlatul Ulama - Membangun Pendidikan Islam Berkualitas",
    "logo": "",
    "favicon": "",
    "contact": {
      "email": "info@lpmaarifnu.or.id",
      "phone": "(021) 3920677",
      "address": "Jl. Kramat Raya No. 164, Jakarta Pusat 10430"
    },
    "social_media": {
      "facebook": "https://facebook.com/lpmaarifnu",
      "twitter": "https://twitter.com/lpmaarifnu",
      "instagram": "https://instagram.com/lpmaarifnu",
      "youtube": "https://youtube.com/@lpmaarifnu"
    }
  }
}
```

---

## 10. Editorial Team

### GET /api/v1/editorial/team - Get Complete Editorial Team

**Request:**
```bash
curl "http://localhost:8080/api/v1/editorial/team"
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Editorial team retrieved successfully",
  "data": {
    "pemimpin_redaksi": {
      "name": "Dr. H. Muhammad Fadhil, M.Pd",
      "position": "Pemimpin Redaksi",
      "photo": "https://ui-avatars.com/api/?name=Muhammad+Fadhil",
      "bio": "Pakar pendidikan Islam dengan pengalaman lebih dari 15 tahun di bidang pendidikan dan jurnalisme pendidikan.",
      "email": "fadhil@lpmaarifnu.or.id",
      "phone": "021-12345678"
    },
    "wakil_pemimpin_redaksi": [
      {
        "name": "Dra. Hj. Nur Azizah, M.Si",
        "position": "Wakil Pemimpin Redaksi I",
        "photo": "https://ui-avatars.com/api/?name=Nur+Azizah",
        "bio": "Spesialis media dan komunikasi publik dengan fokus pada isu-isu pendidikan Islam.",
        "email": "azizah@lpmaarifnu.or.id"
      },
      {
        "name": "Dr. Zainuddin Rahman, M.A",
        "position": "Wakil Pemimpin Redaksi II",
        "photo": "https://ui-avatars.com/api/?name=Zainuddin+Rahman",
        "bio": "Akademisi dan praktisi media dengan keahlian di bidang jurnalisme investigasi pendidikan.",
        "email": "zainuddin@lpmaarifnu.or.id"
      }
    ],
    "redaktur_pelaksana": {
      "name": "Ahmad Syarif, S.Sos, M.I.Kom",
      "position": "Redaktur Pelaksana",
      "photo": "https://ui-avatars.com/api/?name=Ahmad+Syarif",
      "bio": "Koordinator harian tim redaksi yang mengatur alur produksi konten dan publikasi.",
      "email": "syarif@lpmaarifnu.or.id"
    },
    "dewan_redaksi": [
      {
        "id": 1,
        "name": "Prof. Dr. KH. Abdullah Shiddiq, MA",
        "institution": "UIN Syarif Hidayatullah Jakarta",
        "expertise": "Pendidikan Islam & Budaya",
        "photo": "https://ui-avatars.com/api/?name=Abdullah+Shiddiq",
        "bio": "Guru Besar bidang Pendidikan Islam, aktif dalam riset dan pengembangan pendidikan berbasis budaya lokal.",
        "email": "abdullah.shiddiq@uinjkt.ac.id",
        "order_number": 1,
        "is_active": true
      },
      {
        "id": 2,
        "name": "Prof. Dr. Hj. Mahmudah Hani, M.Pd",
        "institution": "Universitas Negeri Malang",
        "expertise": "Manajemen Pendidikan",
        "photo": "https://ui-avatars.com/api/?name=Mahmudah+Hani",
        "bio": "Pakar manajemen pendidikan dengan pengalaman lebih dari 25 tahun dalam pengembangan sistem pendidikan.",
        "email": "mahmudah.hani@um.ac.id",
        "order_number": 2,
        "is_active": true
      },
      {
        "id": 3,
        "name": "Dr. H. Yusuf Hanafi, M.Ag",
        "institution": "UIN Maulana Malik Ibrahim Malang",
        "expertise": "Kurikulum & Pembelajaran",
        "photo": "https://ui-avatars.com/api/?name=Yusuf+Hanafi",
        "bio": "Spesialis kurikulum dan inovasi pembelajaran, aktif mengembangkan model pembelajaran inovatif.",
        "email": "yusuf.hanafi@uin-malang.ac.id",
        "order_number": 3,
        "is_active": true
      },
      {
        "id": 4,
        "name": "Dr. Hj. Siti Maesaroh, M.Pd.I",
        "institution": "IAIN Surakarta",
        "expertise": "Teknologi Pendidikan",
        "photo": "https://ui-avatars.com/api/?name=Siti+Maesaroh",
        "bio": "Ahli teknologi pendidikan yang fokus pada digitalisasi pembelajaran di era modern.",
        "email": "siti.maesaroh@iain-surakarta.ac.id",
        "order_number": 4,
        "is_active": true
      }
    ],
    "tim_redaksi": [
      {
        "id": 5,
        "name": "Rizki Aulia Rahman, S.Pd",
        "position": "Editor Berita",
        "role_type": "tim_redaksi",
        "photo": "https://ui-avatars.com/api/?name=Rizki+Rahman",
        "bio": "Editor berita dengan keahlian dalam meliput dan menulis berita pendidikan.",
        "email": "rizki@lpmaarifnu.or.id",
        "phone": "021-12345681"
      },
      {
        "id": 6,
        "name": "Fatimah Az-Zahra, S.Sos",
        "position": "Editor Opini",
        "role_type": "tim_redaksi",
        "photo": "https://ui-avatars.com/api/?name=Fatimah+Zahra",
        "bio": "Mengelola rubrik opini dan artikel mendalam tentang isu-isu pendidikan.",
        "email": "fatimah@lpmaarifnu.or.id",
        "phone": "021-12345682"
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

## 11. Contact Form

### POST /api/v1/contact/submit - Submit Contact Message

**Request:**
```bash
curl -X POST "http://localhost:8080/api/v1/contact/submit" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Ahmad Fauzi",
    "email": "ahmad.fauzi@example.com",
    "phone": "081234567890",
    "subject": "Pertanyaan tentang Pendaftaran Beasiswa 2024",
    "message": "Assalamualaikum, saya ingin menanyakan persyaratan dan prosedur pendaftaran beasiswa unggulan 2024. Saya siswa SMA kelas 12 dengan prestasi akademik rata-rata 9.0. Mohon informasi lengkapnya. Terima kasih."
  }'
```

**Response:** `201 Created`
```json
{
  "success": true,
  "message": "Pesan Anda telah terkirim. Kami akan segera menghubungi Anda.",
  "data": {
    "ticket_id": "CTK-2024-0005",
    "name": "Ahmad Fauzi",
    "email": "ahmad.fauzi@example.com",
    "subject": "Pertanyaan tentang Pendaftaran Beasiswa 2024",
    "status": "new",
    "created_at": "2025-01-14T15:30:45Z"
  }
}
```

---

### POST /api/v1/contact/submit - Validation Error Example

**Request with missing required field:**
```bash
curl -X POST "http://localhost:8080/api/v1/contact/submit" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Ahmad",
    "email": "invalid-email",
    "subject": "Test"
  }'
```

**Response:** `400 Bad Request`
```json
{
  "success": false,
  "message": "Validation failed",
  "error": "Key: 'ContactRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"
}
```

---

### Contact Form - Rate Limit Example

**Request after exceeding rate limit (6th request within 1 hour):**
```bash
curl -X POST "http://localhost:8080/api/v1/contact/submit" \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","email":"test@example.com","subject":"Test","message":"Test message"}'
```

**Response:** `429 Too Many Requests`
```json
{
  "success": false,
  "message": "Rate limit exceeded. Please try again later."
}
```

**Note:** Contact form has a strict rate limit of **5 requests per hour per IP address**.

---

## 12. Event Flayers

### GET /api/v1/events/flayers - Get All Event Flayers

**Request:**
```bash
curl "http://localhost:8080/api/v1/events/flayers?limit=10&active=true"
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Event flayers retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "Seminar Nasional Pendidikan Islam 2024",
      "description": "Seminar nasional dengan tema 'Transformasi Pendidikan Islam di Era Digital' yang menghadirkan para pakar pendidikan Islam dari berbagai universitas terkemuka.",
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
    },
    {
      "id": 2,
      "title": "Workshop Kurikulum Merdeka untuk Guru Ma'arif",
      "description": "Workshop intensif implementasi Kurikulum Merdeka di sekolah-sekolah Ma'arif dengan pendampingan langsung dari tim ahli.",
      "image": "https://images.unsplash.com/photo-1524178232363-1fb2b075b655?w=800",
      "event_date": "2024-04-20",
      "event_location": "Gedung LP Ma'arif NU Pusat",
      "registration_url": "https://forms.lpmaarifnu.or.id/workshop-kurikulum",
      "contact": {
        "person": "Tim Pengembangan Kurikulum",
        "phone": "021-3920680",
        "email": "kurikulum@lpmaarifnu.or.id"
      },
      "display_period": {
        "start": "2024-02-01T00:00:00Z",
        "end": "2024-04-19T23:59:59Z"
      },
      "order": 2,
      "is_active": true
    },
    {
      "id": 3,
      "title": "Jambore Pramuka Ma'arif Nasional 2024",
      "description": "Jambore tingkat nasional untuk seluruh pramuka sekolah Ma'arif di Indonesia dengan berbagai kegiatan kepramukaan dan pembinaan karakter.",
      "image": "https://images.unsplash.com/photo-1504384764586-bb4cdc1707b0?w=800",
      "event_date": "2024-07-10",
      "event_location": "Bumi Perkemahan Cibubur, Jakarta Timur",
      "registration_url": "https://forms.lpmaarifnu.or.id/jambore-pramuka",
      "contact": {
        "person": "Kwartir Pramuka Ma'arif",
        "phone": "021-3920685",
        "email": "pramuka@lpmaarifnu.or.id"
      },
      "display_period": {
        "start": "2024-04-01T00:00:00Z",
        "end": "2024-07-09T23:59:59Z"
      },
      "order": 3,
      "is_active": true
    }
  ]
}
```

---

### GET /api/v1/events/flayers?upcoming=true - Get Upcoming Events Only

**Request:**
```bash
curl "http://localhost:8080/api/v1/events/flayers?limit=5&upcoming=true"
```

**Response:** `200 OK` - Returns only events with `event_date` in the future

---

### GET /api/v1/events/flayers?limit=3 - Get Limited Events

**Request:**
```bash
curl "http://localhost:8080/api/v1/events/flayers?limit=3&active=true"
```

**Response:** `200 OK` - Returns maximum 3 active events, ordered by `order_number` and `event_date`

---

## 🎯 Frontend Implementation Examples

### React/Next.js Example

```javascript
// Fetch featured news
const fetchFeaturedNews = async () => {
  const response = await fetch('http://localhost:8080/api/v1/news/featured?limit=3');
  const data = await response.json();
  return data.data; // Array of featured news
};

// Fetch news with pagination
const fetchNews = async (page = 1, limit = 10) => {
  const response = await fetch(`http://localhost:8080/api/v1/news?page=${page}&limit=${limit}`);
  const data = await response.json();
  return {
    articles: data.data.articles,
    pagination: data.data.pagination
  };
};

// Fetch single news
const fetchNewsDetail = async (slug) => {
  const response = await fetch(`http://localhost:8080/api/v1/news/${slug}`);
  const data = await response.json();
  return data.data;
};

// Fetch pengurus
const fetchPengurus = async (periode = '2024-2029', kategori = '') => {
  const url = new URL('http://localhost:8080/api/v1/organization/pengurus');
  url.searchParams.set('periode', periode);
  if (kategori) url.searchParams.set('kategori', kategori);

  const response = await fetch(url.toString());
  const data = await response.json();
  return data.data;
};

// Fetch editorial team
const fetchEditorialTeam = async () => {
  const response = await fetch('http://localhost:8080/api/v1/editorial/team');
  const data = await response.json();
  return data.data;
};

// Submit contact form
const submitContact = async (formData) => {
  const response = await fetch('http://localhost:8080/api/v1/contact/submit', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(formData)
  });
  const data = await response.json();
  if (!data.success) throw new Error(data.message);
  return data.data;
};

// Fetch event flayers
const fetchEvents = async (limit = 10, upcoming = false) => {
  const url = new URL('http://localhost:8080/api/v1/events/flayers');
  url.searchParams.set('limit', limit);
  if (upcoming) url.searchParams.set('upcoming', 'true');

  const response = await fetch(url.toString());
  const data = await response.json();
  return data.data;
};

// Usage in component - News
function HomePage() {
  const [news, setNews] = useState([]);

  useEffect(() => {
    fetchFeaturedNews().then(setNews);
  }, []);

  return (
    <div>
      {news.map(article => (
        <div key={article.id}>
          <h2>{article.title}</h2>
          <p>{article.excerpt}</p>
        </div>
      ))}
    </div>
  );
}

// Usage in component - Pengurus
function PengurusPage() {
  const [pengurus, setPengurus] = useState({ periode: '', pengurus: [] });

  useEffect(() => {
    fetchPengurus('2024-2029', 'pimpinan_utama').then(setPengurus);
  }, []);

  return (
    <div>
      <h1>Pengurus Periode {pengurus.periode}</h1>
      {pengurus.pengurus.map(p => (
        <div key={p.id}>
          <img src={p.foto} alt={p.nama} />
          <h2>{p.nama}</h2>
          <p>{p.jabatan}</p>
          <p>{p.bio}</p>
        </div>
      ))}
    </div>
  );
}

// Usage in component - Contact Form
function ContactForm() {
  const [formData, setFormData] = useState({
    name: '', email: '', phone: '', subject: '', message: ''
  });
  const [status, setStatus] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);

    try {
      const result = await submitContact(formData);
      setStatus(`Pesan terkirim! Ticket ID: ${result.ticket_id}`);
      setFormData({ name: '', email: '', phone: '', subject: '', message: '' });
    } catch (error) {
      setStatus(`Error: ${error.message}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
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
        value={formData.phone}
        onChange={(e) => setFormData({...formData, phone: e.target.value})}
        placeholder="Telepon (opsional)"
      />
      <input
        value={formData.subject}
        onChange={(e) => setFormData({...formData, subject: e.target.value})}
        placeholder="Subjek"
        required
      />
      <textarea
        value={formData.message}
        onChange={(e) => setFormData({...formData, message: e.target.value})}
        placeholder="Pesan (minimal 10 karakter)"
        required
      />
      <button type="submit" disabled={loading}>
        {loading ? 'Mengirim...' : 'Kirim Pesan'}
      </button>
      {status && <p>{status}</p>}
    </form>
  );
}

// Usage in component - Events
function EventsPage() {
  const [events, setEvents] = useState([]);

  useEffect(() => {
    fetchEvents(10, true).then(setEvents); // Get 10 upcoming events
  }, []);

  return (
    <div>
      <h1>Event & Kegiatan Mendatang</h1>
      {events.map(event => (
        <div key={event.id}>
          <img src={event.image} alt={event.title} />
          <h2>{event.title}</h2>
          <p>{event.description}</p>
          <p>📅 {event.event_date}</p>
          <p>📍 {event.event_location}</p>
          {event.registration_url && (
            <a href={event.registration_url}>Daftar Sekarang</a>
          )}
        </div>
      ))}
    </div>
  );
}
```

---

### Vue.js Example

```javascript
// composables/useNews.js
export const useNews = () => {
  const fetchFeaturedNews = async (limit = 3) => {
    const response = await fetch(`http://localhost:8080/api/v1/news/featured?limit=${limit}`);
    const data = await response.json();
    return data.data;
  };

  return { fetchFeaturedNews };
};

// Usage in component
<script setup>
import { ref, onMounted } from 'vue';
import { useNews } from '@/composables/useNews';

const { fetchFeaturedNews } = useNews();
const news = ref([]);

onMounted(async () => {
  news.value = await fetchFeaturedNews(3);
});
</script>

<template>
  <div v-for="article in news" :key="article.id">
    <h2>{{ article.title }}</h2>
    <p>{{ article.excerpt }}</p>
  </div>
</template>
```

---

## 📱 Response Status Codes

| Code | Status | Description |
|------|--------|-------------|
| 200 | OK | Request successful |
| 400 | Bad Request | Invalid parameters |
| 404 | Not Found | Resource not found |
| 429 | Too Many Requests | Rate limit exceeded |
| 500 | Internal Server Error | Server error |

---

## 🔗 Quick Reference

**All endpoints return real data after seeding database!**

| Endpoint | Method | Real Data |
|----------|--------|-----------|
| `/health` | GET | ✅ Always works |
| `/api/v1/news` | GET | ✅ 7 articles |
| `/api/v1/news/featured` | GET | ✅ 3 featured |
| `/api/v1/opinions` | GET | ✅ 3 articles |
| `/api/v1/documents` | GET | ✅ 10 documents |
| `/api/v1/hero-slides` | GET | ✅ 3 slides |
| `/api/v1/organization/structure` | GET | ✅ 5 members |
| `/api/v1/pages/:slug` | GET | ✅ 4 pages |
| `/api/v1/categories` | GET | ✅ 9 categories |
| `/api/v1/settings` | GET | ✅ Full settings |
| `/api/v1/organization/pengurus` | GET | ✅ 12 pengurus |
| `/api/v1/editorial/team` | GET | ✅ 14 members |
| `/api/v1/contact/submit` | POST | ✅ Working |
| `/api/v1/events/flayers` | GET | ✅ 3 events |

---

## ⚠️ Important Notes

### Contact Form Rate Limiting
- **Limit:** 5 requests per hour per IP address
- **Error Code:** 429 Too Many Requests
- **Best Practice:** Validate on client-side before submission

### Event Flayers Auto-Filtering
- Events automatically filtered by `start_display_date` and `end_display_date`
- Use `upcoming=true` to get only future events
- Events ordered by `order_number` ASC, then `event_date` DESC

### Pengurus Filtering
- Default: Returns active pengurus only
- Use `active=false` to include inactive members
- Use `kategori` to filter by type (pimpinan_utama, sekretariat, bendahara, bidang)

### Editorial Team Structure
- Response grouped by role type for easy frontend rendering
- Includes complete contact information
- All members are active by default

---

**Last Updated:** 2025-01-14
**Version:** 1.1.0
**Status:** ✅ Ready to Use

🚀 **Start building your frontend now with real data!**

💡 **New APIs Added:** Pengurus, Editorial Team, Contact Form, Event Flayers
