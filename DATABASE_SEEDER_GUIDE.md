# Database Seeder Guide

Panduan lengkap untuk menggunakan database seeder dan populate data dummy untuk testing.

## 📦 File Seeder

**File:** `database_seeder.sql`

**Size:** ~40KB

**Purpose:** Insert data dummy untuk testing API

---

## 📊 Data Yang Di-Insert

### Summary

| Table | Records | Description |
|-------|---------|-------------|
| **users** | 4 | Admin, editor, dan viewer |
| **tags** | 15 | Tags untuk artikel |
| **categories** | 9 | Sudah ada dari schema (4 news, 5 document) |
| **news_articles** | 7 | 3 featured + 4 regular |
| **news_tags** | 14 | Relasi many-to-many |
| **opinion_articles** | 3 | Artikel opini |
| **opinion_tags** | 9 | Relasi many-to-many |
| **documents** | 10 | Berbagai kategori dokumen |
| **hero_slides** | 3 | Homepage slider |
| **organization_positions** | 8 | Sudah ada 5 dari schema + 3 baru |
| **board_members** | 5 | Pengurus organisasi |
| **pages** | 4 | Sudah ada dari schema, di-update |
| **settings** | 9 | Sudah ada dari schema, di-update |
| **media** | 3 | Sample media files |
| **page_views** | 4 | Sample analytics |
| **download_logs** | 4 | Sample downloads |
| **TOTAL** | **100+** | **Complete test data** |

---

## 🚀 Cara Menggunakan

### Option 1: Import Langsung (Recommended)

```bash
# Pastikan database sudah dibuat dan schema sudah di-import
mysql -u root -p lp_maarif_nu < database_seeder.sql
```

### Option 2: Import via MySQL Client

```bash
# Login ke MySQL
mysql -u root -p

# Pilih database
USE lp_maarif_nu;

# Import seeder
SOURCE /path/to/database_seeder.sql;
```

### Option 3: Via phpMyAdmin

1. Login ke phpMyAdmin
2. Pilih database `lp_maarif_nu`
3. Tab **Import**
4. Choose file: `database_seeder.sql`
5. Click **Go**

---

## 📋 Data Details

### 1. Users (4 records)

| Email | Password | Role | Status |
|-------|----------|------|--------|
| admin@lpmaarifnu.or.id | password | super_admin | ✅ Active |
| editor@lpmaarifnu.or.id | password | editor | ✅ Active |
| dokumen@lpmaarifnu.or.id | password | admin | ✅ Active |
| viewer@lpmaarifnu.or.id | password | viewer | ✅ Active |

**Note:** Password untuk semua user adalah `"password"` (sudah di-hash dengan bcrypt)

### 2. Tags (15 records)

Categories untuk artikel:
- Pendidikan, Beasiswa, Kurikulum
- Pelatihan, Seminar, Workshop
- Kegiatan, Prestasi, Teknologi
- Karakter, Literasi, Inovasi
- Kerjasama, Pengabdian, Penelitian

### 3. News Articles (7 records)

#### Featured News (3):
1. **Peluncuran Program Beasiswa Unggulan 2024**
   - Views: 1,520
   - Category: Nasional
   - Tags: Pendidikan, Beasiswa, Prestasi
   - Status: Published

2. **Implementasi Kurikulum Merdeka di Sekolah Ma'arif**
   - Views: 2,340
   - Category: Nasional
   - Tags: Pendidikan, Kurikulum, Inovasi
   - Status: Published

3. **Pelatihan Guru Digital untuk Meningkatkan Kualitas**
   - Views: 1,890
   - Category: Daerah
   - Tags: Pendidikan, Pelatihan, Teknologi
   - Status: Published

#### Regular News (4):
4. Workshop Pengembangan Karakter Siswa
5. Pelaksanaan Ujian Nasional Berjalan Lancar
6. Seminar Nasional Pendidikan Islam Modern
7. Kerjasama Internasional dengan Malaysia

### 4. Opinion Articles (3 records)

1. **Pendidikan Karakter di Era Digital**
   - Author: Prof. Dr. KH. Ahmad Syafi'i, MA
   - Views: 890
   - Featured: Yes

2. **Integrasi Nilai-Nilai Keislaman dalam Kurikulum**
   - Author: Dr. H. Muhammad Yusuf, M.Pd
   - Views: 1,245

3. **Peran Guru dalam Membentuk Generasi Literat**
   - Author: Dr. Hj. Siti Aisyah, M.Pd
   - Views: 678

### 5. Documents (10 records)

**By Category:**

**Pedoman (1):**
- Pedoman Penyelenggaraan Pendidikan Ma'arif NU 2024

**Kurikulum (2):**
- Kurikulum Integratif SMP/MTs
- Kurikulum Integratif SMA/MA

**Regulasi (1):**
- Peraturan Standar Akreditasi Sekolah

**Panduan (4):**
- Panduan Implementasi Kurikulum Merdeka
- Panduan Teknis Penilaian PBP
- Juknis Penyelenggaraan Ujian Akhir
- Modul Pelatihan Guru Digital

**Formulir (2):**
- Formulir Pendaftaran Beasiswa 2024
- Formulir Usulan Kegiatan Sekolah

### 6. Hero Slides (3 records)

1. **Membangun Pendidikan Islam Berkualitas**
   - CTA: "Pelajari Lebih Lanjut" → /tentang/visi-misi
   - Secondary: "Hubungi Kami" → /kontak

2. **Program Beasiswa Unggulan 2024**
   - CTA: "Daftar Sekarang" → /beasiswa
   - Secondary: "Info Lengkap" → /berita/...

3. **Digitalisasi Pendidikan Ma'arif**
   - CTA: "Ikuti Program" → /program/digital
   - Secondary: "Lihat Dokumentasi" → /dokumen

### 7. Board Members (5 records)

1. Prof. Dr. KH. Said Aqil Siradj, MA - Ketua Umum
2. Dr. H. Ahmad Lutfi, M.Pd - Wakil Ketua I
3. Dr. H. Abdurrahman Wahid, M.Ag - Wakil Ketua II
4. Dr. Hj. Siti Aisyah, M.Pd - Sekretaris Umum
5. H. Abdul Rahman, SE, M.Ak - Bendahara Umum

Period: 2024-2029

### 8. Pages (4 records)

1. **visi-misi** - Visi, Misi, dan Nilai-nilai LP Ma'arif NU
2. **sejarah** - Timeline sejarah dari 1916 - 2024
3. **program-strategis** - 4 program strategis utama
4. **pramuka** - Gerakan Pramuka Ma'arif

### 9. Settings (9 records)

All public settings updated:
- Site name, description
- Contact: email, phone, address
- Social media: Facebook, Twitter, Instagram, YouTube

---

## 🧪 Testing Setelah Import

### 1. Verify Import

```sql
-- Check record counts
SELECT 'Users' AS Table_Name, COUNT(*) AS Count FROM users
UNION ALL
SELECT 'News Articles', COUNT(*) FROM news_articles
UNION ALL
SELECT 'Documents', COUNT(*) FROM documents;
```

### 2. Test Data Access

```sql
-- Get featured news
SELECT * FROM news_articles WHERE is_featured = TRUE;

-- Get published articles
SELECT * FROM news_articles WHERE status = 'published';

-- Get documents by category
SELECT d.*, c.name as category_name
FROM documents d
JOIN categories c ON d.category_id = c.id
WHERE c.slug = 'pedoman';
```

### 3. Test API Endpoints

Start the API and test:

```bash
# Start API
go run cmd/api/main.go

# Test in another terminal
curl http://localhost:8080/api/v1/news
curl http://localhost:8080/api/v1/news/featured
curl http://localhost:8080/api/v1/categories
curl http://localhost:8080/api/v1/documents
curl http://localhost:8080/api/v1/hero-slides
```

---

## 🎯 Use Cases

### For Frontend Development

Data tersedia untuk testing:
- ✅ Homepage with hero slides
- ✅ News listing with pagination
- ✅ News detail with related articles
- ✅ Opinion articles
- ✅ Document library
- ✅ About pages (visi-misi, sejarah, etc.)
- ✅ Organization structure

### For API Testing

Test all endpoints dengan data real:
- ✅ Pagination working
- ✅ Search working
- ✅ Filter by category
- ✅ Sort by date/views
- ✅ Featured articles
- ✅ Related articles

### For Demo

Ready untuk demo dengan:
- ✅ Complete content
- ✅ Images (from Unsplash)
- ✅ Realistic data
- ✅ Proper relationships

---

## 🔄 Reset Data

Jika perlu reset dan re-seed:

```bash
# Drop dan re-create database
mysql -u root -p -e "DROP DATABASE lp_maarif_nu;"
mysql -u root -p -e "CREATE DATABASE lp_maarif_nu CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# Import schema
mysql -u root -p lp_maarif_nu < database_schema.sql

# Import seeder
mysql -u root -p lp_maarif_nu < database_seeder.sql
```

---

## 📝 Notes

### Images

All images use placeholder dari:
- **Unsplash** - High quality photos
- **UI Avatars** - Avatar generator

Format: `https://images.unsplash.com/photo-{id}?w=800`

### Passwords

All user passwords: `"password"` (bcrypt hash)

Hash: `$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi`

### Dates

- News published dates: Last 25 days
- Hero slides: Active now
- Board members: Period 2024-2029
- Page views: Last 3 days

### Content

- All content in Bahasa Indonesia
- Realistic and contextual
- HTML formatted untuk articles
- JSON formatted untuk pages

---

## ⚠️ Important

### Before Import

1. ✅ Database `lp_maarif_nu` harus sudah ada
2. ✅ Schema sudah di-import (`database_schema.sql`)
3. ✅ Foreign key checks akan di-disable temporarily
4. ✅ Backup database jika sudah ada data

### After Import

1. ✅ Verify record counts
2. ✅ Test API endpoints
3. ✅ Check relationships
4. ✅ Verify images loading

---

## 🆘 Troubleshooting

### Error: Database doesn't exist

**Solution:**
```sql
CREATE DATABASE lp_maarif_nu CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### Error: Table doesn't exist

**Solution:** Import schema first
```bash
mysql -u root -p lp_maarif_nu < database_schema.sql
```

### Error: Duplicate entry

**Solution:** Database sudah memiliki data. Reset atau skip duplicates.

### Some data missing

**Solution:** Check logs dan re-run import
```bash
mysql -u root -p lp_maarif_nu < database_seeder.sql > import.log 2>&1
```

---

## 📚 Related Files

- **database_schema.sql** - Database structure
- **database_seeder.sql** - This seeder file
- **API_DOCUMENTATION.md** - API endpoints
- **POSTMAN_GUIDE.md** - Testing with Postman

---

## ✅ Checklist

Before deploying to production:

- [ ] Remove or modify dummy data
- [ ] Update admin credentials
- [ ] Use real images
- [ ] Update contact information
- [ ] Review all content
- [ ] Update settings

---

**Created:** 2025-01-11
**Version:** 1.0.0
**Total Records:** 100+
**Status:** ✅ Ready to Use

Happy Testing! 🚀
