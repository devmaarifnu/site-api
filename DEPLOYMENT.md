# Deployment Guide - LP Ma'arif NU Backend API

## Production Deployment

### Prerequisites
- Linux server (Ubuntu 20.04+ recommended)
- MySQL 8.0+ installed
- Domain name (e.g., api.lpmaarifnu.or.id)
- SSL certificate

## Option 1: Systemd Service (Recommended)

### 1. Build for Production

On your local machine:
```bash
# Build for Linux
GOOS=linux GOARCH=amd64 go build -o bin/api-linux cmd/api/main.go

# Or use Make
make build-linux
```

### 2. Upload to Server

```bash
# Create directory
ssh user@server "mkdir -p /var/www/lpmaarifnu-site-api"

# Upload files
scp -r . user@server:/var/www/lpmaarifnu-site-api/

# Or use rsync
rsync -avz --exclude 'bin' --exclude 'tmp' . user@server:/var/www/lpmaarifnu-site-api/
```

### 3. Setup Database

On server:
```bash
# Create database
mysql -u root -p -e "CREATE DATABASE lp_maarif_nu CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# Import schema
mysql -u root -p lp_maarif_nu < /var/www/lpmaarifnu-site-api/database_schema.sql
```

### 4. Configure Application

Edit `/var/www/lpmaarifnu-site-api/config.yaml`:

```yaml
app:
  name: "LP Maarif API"
  env: "production"
  port: 8080
  version: "1.0.0"

database:
  host: "localhost"
  port: 3306
  user: "lpmaarif_user"
  password: "secure_password_here"
  name: "lp_maarif_nu"
  max_idle_conns: 10
  max_open_conns: 100
  conn_max_lifetime: 3600

cors:
  allowed_origins:
    - "https://lpmaarifnu.or.id"
    - "https://www.lpmaarifnu.or.id"
  allowed_methods:
    - "GET"
    - "OPTIONS"
  allowed_headers:
    - "Content-Type"
    - "Authorization"

rate_limit:
  enabled: true
  requests: 100
  window: 60

storage:
  cdn_url: "https://cdn.lpmaarifnu.or.id"

logging:
  level: "info"
  format: "json"
```

### 5. Create Systemd Service

Create `/etc/systemd/system/lpmaarifnu-site-api.service`:

```ini
[Unit]
Description=LP Maarif NU Backend API
Documentation=https://github.com/lpmaarifnu/api
After=network.target mysql.service

[Service]
Type=simple
User=www-data
Group=www-data
WorkingDirectory=/var/www/lpmaarifnu-site-api
ExecStart=/var/www/lpmaarifnu-site-api/bin/api-linux
Restart=on-failure
RestartSec=5s
StandardOutput=journal
StandardError=journal
SyslogIdentifier=lpmaarifnu-site-api

# Security settings
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ProtectHome=true
ReadWritePaths=/var/www/lpmaarifnu-site-api

# Environment
Environment="GIN_MODE=release"

[Install]
WantedBy=multi-user.target
```

### 6. Start Service

```bash
# Set permissions
sudo chown -R www-data:www-data /var/www/lpmaarifnu-site-api
sudo chmod +x /var/www/lpmaarifnu-site-api/bin/api-linux

# Reload systemd
sudo systemctl daemon-reload

# Enable and start service
sudo systemctl enable lpmaarifnu-site-api
sudo systemctl start lpmaarifnu-site-api

# Check status
sudo systemctl status lpmaarifnu-site-api

# View logs
sudo journalctl -u lpmaarifnu-site-api -f
```

### 7. Setup Nginx Reverse Proxy

Create `/etc/nginx/sites-available/lpmaarifnu-site-api`:

```nginx
server {
    listen 80;
    server_name api.lpmaarifnu.or.id;

    # Redirect to HTTPS
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name api.lpmaarifnu.or.id;

    # SSL Configuration
    ssl_certificate /etc/letsencrypt/live/api.lpmaarifnu.or.id/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/api.lpmaarifnu.or.id/privkey.pem;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers HIGH:!aNULL:!MD5;

    # Logging
    access_log /var/log/nginx/lpmaarifnu-site-api-access.log;
    error_log /var/log/nginx/lpmaarifnu-site-api-error.log;

    # Proxy to Go application
    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;

        # Timeouts
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }

    # Health check endpoint
    location /health {
        proxy_pass http://localhost:8080/health;
        access_log off;
    }

    # Security headers
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header Referrer-Policy "no-referrer-when-downgrade" always;
}
```

Enable site:
```bash
sudo ln -s /etc/nginx/sites-available/lpmaarifnu-site-api /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

### 8. Setup SSL with Let's Encrypt

```bash
# Install certbot
sudo apt install certbot python3-certbot-nginx

# Get certificate
sudo certbot --nginx -d api.lpmaarifnu.or.id

# Auto-renewal is set up automatically
# Test renewal
sudo certbot renew --dry-run
```

### 9. Verify Deployment

```bash
# Check service
curl http://localhost:8080/health

# Check through Nginx
curl https://api.lpmaarifnu.or.id/health

# Test API endpoint
curl https://api.lpmaarifnu.or.id/api/v1/categories
```

## Option 2: Docker Deployment

### 1. Create Dockerfile

Create `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/api/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy binary and config
COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .

EXPOSE 8080

CMD ["./main"]
```

### 2. Create docker-compose.yml

```yaml
version: '3.8'

services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - GIN_MODE=release
    volumes:
      - ./config.yaml:/root/config.yaml:ro
    depends_on:
      - db
    restart: unless-stopped

  db:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: root_password
      MYSQL_DATABASE: lp_maarif_nu
      MYSQL_USER: lpmaarif_user
      MYSQL_PASSWORD: secure_password
    volumes:
      - mysql_data:/var/lib/mysql
      - ./database_schema.sql:/docker-entrypoint-initdb.d/schema.sql
    ports:
      - "3306:3306"
    restart: unless-stopped

volumes:
  mysql_data:
```

### 3. Deploy with Docker

```bash
# Build and start
docker-compose up -d

# View logs
docker-compose logs -f api

# Stop
docker-compose down

# Rebuild
docker-compose up -d --build
```

## Monitoring & Maintenance

### View Logs

```bash
# Systemd
sudo journalctl -u lpmaarifnu-site-api -f

# Docker
docker-compose logs -f api

# Nginx access logs
sudo tail -f /var/log/nginx/lpmaarifnu-site-api-access.log

# Nginx error logs
sudo tail -f /var/log/nginx/lpmaarifnu-site-api-error.log
```

### Restart Service

```bash
# Systemd
sudo systemctl restart lpmaarifnu-site-api

# Docker
docker-compose restart api
```

### Update Application

```bash
# Systemd
cd /var/www/lpmaarifnu-site-api
git pull  # or upload new files
sudo systemctl restart lpmaarifnu-site-api

# Docker
docker-compose down
docker-compose up -d --build
```

### Database Backup

```bash
# Backup
mysqldump -u root -p lp_maarif_nu > backup_$(date +%Y%m%d).sql

# Restore
mysql -u root -p lp_maarif_nu < backup_20250111.sql

# Automated backup (crontab)
0 2 * * * mysqldump -u root -p'password' lp_maarif_nu > /backup/lpmaarif_$(date +\%Y\%m\%d).sql
```

## Performance Tuning

### MySQL Optimization

Edit `/etc/mysql/mysql.conf.d/mysqld.cnf`:

```ini
[mysqld]
max_connections = 500
innodb_buffer_pool_size = 1G
innodb_log_file_size = 256M
query_cache_size = 0
query_cache_type = 0
```

### Nginx Optimization

```nginx
# Enable gzip
gzip on;
gzip_vary on;
gzip_min_length 1024;
gzip_types application/json text/css application/javascript;

# Buffer sizes
client_body_buffer_size 128k;
client_max_body_size 10m;
```

## Security Checklist

- [ ] Change all default passwords
- [ ] Enable firewall (ufw)
- [ ] Setup SSL certificate
- [ ] Configure CORS properly
- [ ] Enable rate limiting
- [ ] Regular security updates
- [ ] Database user with limited privileges
- [ ] Disable root MySQL login
- [ ] Setup fail2ban
- [ ] Regular backups

## Troubleshooting

### Service won't start
```bash
# Check logs
sudo journalctl -u lpmaarifnu-site-api -n 50

# Check config
cd /var/www/lpmaarifnu-site-api
./bin/api-linux  # Run manually to see errors
```

### Database connection error
```bash
# Check MySQL is running
sudo systemctl status mysql

# Test connection
mysql -u lpmaarif_user -p lp_maarif_nu
```

### High memory usage
```bash
# Check memory
free -h

# Optimize MySQL
# Reduce innodb_buffer_pool_size

# Restart service
sudo systemctl restart lpmaarifnu-site-api
```

## Support

For deployment issues:
- GitHub Issues: https://github.com/lpmaarifnu/api/issues
- Email: dev@lpmaarifnu.or.id
- Documentation: See README.md and SETUP.md
