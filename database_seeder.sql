-- ============================================
-- LP MA'ARIF NU DATABASE SEEDER
-- Version: 1.0.0
-- Date: 2025-01-11
-- Purpose: Insert dummy data for testing
-- ============================================

USE lp_maarif_nu;

-- Disable foreign key checks temporarily
SET FOREIGN_KEY_CHECKS = 0;

-- ============================================
-- 1. USERS (Already has 1 admin from schema)
-- ============================================

INSERT INTO users (name, email, password, role, avatar, phone, is_active, email_verified_at) VALUES
('Editor Berita', 'editor@lpmaarifnu.or.id', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'editor', 'https://ui-avatars.com/api/?name=Editor+Berita', '081234567890', TRUE, NOW()),
('Admin Dokumen', 'dokumen@lpmaarifnu.or.id', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'admin', 'https://ui-avatars.com/api/?name=Admin+Dokumen', '081234567891', TRUE, NOW()),
('Viewer User', 'viewer@lpmaarifnu.or.id', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'viewer', 'https://ui-avatars.com/api/?name=Viewer+User', '081234567892', TRUE, NOW());

-- ============================================
-- 2. TAGS
-- ============================================

INSERT INTO tags (name, slug) VALUES
('Pendidikan', 'pendidikan'),
('Beasiswa', 'beasiswa'),
('Kurikulum', 'kurikulum'),
('Pelatihan', 'pelatihan'),
('Seminar', 'seminar'),
('Workshop', 'workshop'),
('Kegiatan', 'kegiatan'),
('Prestasi', 'prestasi'),
('Teknologi', 'teknologi'),
('Karakter', 'karakter'),
('Literasi', 'literasi'),
('Inovasi', 'inovasi'),
('Kerjasama', 'kerjasama'),
('Pengabdian', 'pengabdian'),
('Penelitian', 'penelitian');

-- ============================================
-- 3. NEWS ARTICLES
-- ============================================

-- Featured News
INSERT INTO news_articles (title, slug, excerpt, content, image, category_id, author_id, status, published_at, views, is_featured, meta_title, meta_description, meta_keywords) VALUES
(
    'Peluncuran Program Beasiswa Unggulan 2024 untuk Siswa Berprestasi',
    'peluncuran-program-beasiswa-unggulan-2024',
    'LP Ma''arif NU meluncurkan program beasiswa unggulan tahun 2024 yang ditujukan untuk siswa berprestasi dari keluarga kurang mampu di seluruh Indonesia.',
    '<h2>Program Beasiswa Unggulan 2024</h2><p>LP Ma''arif NU dengan bangga mengumumkan peluncuran Program Beasiswa Unggulan 2024 yang merupakan komitmen kami dalam meningkatkan akses pendidikan berkualitas bagi siswa berprestasi dari berbagai latar belakang ekonomi.</p><h3>Latar Belakang</h3><p>Program ini diluncurkan sebagai respons terhadap kebutuhan akan dukungan pendidikan yang lebih luas dan merata di Indonesia. Dengan mempertimbangkan kondisi ekonomi yang masih menantang, LP Ma''arif NU berupaya memberikan kesempatan kepada siswa-siswa berbakat untuk melanjutkan pendidikan mereka tanpa terkendala masalah finansial.</p><h3>Kriteria Penerima</h3><ul><li>Siswa aktif jenjang SMP/MTs dan SMA/MA</li><li>Memiliki prestasi akademik minimal rata-rata 8.0</li><li>Berasal dari keluarga kurang mampu (dibuktikan dengan SKTM)</li><li>Aktif dalam kegiatan ekstrakurikuler</li><li>Memiliki sikap dan karakter yang baik</li></ul><h3>Fasilitas yang Diberikan</h3><ol><li>Biaya pendidikan penuh selama 1 tahun ajaran</li><li>Uang saku bulanan</li><li>Bantuan buku dan alat tulis</li><li>Program mentoring dan pembinaan karakter</li><li>Kesempatan mengikuti pelatihan kepemimpinan</li></ol><h3>Cara Pendaftaran</h3><p>Pendaftaran dapat dilakukan secara online melalui website resmi LP Ma''arif NU mulai tanggal 1 Februari 2024 hingga 30 Maret 2024. Berkas yang diperlukan meliputi:</p><ul><li>Formulir pendaftaran</li><li>Fotokopi rapor 2 semester terakhir</li><li>Surat Keterangan Tidak Mampu (SKTM)</li><li>Surat rekomendasi dari sekolah</li><li>Essay motivasi</li><li>Foto terbaru</li></ul><p>Untuk informasi lebih lanjut, silakan hubungi sekretariat LP Ma''arif NU atau kunjungi website resmi kami.</p>',
    'https://images.unsplash.com/photo-1523050854058-8df90110c9f1?w=800',
    1, 1, 'published', NOW() - INTERVAL 2 DAY, 1520, TRUE,
    'Program Beasiswa Unggulan 2024 LP Ma''arif NU',
    'LP Ma''arif NU meluncurkan program beasiswa unggulan 2024 untuk siswa berprestasi. Daftar sekarang!',
    'beasiswa, pendidikan, ma''arif nu, siswa berprestasi, 2024'
),
(
    'Implementasi Kurikulum Merdeka di Sekolah Ma''arif Se-Indonesia',
    'implementasi-kurikulum-merdeka-sekolah-maarif',
    'Lebih dari 5000 sekolah Ma''arif di seluruh Indonesia telah berhasil mengimplementasikan Kurikulum Merdeka dengan pendampingan intensif dari LP Ma''arif NU.',
    '<h2>Kurikulum Merdeka di Sekolah Ma''arif</h2><p>Dalam upaya meningkatkan kualitas pendidikan di lingkungan Ma''arif, LP Ma''arif NU telah melakukan pendampingan implementasi Kurikulum Merdeka kepada lebih dari 5000 sekolah di seluruh Indonesia.</p><h3>Tahapan Implementasi</h3><p>Proses implementasi dilakukan secara bertahap dengan mempertimbangkan kesiapan setiap sekolah:</p><ol><li><strong>Fase Persiapan:</strong> Sosialisasi dan pelatihan guru</li><li><strong>Fase Pilot Project:</strong> Uji coba di sekolah percontohan</li><li><strong>Fase Implementasi:</strong> Pelaksanaan di seluruh sekolah</li><li><strong>Fase Evaluasi:</strong> Monitoring dan evaluasi berkelanjutan</li></ol><h3>Program Pendampingan</h3><p>LP Ma''arif NU menyediakan program pendampingan komprehensif yang meliputi:</p><ul><li>Pelatihan guru tentang Kurikulum Merdeka</li><li>Workshop pengembangan modul ajar</li><li>Asesmen pembelajaran berbasis proyek</li><li>Pendampingan implementasi di kelas</li><li>Sharing session antar sekolah</li></ul><h3>Hasil dan Dampak</h3><p>Setelah satu semester implementasi, terlihat dampak positif yang signifikan:</p><ul><li>Peningkatan motivasi belajar siswa sebesar 40%</li><li>Guru lebih kreatif dalam mengembangkan pembelajaran</li><li>Siswa lebih aktif dan mandiri dalam belajar</li><li>Pembelajaran lebih kontekstual dan bermakna</li></ul>',
    'https://images.unsplash.com/photo-1503676260728-1c00da094a0b?w=800',
    1, 1, 'published', NOW() - INTERVAL 5 DAY, 2340, TRUE,
    'Implementasi Kurikulum Merdeka Sekolah Ma''arif',
    'Lebih dari 5000 sekolah Ma''arif telah sukses implementasi Kurikulum Merdeka dengan pendampingan LP Ma''arif NU.',
    'kurikulum merdeka, sekolah maarif, pendidikan, implementasi'
),
(
    'Pelatihan Guru Digital untuk Meningkatkan Kualitas Pembelajaran',
    'pelatihan-guru-digital-kualitas-pembelajaran',
    'LP Ma''arif NU menyelenggarakan pelatihan guru digital yang diikuti oleh 2000 guru dari berbagai daerah untuk meningkatkan kompetensi dalam pembelajaran berbasis teknologi.',
    '<h2>Pelatihan Guru Digital</h2><p>Dalam era digital ini, kompetensi guru dalam menggunakan teknologi pembelajaran menjadi sangat penting. LP Ma''arif NU menyelenggarakan program pelatihan guru digital yang komprehensif.</p><h3>Materi Pelatihan</h3><ul><li>Penggunaan Learning Management System (LMS)</li><li>Pembuatan konten pembelajaran digital</li><li>Penggunaan aplikasi pembelajaran interaktif</li><li>Asesmen online dan feedback digital</li><li>Keamanan digital dan etika online</li></ul><h3>Metodologi</h3><p>Pelatihan dilakukan dengan pendekatan blended learning:</p><ul><li>Sesi online melalui platform Zoom</li><li>Praktek langsung dengan pendampingan</li><li>Diskusi dan sharing session</li><li>Tugas proyek pembuatan konten</li><li>Presentasi hasil karya</li></ul>',
    'https://images.unsplash.com/photo-1516321318423-f06f85e504b3?w=800',
    2, 2, 'published', NOW() - INTERVAL 7 DAY, 1890, TRUE,
    'Pelatihan Guru Digital LP Ma''arif NU',
    'Program pelatihan guru digital untuk meningkatkan kompetensi pembelajaran berbasis teknologi di era digital.',
    'pelatihan guru, digital, teknologi, pembelajaran, ma''arif'
);

-- Regular News
INSERT INTO news_articles (title, slug, excerpt, content, image, category_id, author_id, status, published_at, views, is_featured, meta_title, meta_description, meta_keywords) VALUES
(
    'Workshop Pengembangan Karakter Siswa Melalui Pendidikan Kepramukaan',
    'workshop-pengembangan-karakter-pramuka',
    'Workshop ini bertujuan mengintegrasikan nilai-nilai kepramukaan dalam pembentukan karakter siswa di lingkungan sekolah Ma''arif.',
    '<h2>Workshop Karakter melalui Pramuka</h2><p>Pendidikan karakter merupakan aspek penting dalam pembentukan generasi muda yang berkualitas. Workshop ini dirancang untuk mengintegrasikan nilai-nilai kepramukaan dalam pendidikan karakter di sekolah.</p><h3>Nilai-Nilai Pramuka</h3><ul><li>Kedisiplinan dan tanggung jawab</li><li>Kerjasama dan kepemimpinan</li><li>Kemandirian dan kreativitas</li><li>Kepedulian terhadap lingkungan</li></ul>',
    'https://images.unsplash.com/photo-1529390079861-591de354faf5?w=800',
    3, 2, 'published', NOW() - INTERVAL 10 DAY, 890, FALSE,
    'Workshop Karakter Siswa Melalui Pramuka',
    'Workshop pengembangan karakter siswa melalui pendidikan kepramukaan di sekolah Ma''arif.',
    'pramuka, karakter, pendidikan, workshop, siswa'
),
(
    'Pelaksanaan Ujian Nasional Sekolah Ma''arif Berjalan Lancar',
    'pelaksanaan-ujian-nasional-maarif-lancar',
    'Ujian Nasional di seluruh sekolah Ma''arif berjalan dengan lancar dan tertib mengikuti protokol kesehatan yang ketat.',
    '<h2>Ujian Nasional Sekolah Ma''arif</h2><p>Pelaksanaan Ujian Nasional di lingkungan sekolah Ma''arif se-Indonesia berjalan dengan lancar dan tertib. Seluruh sekolah telah menerapkan protokol yang ketat untuk memastikan kelancaran ujian.</p><h3>Persiapan</h3><p>Persiapan yang matang dilakukan meliputi:</p><ul><li>Koordinasi dengan dinas pendidikan</li><li>Pengecekan sistem dan jaringan</li><li>Briefing pengawas</li><li>Persiapan ruang ujian</li></ul>',
    'https://images.unsplash.com/photo-1434030216411-0b793f4b4173?w=800',
    4, 1, 'published', NOW() - INTERVAL 15 DAY, 1234, FALSE,
    'Ujian Nasional Sekolah Ma''arif',
    'Pelaksanaan Ujian Nasional di sekolah Ma''arif berjalan lancar dengan protokol ketat.',
    'ujian nasional, sekolah, ma''arif, pendidikan'
),
(
    'Seminar Nasional Pendidikan Islam Modern di Era Digital',
    'seminar-nasional-pendidikan-islam-modern',
    'Seminar yang dihadiri 500 peserta ini membahas tantangan dan peluang pendidikan Islam di era digital dengan narasumber kompeten.',
    '<h2>Seminar Pendidikan Islam Modern</h2><p>LP Ma''arif NU menyelenggarakan Seminar Nasional bertema "Pendidikan Islam Modern di Era Digital" yang dihadiri oleh 500 peserta dari berbagai daerah.</p><h3>Narasumber</h3><ul><li>Prof. Dr. KH. Ahmad Syafi''i, MA - Pakar Pendidikan Islam</li><li>Dr. Miftahul Huda, M.Pd - Ahli Teknologi Pendidikan</li><li>Dr. Siti Aisyah, M.Pd - Praktisi Pendidikan</li></ul>',
    'https://images.unsplash.com/photo-1540575467063-178a50c2df87?w=800',
    3, 2, 'published', NOW() - INTERVAL 20 DAY, 2100, FALSE,
    'Seminar Pendidikan Islam Modern',
    'Seminar Nasional membahas tantangan dan peluang pendidikan Islam di era digital.',
    'seminar, pendidikan islam, modern, digital, ma''arif'
),
(
    'Kerjasama Internasional LP Ma''arif NU dengan Universitas Malaysia',
    'kerjasama-internasional-universitas-malaysia',
    'Penandatanganan MoU antara LP Ma''arif NU dengan Universitas Islam Malaysia membuka peluang pertukaran pelajar dan dosen.',
    '<h2>Kerjasama Internasional</h2><p>LP Ma''arif NU menjalin kerjasama internasional dengan Universitas Islam Malaysia dalam bidang pendidikan dan penelitian.</p><h3>Program Kerjasama</h3><ul><li>Pertukaran pelajar</li><li>Pertukaran dosen</li><li>Penelitian bersama</li><li>Joint conference</li><li>Pengembangan kurikulum</li></ul>',
    'https://images.unsplash.com/photo-1521737711867-e3b97375f902?w=800',
    1, 1, 'published', NOW() - INTERVAL 25 DAY, 1567, FALSE,
    'Kerjasama LP Ma''arif NU dengan Malaysia',
    'MoU LP Ma''arif NU dengan Universitas Islam Malaysia untuk pertukaran pelajar dan dosen.',
    'kerjasama, internasional, malaysia, ma''arif, pendidikan'
);

-- ============================================
-- 4. NEWS TAGS (Many-to-Many Relationships)
-- ============================================

-- Article 1: Beasiswa
INSERT INTO news_tags (news_article_id, tag_id) VALUES
(1, 1), (1, 2), (1, 8);

-- Article 2: Kurikulum Merdeka
INSERT INTO news_tags (news_article_id, tag_id) VALUES
(2, 1), (2, 3), (2, 12);

-- Article 3: Pelatihan Guru
INSERT INTO news_tags (news_article_id, tag_id) VALUES
(3, 1), (3, 4), (3, 9);

-- Article 4: Pramuka
INSERT INTO news_tags (news_article_id, tag_id) VALUES
(4, 10), (4, 7), (4, 1);

-- Article 5: Ujian
INSERT INTO news_tags (news_article_id, tag_id) VALUES
(5, 1), (5, 7);

-- Article 6: Seminar
INSERT INTO news_tags (news_article_id, tag_id) VALUES
(6, 5), (6, 1), (6, 9);

-- Article 7: Kerjasama
INSERT INTO news_tags (news_article_id, tag_id) VALUES
(7, 13), (7, 1), (7, 15);

-- ============================================
-- 5. OPINION ARTICLES
-- ============================================

INSERT INTO opinion_articles (title, slug, excerpt, content, image, author_name, author_title, author_image, author_bio, status, published_at, views, is_featured, meta_title, meta_description, meta_keywords) VALUES
(
    'Pendidikan Karakter di Era Digital: Tantangan dan Solusi',
    'pendidikan-karakter-era-digital',
    'Pentingnya menanamkan nilai-nilai karakter kepada generasi muda di tengah gempuran teknologi digital yang semakin masif.',
    '<h2>Pendidikan Karakter di Era Digital</h2><p>Era digital membawa perubahan besar dalam cara kita mendidik generasi muda. Kemudahan akses informasi di satu sisi memberikan manfaat, namun di sisi lain juga membawa tantangan tersendiri dalam pembentukan karakter.</p><h3>Tantangan</h3><ul><li>Paparan konten negatif di media sosial</li><li>Kecanduan gadget</li><li>Berkurangnya interaksi sosial langsung</li><li>Instant gratification mindset</li></ul><h3>Solusi</h3><p>Untuk menghadapi tantangan tersebut, diperlukan strategi khusus:</p><ol><li>Pendidikan literasi digital</li><li>Pembentukan filter konten yang baik</li><li>Penanaman nilai-nilai agama dan budaya</li><li>Pendampingan intensif dari orangtua dan guru</li></ol>',
    'https://images.unsplash.com/photo-1509062522246-3755977927d7?w=800',
    'Prof. Dr. KH. Ahmad Syafi''i, MA',
    'Pakar Pendidikan Islam',
    'https://ui-avatars.com/api/?name=Ahmad+Syafii&background=4CAF50&color=fff&size=200',
    'Pakar pendidikan Islam dengan pengalaman lebih dari 25 tahun di bidang pendidikan. Aktif menulis dan memberikan ceramah tentang pendidikan karakter.',
    'published', NOW() - INTERVAL 3 DAY, 890, TRUE,
    'Pendidikan Karakter di Era Digital',
    'Pentingnya pendidikan karakter di era digital dengan solusi praktis menghadapi tantangan teknologi.',
    'pendidikan, karakter, digital, generasi muda'
),
(
    'Integrasi Nilai-Nilai Keislaman dalam Kurikulum Modern',
    'integrasi-nilai-keislaman-kurikulum',
    'Bagaimana mengintegrasikan nilai-nilai keislaman dalam kurikulum modern tanpa mengurangi kompetensi akademik siswa.',
    '<h2>Integrasi Nilai Keislaman</h2><p>Pendidikan Islam modern harus mampu mengintegrasikan nilai-nilai keislaman dengan perkembangan ilmu pengetahuan dan teknologi terkini.</p><h3>Prinsip Integrasi</h3><ul><li>Tidak dikotomi antara ilmu agama dan umum</li><li>Menggunakan pendekatan kontekstual</li><li>Menekankan praktik, bukan hanya teori</li><li>Mengembangkan critical thinking</li></ul>',
    'https://images.unsplash.com/photo-1456513080510-7bf3a84b82f8?w=800',
    'Dr. H. Muhammad Yusuf, M.Pd',
    'Dosen Pendidikan Islam',
    'https://ui-avatars.com/api/?name=Muhammad+Yusuf&background=2196F3&color=fff&size=200',
    'Dosen di bidang Pendidikan Islam dengan fokus penelitian pada integrasi nilai keislaman dalam pendidikan modern.',
    'published', NOW() - INTERVAL 8 DAY, 1245, FALSE,
    'Integrasi Nilai Keislaman dalam Kurikulum',
    'Strategi mengintegrasikan nilai-nilai keislaman dalam kurikulum pendidikan modern.',
    'islam, kurikulum, pendidikan, integrasi'
),
(
    'Peran Guru dalam Membentuk Generasi Literat dan Berakhlak',
    'peran-guru-membentuk-generasi-literat',
    'Guru memiliki peran sentral dalam membentuk generasi yang tidak hanya literat tetapi juga memiliki akhlak mulia.',
    '<h2>Peran Guru</h2><p>Guru adalah garda terdepan dalam proses pendidikan. Peran guru sangat strategis dalam membentuk karakter dan kompetensi siswa.</p><h3>Tugas Guru Modern</h3><ul><li>Fasilitator pembelajaran</li><li>Motivator siswa</li><li>Role model</li><li>Inovator pendidikan</li></ul>',
    'https://images.unsplash.com/photo-1524178232363-1fb2b075b655?w=800',
    'Dr. Hj. Siti Aisyah, M.Pd',
    'Praktisi Pendidikan',
    'https://ui-avatars.com/api/?name=Siti+Aisyah&background=E91E63&color=fff&size=200',
    'Praktisi pendidikan dengan pengalaman 20 tahun. Aktif dalam pengembangan kompetensi guru dan inovasi pembelajaran.',
    'published', NOW() - INTERVAL 12 DAY, 678, FALSE,
    'Peran Guru Membentuk Generasi Literat',
    'Peran strategis guru dalam membentuk generasi yang literat dan berakhlak mulia.',
    'guru, literasi, akhlak, pendidikan'
);

-- ============================================
-- 6. OPINION TAGS
-- ============================================

INSERT INTO opinion_tags (opinion_article_id, tag_id) VALUES
(1, 10), (1, 9), (1, 1),
(2, 1), (2, 3), (2, 12),
(3, 1), (3, 11), (3, 10);

-- ============================================
-- 7. DOCUMENTS
-- ============================================

INSERT INTO documents (title, description, category_id, file_name, file_path, file_type, file_size, mime_type, download_count, is_public, uploaded_by, status) VALUES
(
    'Pedoman Penyelenggaraan Pendidikan Ma''arif NU 2024',
    'Pedoman lengkap penyelenggaraan pendidikan di lingkungan Ma''arif NU yang mencakup standar operasional, kurikulum, dan tata kelola sekolah.',
    5, 'pedoman-pendidikan-maarif-2024.pdf',
    'https://cdn.lpmaarifnu.or.id/documents/pedoman-pendidikan-maarif-2024.pdf',
    'PDF', 2621440, 'application/pdf', 1534, TRUE, 1, 'active'
),
(
    'Kurikulum Integratif Ma''arif NU Jenjang SMP/MTs',
    'Dokumen kurikulum integratif yang mengintegrasikan nilai-nilai keislaman dengan kurikulum nasional untuk jenjang SMP/MTs.',
    6, 'kurikulum-integratif-smp-mts.pdf',
    'https://cdn.lpmaarifnu.or.id/documents/kurikulum-integratif-smp-mts.pdf',
    'PDF', 3145728, 'application/pdf', 2341, TRUE, 1, 'active'
),
(
    'Kurikulum Integratif Ma''arif NU Jenjang SMA/MA',
    'Dokumen kurikulum integratif untuk jenjang SMA/MA dengan pendekatan saintifik dan nilai-nilai keislaman.',
    6, 'kurikulum-integratif-sma-ma.pdf',
    'https://cdn.lpmaarifnu.or.id/documents/kurikulum-integratif-sma-ma.pdf',
    'PDF', 3670016, 'application/pdf', 1987, TRUE, 1, 'active'
),
(
    'Peraturan Standar Akreditasi Sekolah Ma''arif',
    'Peraturan dan standar akreditasi yang harus dipenuhi oleh sekolah-sekolah di bawah naungan LP Ma''arif NU.',
    7, 'standar-akreditasi-sekolah.pdf',
    'https://cdn.lpmaarifnu.or.id/documents/standar-akreditasi-sekolah.pdf',
    'PDF', 1572864, 'application/pdf', 876, TRUE, 3, 'active'
),
(
    'Panduan Implementasi Kurikulum Merdeka di Sekolah Ma''arif',
    'Panduan praktis implementasi Kurikulum Merdeka dengan tetap mempertahankan nilai-nilai Ma''arif NU.',
    8, 'panduan-kurikulum-merdeka.pdf',
    'https://cdn.lpmaarifnu.or.id/documents/panduan-kurikulum-merdeka.pdf',
    'PDF', 2097152, 'application/pdf', 3456, TRUE, 2, 'active'
),
(
    'Panduan Teknis Penilaian Pembelajaran Berbasis Proyek',
    'Panduan teknis untuk guru dalam melakukan penilaian pembelajaran berbasis proyek sesuai Kurikulum Merdeka.',
    8, 'panduan-penilaian-pbp.pdf',
    'https://cdn.lpmaarifnu.or.id/documents/panduan-penilaian-pbp.pdf',
    'PDF', 1048576, 'application/pdf', 1234, TRUE, 2, 'active'
),
(
    'Formulir Pendaftaran Beasiswa Unggulan 2024',
    'Formulir pendaftaran untuk program beasiswa unggulan LP Ma''arif NU tahun 2024.',
    9, 'formulir-beasiswa-2024.pdf',
    'https://cdn.lpmaarifnu.or.id/documents/formulir-beasiswa-2024.pdf',
    'PDF', 524288, 'application/pdf', 4567, TRUE, 1, 'active'
),
(
    'Formulir Usulan Kegiatan Sekolah',
    'Template formulir untuk mengajukan usulan kegiatan sekolah kepada LP Ma''arif NU.',
    9, 'formulir-usulan-kegiatan.docx',
    'https://cdn.lpmaarifnu.or.id/documents/formulir-usulan-kegiatan.docx',
    'DOCX', 204800, 'application/vnd.openxmlformats-officedocument.wordprocessingml.document', 789, TRUE, 3, 'active'
),
(
    'Juknis Penyelenggaraan Ujian Akhir Sekolah',
    'Petunjuk teknis penyelenggaraan ujian akhir sekolah di lingkungan Ma''arif NU.',
    8, 'juknis-ujian-akhir.pdf',
    'https://cdn.lpmaarifnu.or.id/documents/juknis-ujian-akhir.pdf',
    'PDF', 1310720, 'application/pdf', 2134, TRUE, 2, 'active'
),
(
    'Modul Pelatihan Guru Digital',
    'Modul lengkap untuk pelatihan kompetensi digital bagi guru-guru Ma''arif.',
    8, 'modul-pelatihan-guru-digital.pdf',
    'https://cdn.lpmaarifnu.or.id/documents/modul-pelatihan-guru-digital.pdf',
    'PDF', 5242880, 'application/pdf', 3890, TRUE, 1, 'active'
);

-- ============================================
-- 8. HERO SLIDES
-- ============================================

INSERT INTO hero_slides (title, description, image, cta_label, cta_href, cta_secondary_label, cta_secondary_href, order_number, is_active, start_date, end_date) VALUES
(
    'Membangun Pendidikan Islam Berkualitas dan Berkarakter',
    'LP Ma''arif NU berkomitmen mengembangkan pendidikan Islam yang berkualitas dengan mengintegrasikan nilai-nilai keislaman, kearifan lokal, dan kemajuan teknologi untuk membentuk generasi yang berakhlak mulia dan berdaya saing tinggi.',
    'https://images.unsplash.com/photo-1427504494785-3a9ca7044f45?w=1920',
    'Pelajari Lebih Lanjut', '/tentang/visi-misi',
    'Hubungi Kami', '/kontak',
    1, TRUE, NOW() - INTERVAL 30 DAY, NULL
),
(
    'Program Beasiswa Unggulan 2024',
    'Raih kesempatan emas untuk mendapatkan beasiswa pendidikan penuh! Daftar sekarang dan wujudkan impian pendidikanmu bersama LP Ma''arif NU.',
    'https://images.unsplash.com/photo-1523050854058-8df90110c9f1?w=1920',
    'Daftar Sekarang', '/beasiswa',
    'Info Lengkap', '/berita/peluncuran-program-beasiswa-unggulan-2024',
    2, TRUE, NOW() - INTERVAL 20 DAY, NOW() + INTERVAL 60 DAY
),
(
    'Digitalisasi Pendidikan Ma''arif',
    'Bergabunglah dalam transformasi digital pendidikan! Pelatihan guru digital, pembelajaran online, dan inovasi teknologi pendidikan untuk masa depan yang lebih baik.',
    'https://images.unsplash.com/photo-1501504905252-473c47e087f8?w=1920',
    'Ikuti Program', '/program/digital',
    'Lihat Dokumentasi', '/dokumen',
    3, TRUE, NOW() - INTERVAL 10 DAY, NULL
);

-- ============================================
-- 9. ORGANIZATION POSITIONS (Already inserted in schema, add more)
-- ============================================

INSERT INTO organization_positions (position_name, position_level, position_type, parent_id, order_number) VALUES
('Wakil Ketua III', 2, 'wakil', NULL, 4),
('Wakil Sekretaris', 3, 'sekretaris', 4, 5),
('Wakil Bendahara', 4, 'bendahara', 5, 6);

-- ============================================
-- 10. BOARD MEMBERS
-- ============================================

INSERT INTO board_members (position_id, name, title, photo, bio, email, phone, social_media, period_start, period_end, is_active, order_number) VALUES
(
    1, 'Prof. Dr. KH. Said Aqil Siradj, MA',
    'Prof. Dr. KH.',
    'https://ui-avatars.com/api/?name=Said+Aqil+Siradj&background=1976D2&color=fff&size=400',
    'Ketua Umum LP Ma''arif NU periode 2024-2029. Beliau adalah ulama dan akademisi terkemuka yang memiliki dedikasi tinggi dalam pengembangan pendidikan Islam di Indonesia.',
    'ketua@lpmaarifnu.or.id', '021-12345678',
    '{"facebook":"https://facebook.com/saidaqilsiradj","twitter":"https://twitter.com/saidaqil","instagram":"https://instagram.com/saidaqilsiradj"}',
    2024, 2029, TRUE, 1
),
(
    2, 'Dr. H. Ahmad Lutfi, M.Pd',
    'Dr. H.',
    'https://ui-avatars.com/api/?name=Ahmad+Lutfi&background=388E3C&color=fff&size=400',
    'Wakil Ketua I LP Ma''arif NU yang membidangi Pendidikan Dasar dan Menengah. Berpengalaman dalam manajemen pendidikan lebih dari 20 tahun.',
    'wakilketua1@lpmaarifnu.or.id', '021-12345679',
    '{"linkedin":"https://linkedin.com/in/ahmadlutfi"}',
    2024, 2029, TRUE, 2
),
(
    3, 'Dr. H. Abdurrahman Wahid, M.Ag',
    'Dr. H.',
    'https://ui-avatars.com/api/?name=Abdurrahman+Wahid&background=F57C00&color=fff&size=400',
    'Wakil Ketua II yang fokus pada pengembangan kurikulum dan peningkatan kualitas pembelajaran di sekolah-sekolah Ma''arif.',
    'wakilketua2@lpmaarifnu.or.id', '021-12345680',
    '{}',
    2024, 2029, TRUE, 3
),
(
    4, 'Dr. Hj. Siti Aisyah, M.Pd',
    'Dr. Hj.',
    'https://ui-avatars.com/api/?name=Siti+Aisyah&background=C62828&color=fff&size=400',
    'Sekretaris Umum LP Ma''arif NU yang bertanggung jawab atas administrasi dan koordinasi kegiatan organisasi.',
    'sekretaris@lpmaarifnu.or.id', '021-12345681',
    '{"instagram":"https://instagram.com/sitiaisyah"}',
    2024, 2029, TRUE, 4
),
(
    5, 'H. Abdul Rahman, SE, M.Ak',
    'H.',
    'https://ui-avatars.com/api/?name=Abdul+Rahman&background=7B1FA2&color=fff&size=400',
    'Bendahara Umum yang mengelola keuangan dan aset LP Ma''arif NU dengan prinsip transparansi dan akuntabilitas.',
    'bendahara@lpmaarifnu.or.id', '021-12345682',
    '{}',
    2024, 2029, TRUE, 5
);

-- ============================================
-- 11. PAGES (Already inserted in schema, update with content)
-- ============================================

UPDATE pages SET
    content = NULL,
    metadata = JSON_OBJECT(
        'visi', 'Menjadi lembaga pendidikan Islam yang unggul, modern, dan berkarakter untuk mewujudkan generasi yang berakhlak mulia, cerdas, dan berdaya saing global.',
        'misi', JSON_ARRAY(
            'Menyelenggarakan pendidikan berkualitas yang berbasis nilai-nilai keislaman',
            'Mengembangkan kurikulum yang integratif antara ilmu agama dan ilmu umum',
            'Memberdayakan sumber daya manusia yang profesional dan berakhlak mulia',
            'Membangun kerjasama dengan berbagai pihak untuk kemajuan pendidikan',
            'Mengimplementasikan teknologi dalam pembelajaran'
        ),
        'nilai_nilai', JSON_ARRAY(
            JSON_OBJECT('icon', 'integrity', 'title', 'Integritas', 'description', 'Menjunjung tinggi kejujuran dan tanggung jawab dalam setiap aspek'),
            JSON_OBJECT('icon', 'professionalism', 'title', 'Profesionalisme', 'description', 'Menjalankan tugas dengan kompeten dan penuh dedikasi'),
            JSON_OBJECT('icon', 'innovation', 'title', 'Inovasi', 'description', 'Senantiasa berinovasi untuk kemajuan pendidikan'),
            JSON_OBJECT('icon', 'collaboration', 'title', 'Kolaborasi', 'description', 'Membangun kerjasama yang sinergis dengan berbagai pihak')
        )
    ),
    is_active = TRUE
WHERE slug = 'visi-misi';

UPDATE pages SET
    content = NULL,
    metadata = JSON_OBJECT(
        'introduction', '<p>LP Ma''arif NU merupakan lembaga pendidikan di bawah naungan Nahdlatul Ulama yang telah berkiprah dalam dunia pendidikan Islam di Indonesia sejak tahun 1916.</p>',
        'timeline', JSON_ARRAY(
            JSON_OBJECT('year', '1916', 'title', 'Berdirinya Nahdlatul Ulama', 'description', 'Nahdlatul Ulama (NU) didirikan di Surabaya oleh KH. Hasyim Asy''ari sebagai organisasi kemasyarakatan Islam.'),
            JSON_OBJECT('year', '1926', 'title', 'Pembentukan Lembaga Pendidikan', 'description', 'NU mulai fokus pada pengembangan lembaga pendidikan Islam di berbagai daerah.'),
            JSON_OBJECT('year', '1960', 'title', 'Formalisasi LP Ma''arif', 'description', 'Pembentukan formal LP Ma''arif NU sebagai lembaga yang menaungi pendidikan di lingkungan NU.'),
            JSON_OBJECT('year', '1990', 'title', 'Ekspansi Nasional', 'description', 'Perluasan jangkauan LP Ma''arif ke seluruh Indonesia dengan ribuan sekolah di bawah naungannya.'),
            JSON_OBJECT('year', '2010', 'title', 'Modernisasi Pendidikan', 'description', 'Implementasi teknologi dan kurikulum modern dengan tetap mempertahankan nilai-nilai keislaman.'),
            JSON_OBJECT('year', '2024', 'title', 'Transformasi Digital', 'description', 'Akselerasi digitalisasi pendidikan dan peningkatan kualitas pembelajaran di era digital.')
        )
    ),
    is_active = TRUE
WHERE slug = 'sejarah';

UPDATE pages SET
    content = NULL,
    metadata = JSON_OBJECT(
        'programs', JSON_ARRAY(
            JSON_OBJECT(
                'title', 'Peningkatan Kualitas Pembelajaran',
                'description', 'Program pelatihan dan pendampingan guru untuk meningkatkan kualitas pembelajaran di kelas.',
                'icon', 'teaching'
            ),
            JSON_OBJECT(
                'title', 'Digitalisasi Pendidikan',
                'description', 'Implementasi teknologi dalam pembelajaran untuk meningkatkan efektivitas dan efisiensi.',
                'icon', 'digital'
            ),
            JSON_OBJECT(
                'title', 'Pengembangan Kurikulum Integratif',
                'description', 'Pengembangan kurikulum yang mengintegrasikan nilai-nilai keislaman dengan ilmu pengetahuan modern.',
                'icon', 'curriculum'
            ),
            JSON_OBJECT(
                'title', 'Beasiswa Pendidikan',
                'description', 'Program beasiswa untuk siswa berprestasi dari keluarga kurang mampu.',
                'icon', 'scholarship'
            )
        )
    ),
    is_active = TRUE
WHERE slug = 'program-strategis';

UPDATE pages SET
    content = NULL,
    metadata = JSON_OBJECT(
        'introduction', '<p>Gerakan Pramuka Ma''arif merupakan bagian integral dari pendidikan karakter di sekolah-sekolah Ma''arif yang bertujuan membentuk karakter siswa yang mandiri, bertanggung jawab, dan peduli terhadap sesama.</p>',
        'activities', JSON_ARRAY(
            'Perkemahan Wajib (Persami)',
            'Jambore Pramuka Ma''arif',
            'Bakti Sosial',
            'Pelatihan Kepemimpinan',
            'Gladian Pemimpin Satuan (Dianpinsat)'
        ),
        'values', JSON_ARRAY(
            'Kedisiplinan dan tanggung jawab',
            'Kepemimpinan dan kerjasama',
            'Kemandirian dan kreativitas',
            'Kepedulian sosial dan lingkungan'
        )
    ),
    is_active = TRUE
WHERE slug = 'pramuka';

-- ============================================
-- 12. SETTINGS (Already inserted in schema, update values)
-- ============================================

UPDATE settings SET setting_value = 'LP Ma''arif NU' WHERE setting_key = 'site_name';
UPDATE settings SET setting_value = 'Lembaga Pendidikan Ma''arif Nahdlatul Ulama - Membangun Pendidikan Islam Berkualitas' WHERE setting_key = 'site_description';
UPDATE settings SET setting_value = 'info@lpmaarifnu.or.id' WHERE setting_key = 'contact_email';
UPDATE settings SET setting_value = '(021) 3920677' WHERE setting_key = 'contact_phone';
UPDATE settings SET setting_value = 'Jl. Kramat Raya No. 164, Jakarta Pusat 10430' WHERE setting_key = 'contact_address';
UPDATE settings SET setting_value = 'https://facebook.com/lpmaarifnu' WHERE setting_key = 'social_facebook';
UPDATE settings SET setting_value = 'https://twitter.com/lpmaarifnu' WHERE setting_key = 'social_twitter';
UPDATE settings SET setting_value = 'https://instagram.com/lpmaarifnu' WHERE setting_key = 'social_instagram';
UPDATE settings SET setting_value = 'https://youtube.com/@lpmaarifnu' WHERE setting_key = 'social_youtube';

-- ============================================
-- 13. MEDIA LIBRARY
-- ============================================

INSERT INTO media (file_name, original_name, file_path, file_url, file_type, mime_type, file_size, width, height, folder, alt_text, caption, uploaded_by) VALUES
('hero-education.jpg', 'hero-education.jpg', '/media/images/hero-education.jpg', 'https://cdn.lpmaarifnu.or.id/images/hero-education.jpg', 'image', 'image/jpeg', 2048000, 1920, 1080, 'hero', 'Pendidikan Berkualitas', 'Siswa sedang belajar di kelas', 1),
('beasiswa-2024.jpg', 'beasiswa-2024.jpg', '/media/images/beasiswa-2024.jpg', 'https://cdn.lpmaarifnu.or.id/images/beasiswa-2024.jpg', 'image', 'image/jpeg', 1536000, 1280, 720, 'news', 'Program Beasiswa 2024', 'Banner program beasiswa unggulan', 1),
('seminar-pendidikan.jpg', 'seminar-pendidikan.jpg', '/media/images/seminar-pendidikan.jpg', 'https://cdn.lpmaarifnu.or.id/images/seminar-pendidikan.jpg', 'image', 'image/jpeg', 1843200, 1600, 900, 'events', 'Seminar Pendidikan', 'Peserta seminar pendidikan Islam', 2);

-- ============================================
-- 14. PAGE VIEWS & DOWNLOAD LOGS
-- ============================================

-- Insert sample page views
INSERT INTO page_views (viewable_type, viewable_id, ip_address, user_agent, viewed_at) VALUES
('news_articles', 1, '192.168.1.100', 'Mozilla/5.0', NOW() - INTERVAL 1 DAY),
('news_articles', 1, '192.168.1.101', 'Mozilla/5.0', NOW() - INTERVAL 2 DAY),
('news_articles', 2, '192.168.1.102', 'Mozilla/5.0', NOW() - INTERVAL 1 DAY),
('opinion_articles', 1, '192.168.1.103', 'Mozilla/5.0', NOW() - INTERVAL 3 DAY);

-- Insert sample download logs
INSERT INTO download_logs (document_id, ip_address, user_agent, downloaded_at) VALUES
(1, '192.168.1.104', 'Mozilla/5.0', NOW() - INTERVAL 1 DAY),
(1, '192.168.1.105', 'Mozilla/5.0', NOW() - INTERVAL 2 DAY),
(2, '192.168.1.106', 'Mozilla/5.0', NOW() - INTERVAL 1 DAY),
(5, '192.168.1.107', 'Mozilla/5.0', NOW() - INTERVAL 3 DAY);

-- Re-enable foreign key checks
SET FOREIGN_KEY_CHECKS = 1;

-- ============================================
-- VERIFY DATA
-- ============================================

-- Show summary of inserted data
SELECT 'Data seeding completed successfully!' AS Status;

SELECT
    'Users' AS Table_Name, COUNT(*) AS Record_Count FROM users
UNION ALL
SELECT 'Tags', COUNT(*) FROM tags
UNION ALL
SELECT 'Categories', COUNT(*) FROM categories
UNION ALL
SELECT 'News Articles', COUNT(*) FROM news_articles
UNION ALL
SELECT 'Opinion Articles', COUNT(*) FROM opinion_articles
UNION ALL
SELECT 'Documents', COUNT(*) FROM documents
UNION ALL
SELECT 'Hero Slides', COUNT(*) FROM hero_slides
UNION ALL
SELECT 'Board Members', COUNT(*) FROM board_members
UNION ALL
SELECT 'Pages', COUNT(*) FROM pages
UNION ALL
SELECT 'Settings', COUNT(*) FROM settings
UNION ALL
SELECT 'Media', COUNT(*) FROM media;

-- ============================================
-- NOTES
-- ============================================

/*
Seeder ini berisi data dummy untuk testing dengan detail:

1. Users: 4 users (1 super_admin, 1 editor, 1 admin, 1 viewer)
2. Tags: 15 tags untuk kategorisasi artikel
3. News Articles: 7 artikel berita (3 featured, 4 regular)
4. Opinion Articles: 3 artikel opini
5. Documents: 10 dokumen berbagai kategori
6. Hero Slides: 3 slides untuk homepage
7. Board Members: 5 pengurus organisasi
8. Pages: 4 halaman statis (visi-misi, sejarah, program, pramuka)
9. Settings: Website settings lengkap
10. Media: 3 sample media files

Semua data menggunakan placeholder images dari Unsplash dan UI Avatars.
Password untuk semua user: "password" (sudah di-hash)

Untuk testing API:
- Categories sudah terisi dari schema (news: 4, document: 5)
- Relasi many-to-many sudah dibuat (news_tags, opinion_tags)
- Sample views dan downloads untuk testing analytics

Total records: 100+ records across all tables
*/
