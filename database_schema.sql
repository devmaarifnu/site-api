-- ============================================
-- LP MA'ARIF NU DATABASE SCHEMA
-- Version: 1.0.0
-- Date: 2025-01-11
-- ============================================

-- Drop existing database if exists
DROP DATABASE IF EXISTS lp_maarif_nu;

-- Create database
CREATE DATABASE lp_maarif_nu CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE lp_maarif_nu;

-- ============================================
-- 1. USERS & AUTHENTICATION
-- ============================================

CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role ENUM('super_admin', 'admin', 'editor', 'viewer') DEFAULT 'viewer',
    avatar VARCHAR(500),
    phone VARCHAR(20),
    is_active BOOLEAN DEFAULT TRUE,
    last_login_at TIMESTAMP NULL,
    email_verified_at TIMESTAMP NULL,
    remember_token VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_email (email),
    INDEX idx_role (role),
    INDEX idx_is_active (is_active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE password_resets (
    email VARCHAR(255) NOT NULL,
    token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_email (email),
    INDEX idx_token (token)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE personal_access_tokens (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    tokenable_type VARCHAR(255) NOT NULL,
    tokenable_id BIGINT UNSIGNED NOT NULL,
    name VARCHAR(255) NOT NULL,
    token VARCHAR(64) NOT NULL UNIQUE,
    abilities TEXT,
    last_used_at TIMESTAMP NULL,
    expires_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_tokenable (tokenable_type, tokenable_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================
-- 2. CATEGORIES & TAGS
-- ============================================

CREATE TABLE categories (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    type ENUM('news', 'opinion', 'document') NOT NULL,
    icon VARCHAR(100),
    color VARCHAR(7),
    is_active BOOLEAN DEFAULT TRUE,
    order_number INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_slug (slug),
    INDEX idx_type (type),
    INDEX idx_is_active (is_active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE tags (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_slug (slug)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================
-- 3. NEWS ARTICLES (BERITA)
-- ============================================

CREATE TABLE news_articles (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(500) NOT NULL,
    slug VARCHAR(500) NOT NULL UNIQUE,
    excerpt TEXT NOT NULL,
    content LONGTEXT NOT NULL,
    image VARCHAR(500),
    category_id INT UNSIGNED,
    author_id BIGINT UNSIGNED,
    status ENUM('draft', 'published', 'archived') DEFAULT 'draft',
    published_at TIMESTAMP NULL,
    views INT UNSIGNED DEFAULT 0,
    is_featured BOOLEAN DEFAULT FALSE,
    meta_title VARCHAR(255),
    meta_description TEXT,
    meta_keywords VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL,
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE SET NULL,
    INDEX idx_slug (slug),
    INDEX idx_status (status),
    INDEX idx_published_at (published_at),
    INDEX idx_views (views),
    INDEX idx_is_featured (is_featured),
    FULLTEXT idx_search (title, excerpt, content)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE news_tags (
    news_article_id BIGINT UNSIGNED NOT NULL,
    tag_id INT UNSIGNED NOT NULL,
    PRIMARY KEY (news_article_id, tag_id),
    FOREIGN KEY (news_article_id) REFERENCES news_articles(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================
-- 4. OPINION ARTICLES (OPINI)
-- ============================================

CREATE TABLE opinion_articles (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(500) NOT NULL,
    slug VARCHAR(500) NOT NULL UNIQUE,
    excerpt TEXT NOT NULL,
    content LONGTEXT NOT NULL,
    image VARCHAR(500),
    author_name VARCHAR(255) NOT NULL,
    author_title VARCHAR(255),
    author_image VARCHAR(500),
    author_bio TEXT,
    status ENUM('draft', 'published', 'archived') DEFAULT 'draft',
    published_at TIMESTAMP NULL,
    views INT UNSIGNED DEFAULT 0,
    is_featured BOOLEAN DEFAULT FALSE,
    meta_title VARCHAR(255),
    meta_description TEXT,
    meta_keywords VARCHAR(500),
    created_by BIGINT UNSIGNED,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
    INDEX idx_slug (slug),
    INDEX idx_status (status),
    INDEX idx_published_at (published_at),
    INDEX idx_views (views),
    INDEX idx_is_featured (is_featured),
    FULLTEXT idx_search (title, excerpt, content)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE opinion_tags (
    opinion_article_id BIGINT UNSIGNED NOT NULL,
    tag_id INT UNSIGNED NOT NULL,
    PRIMARY KEY (opinion_article_id, tag_id),
    FOREIGN KEY (opinion_article_id) REFERENCES opinion_articles(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================
-- 5. DOCUMENTS (DOKUMEN)
-- ============================================

CREATE TABLE documents (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(500) NOT NULL,
    description TEXT,
    category_id INT UNSIGNED,
    file_name VARCHAR(255) NOT NULL,
    file_path VARCHAR(500) NOT NULL,
    file_type VARCHAR(50) NOT NULL,
    file_size BIGINT UNSIGNED NOT NULL COMMENT 'in bytes',
    mime_type VARCHAR(100),
    download_count INT UNSIGNED DEFAULT 0,
    is_public BOOLEAN DEFAULT TRUE,
    uploaded_by BIGINT UNSIGNED,
    status ENUM('active', 'archived') DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL,
    FOREIGN KEY (uploaded_by) REFERENCES users(id) ON DELETE SET NULL,
    INDEX idx_file_type (file_type),
    INDEX idx_is_public (is_public),
    INDEX idx_status (status),
    INDEX idx_download_count (download_count),
    FULLTEXT idx_search (title, description)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================
-- 6. HERO SLIDES (HOME PAGE)
-- ============================================

CREATE TABLE hero_slides (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(500) NOT NULL,
    description TEXT,
    image VARCHAR(500) NOT NULL,
    cta_label VARCHAR(100),
    cta_href VARCHAR(500),
    cta_secondary_label VARCHAR(100),
    cta_secondary_href VARCHAR(500),
    order_number INT DEFAULT 0,
    is_active BOOLEAN DEFAULT TRUE,
    start_date TIMESTAMP NULL,
    end_date TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_order_number (order_number),
    INDEX idx_is_active (is_active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================
-- 7. ORGANIZATION (STRUKTUR ORGANISASI)
-- ============================================

CREATE TABLE organization_positions (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    position_name VARCHAR(255) NOT NULL,
    position_level INT NOT NULL COMMENT '1=Ketua, 2=Wakil, 3=Sekretaris, 4=Bendahara, 5=Bidang',
    position_type ENUM('ketua', 'wakil', 'sekretaris', 'bendahara', 'bidang') NOT NULL,
    parent_id INT UNSIGNED,
    order_number INT DEFAULT 0,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES organization_positions(id) ON DELETE SET NULL,
    INDEX idx_position_level (position_level),
    INDEX idx_position_type (position_type),
    INDEX idx_order_number (order_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE board_members (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    position_id INT UNSIGNED NOT NULL,
    name VARCHAR(255) NOT NULL,
    title VARCHAR(255) COMMENT 'Gelar akademik',
    photo VARCHAR(500),
    bio TEXT,
    email VARCHAR(255),
    phone VARCHAR(20),
    social_media JSON COMMENT '{"facebook": "", "twitter": "", "linkedin": ""}',
    period_start YEAR NOT NULL,
    period_end YEAR NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    order_number INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (position_id) REFERENCES organization_positions(id) ON DELETE CASCADE,
    INDEX idx_period (period_start, period_end),
    INDEX idx_is_active (is_active),
    INDEX idx_order_number (order_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE departments (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    head_name VARCHAR(255),
    order_number INT DEFAULT 0,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_order_number (order_number),
    INDEX idx_is_active (is_active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================
-- 8. CONTENT PAGES (STATIC PAGES)
-- ============================================

CREATE TABLE pages (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    slug VARCHAR(255) NOT NULL UNIQUE,
    title VARCHAR(500) NOT NULL,
    content LONGTEXT,
    metadata JSON COMMENT 'Flexible content based on page type',
    template VARCHAR(100) DEFAULT 'default',
    is_active BOOLEAN DEFAULT TRUE,
    meta_title VARCHAR(255),
    meta_description TEXT,
    meta_keywords VARCHAR(500),
    last_updated_by BIGINT UNSIGNED,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (last_updated_by) REFERENCES users(id) ON DELETE SET NULL,
    INDEX idx_slug (slug),
    INDEX idx_is_active (is_active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================
-- 9. MEDIA LIBRARY
-- ============================================

CREATE TABLE media (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    file_name VARCHAR(255) NOT NULL,
    original_name VARCHAR(255) NOT NULL,
    file_path VARCHAR(500) NOT NULL,
    file_url VARCHAR(500) NOT NULL,
    file_type VARCHAR(50) NOT NULL,
    mime_type VARCHAR(100),
    file_size BIGINT UNSIGNED NOT NULL COMMENT 'in bytes',
    width INT UNSIGNED COMMENT 'for images',
    height INT UNSIGNED COMMENT 'for images',
    folder VARCHAR(100) DEFAULT 'general',
    alt_text VARCHAR(255),
    caption TEXT,
    uploaded_by BIGINT UNSIGNED,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (uploaded_by) REFERENCES users(id) ON DELETE SET NULL,
    INDEX idx_file_type (file_type),
    INDEX idx_folder (folder),
    INDEX idx_uploaded_by (uploaded_by)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================
-- 10. SETTINGS
-- ============================================

CREATE TABLE settings (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    setting_key VARCHAR(100) NOT NULL UNIQUE,
    setting_value TEXT,
    setting_type ENUM('string', 'text', 'number', 'boolean', 'json') DEFAULT 'string',
    setting_group VARCHAR(50) DEFAULT 'general',
    description TEXT,
    is_public BOOLEAN DEFAULT FALSE COMMENT 'Can be accessed without auth',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_setting_key (setting_key),
    INDEX idx_setting_group (setting_group),
    INDEX idx_is_public (is_public)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================
-- 11. ACTIVITY LOGS
-- ============================================

CREATE TABLE activity_logs (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    log_name VARCHAR(100),
    description TEXT,
    subject_type VARCHAR(255),
    subject_id BIGINT UNSIGNED,
    causer_type VARCHAR(255),
    causer_id BIGINT UNSIGNED,
    properties JSON,
    ip_address VARCHAR(45),
    user_agent TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_log_name (log_name),
    INDEX idx_subject (subject_type, subject_id),
    INDEX idx_causer (causer_type, causer_id),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================
-- 12. ANALYTICS
-- ============================================

CREATE TABLE page_views (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    viewable_type VARCHAR(255) NOT NULL,
    viewable_id BIGINT UNSIGNED NOT NULL,
    ip_address VARCHAR(45),
    user_agent TEXT,
    referer VARCHAR(500),
    session_id VARCHAR(255),
    viewed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_viewable (viewable_type, viewable_id),
    INDEX idx_viewed_at (viewed_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE download_logs (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    document_id BIGINT UNSIGNED NOT NULL,
    ip_address VARCHAR(45),
    user_agent TEXT,
    downloaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (document_id) REFERENCES documents(id) ON DELETE CASCADE,
    INDEX idx_document_id (document_id),
    INDEX idx_downloaded_at (downloaded_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================
-- 13. NOTIFICATIONS
-- ============================================

CREATE TABLE notifications (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    type VARCHAR(100) NOT NULL,
    title VARCHAR(255) NOT NULL,
    message TEXT,
    data JSON,
    read_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_user_id (user_id),
    INDEX idx_read_at (read_at),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================
-- 14. CACHE TABLE (Optional - for database caching)
-- ============================================

CREATE TABLE cache (
    cache_key VARCHAR(255) PRIMARY KEY,
    cache_value MEDIUMTEXT NOT NULL,
    expiration INT NOT NULL,
    INDEX idx_expiration (expiration)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE cache_locks (
    lock_key VARCHAR(255) PRIMARY KEY,
    owner VARCHAR(255) NOT NULL,
    expiration INT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================
-- INSERT DEFAULT DATA
-- ============================================

-- Insert default admin user (password: password)
INSERT INTO users (name, email, password, role, is_active) VALUES
('Super Admin', 'admin@lpmaarifnu.or.id', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'super_admin', TRUE);

-- Insert default categories for news
INSERT INTO categories (name, slug, description, type, is_active, order_number) VALUES
('Nasional', 'nasional', 'Berita tingkat nasional', 'news', TRUE, 1),
('Daerah', 'daerah', 'Berita tingkat daerah', 'news', TRUE, 2),
('Program', 'program', 'Program dan kegiatan', 'news', TRUE, 3),
('Pengumuman', 'pengumuman', 'Pengumuman resmi', 'news', TRUE, 4);

-- Insert default categories for documents
INSERT INTO categories (name, slug, description, type, is_active, order_number) VALUES
('Pedoman', 'pedoman', 'Pedoman dan panduan', 'document', TRUE, 1),
('Kurikulum', 'kurikulum', 'Dokumen kurikulum', 'document', TRUE, 2),
('Regulasi', 'regulasi', 'Peraturan dan regulasi', 'document', TRUE, 3),
('Panduan', 'panduan', 'Panduan teknis', 'document', TRUE, 4),
('Formulir', 'formulir', 'Formulir dan template', 'document', TRUE, 5);

-- Insert organization positions
INSERT INTO organization_positions (position_name, position_level, position_type, order_number) VALUES
('Ketua Umum', 1, 'ketua', 1),
('Wakil Ketua I', 2, 'wakil', 2),
('Wakil Ketua II', 2, 'wakil', 3),
('Sekretaris Umum', 3, 'sekretaris', 4),
('Bendahara Umum', 4, 'bendahara', 5);

-- Insert departments
INSERT INTO departments (name, description, order_number, is_active) VALUES
('Bidang Pendidikan Dasar', 'Mengelola pengembangan pendidikan tingkat SD/MI', 1, TRUE),
('Bidang Pendidikan Menengah', 'Mengelola pengembangan pendidikan tingkat SMP/MTs dan SMA/MA', 2, TRUE),
('Bidang Pendidikan Tinggi', 'Mengelola perguruan tinggi di bawah LP Ma''arif NU', 3, TRUE),
('Bidang Kurikulum', 'Mengembangkan kurikulum berbasis Ma''arif NU', 4, TRUE),
('Bidang SDM dan Kemitraan', 'Mengembangkan SDM dan kerjasama kelembagaan', 5, TRUE),
('Bidang Penelitian dan Pengembangan', 'Melakukan penelitian dan pengembangan pendidikan', 6, TRUE);

-- Insert default settings
INSERT INTO settings (setting_key, setting_value, setting_type, setting_group, description, is_public) VALUES
('site_name', 'LP Ma''arif NU', 'string', 'general', 'Nama website', TRUE),
('site_description', 'Lembaga Pendidikan Ma''arif Nahdlatul Ulama', 'text', 'general', 'Deskripsi website', TRUE),
('contact_email', 'info@lpmaarifnu.or.id', 'string', 'contact', 'Email kontak', TRUE),
('contact_phone', '021-1234567', 'string', 'contact', 'Nomor telepon', TRUE),
('contact_address', 'Jakarta', 'text', 'contact', 'Alamat', TRUE),
('social_facebook', '', 'string', 'social', 'Facebook URL', TRUE),
('social_twitter', '', 'string', 'social', 'Twitter URL', TRUE),
('social_instagram', '', 'string', 'social', 'Instagram URL', TRUE),
('social_youtube', '', 'string', 'social', 'YouTube URL', TRUE);

-- Insert default pages
INSERT INTO pages (slug, title, template, is_active) VALUES
('visi-misi', 'Visi & Misi', 'visi-misi', TRUE),
('sejarah', 'Sejarah LP Ma''arif NU', 'sejarah', TRUE),
('program-strategis', 'Program Strategis', 'program-strategis', TRUE),
('pramuka', 'Pramuka Ma''arif', 'pramuka', TRUE);

-- ============================================
-- CREATE VIEWS FOR COMMON QUERIES
-- ============================================

-- View for published news with category and tags
CREATE VIEW v_published_news AS
SELECT
    n.id,
    n.title,
    n.slug,
    n.excerpt,
    n.content,
    n.image,
    n.published_at,
    n.views,
    n.is_featured,
    c.name as category_name,
    c.slug as category_slug,
    u.name as author_name,
    GROUP_CONCAT(t.name) as tags
FROM news_articles n
LEFT JOIN categories c ON n.category_id = c.id
LEFT JOIN users u ON n.author_id = u.id
LEFT JOIN news_tags nt ON n.id = nt.news_article_id
LEFT JOIN tags t ON nt.tag_id = t.id
WHERE n.status = 'published' AND n.deleted_at IS NULL
GROUP BY n.id;

-- ============================================
-- CREATE STORED PROCEDURES
-- ============================================

DELIMITER $$

-- Procedure to increment view count
CREATE PROCEDURE increment_view_count(
    IN p_viewable_type VARCHAR(255),
    IN p_viewable_id BIGINT UNSIGNED
)
BEGIN
    IF p_viewable_type = 'news_articles' THEN
        UPDATE news_articles SET views = views + 1 WHERE id = p_viewable_id;
    ELSEIF p_viewable_type = 'opinion_articles' THEN
        UPDATE opinion_articles SET views = views + 1 WHERE id = p_viewable_id;
    END IF;
END$$

-- Procedure to get popular articles
CREATE PROCEDURE get_popular_articles(
    IN p_limit INT,
    IN p_days INT
)
BEGIN
    SELECT
        id, title, slug, views, published_at
    FROM news_articles
    WHERE
        status = 'published'
        AND deleted_at IS NULL
        AND published_at >= DATE_SUB(NOW(), INTERVAL p_days DAY)
    ORDER BY views DESC
    LIMIT p_limit;
END$$

DELIMITER ;

-- ============================================
-- CREATE TRIGGERS
-- ============================================

DELIMITER $$

-- Trigger to auto-generate slug before insert
CREATE TRIGGER news_articles_before_insert
BEFORE INSERT ON news_articles
FOR EACH ROW
BEGIN
    IF NEW.slug IS NULL OR NEW.slug = '' THEN
        SET NEW.slug = LOWER(REPLACE(NEW.title, ' ', '-'));
    END IF;
END$$

-- Trigger to log document downloads
CREATE TRIGGER after_document_download
AFTER INSERT ON download_logs
FOR EACH ROW
BEGIN
    UPDATE documents
    SET download_count = download_count + 1
    WHERE id = NEW.document_id;
END$$

DELIMITER ;

-- ============================================
-- INDEXES FOR PERFORMANCE
-- ============================================

-- Additional composite indexes
CREATE INDEX idx_news_published ON news_articles(status, published_at, is_featured);
CREATE INDEX idx_opinion_published ON opinion_articles(status, published_at, is_featured);
CREATE INDEX idx_documents_category_type ON documents(category_id, file_type, status);

-- ============================================
-- COMPLETED
-- ============================================

-- Show summary
SELECT
    'Database schema created successfully!' as Status,
    'lp_maarif_nu' as Database_Name,
    COUNT(*) as Total_Tables
FROM information_schema.tables
WHERE table_schema = 'lp_maarif_nu';
