# API Documentation - LP Ma'arif NU

## Base Information

- **Base URL Development:** `http://localhost:8080/api/v1`
- **Base URL Production:** `https://api.lpmaarifnu.or.id/v1`
- **Response Format:** JSON
- **API Type:** Read-Only (GET only)
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

### 9. Health Check

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

- **Limit:** 100 requests per minute per IP
- **Response when exceeded:**
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
```

### cURL
```bash
# Get categories
curl http://localhost:8080/api/v1/categories

# Get news with filters
curl "http://localhost:8080/api/v1/news?category=nasional&page=1&limit=10"

# Get settings
curl http://localhost:8080/api/v1/settings
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
```

## Best Practices

1. **Always handle errors** - Check `success` field in response
2. **Use pagination** - Don't request too many items at once
3. **Cache responses** - Use client-side caching for better performance
4. **Respect rate limits** - Implement retry logic with exponential backoff
5. **Handle 404s gracefully** - Show user-friendly error messages

## Support

For API support or questions:
- GitHub Issues: https://github.com/lpmaarifnu/api/issues
- Email: dev@lpmaarifnu.or.id
- Documentation: See `TODO BACKEND - READ ONLY API.md` for detailed specs
