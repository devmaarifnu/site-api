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

// Usage in component
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

---

**Last Updated:** 2025-01-11
**Version:** 1.0.0
**Status:** ✅ Ready to Use

🚀 **Start building your frontend now with real data!**
