# 🖼️ Menu Image Feature

## Overview

Field `image` telah ditambahkan ke menu items untuk mendukung tampilan visual di frontend (React/Next.js web dan Flutter mobile).

---

## Database Schema

### Menu Table Structure

```sql
CREATE TABLE menus (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  stock INT DEFAULT 0,
  image VARCHAR(255),              -- ✨ NEW FIELD
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Field Details:**
- **Type**: VARCHAR(255)
- **Nullable**: Yes (NULL allowed)
- **Purpose**: Store URL to menu item image
- **Location**: After `stock` column

---

## Go Model

### Updated Menu Model

```go
type Menu struct {
    ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Name      string    `gorm:"type:varchar(100);not null" json:"name"`
    Price     float64   `gorm:"type:decimal(10,2);not null" json:"price"`
    Stock     int       `gorm:"type:int;default:0" json:"stock"`
    Image     string    `gorm:"type:varchar(255)" json:"image"`  // ✨ NEW
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
```

---

## API Response

### GET /api/menus

Response sekarang include field `image`:

```json
{
  "success": true,
  "message": "Menus retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Espresso",
      "price": 25000.00,
      "stock": 100,
      "image": "https://images.unsplash.com/photo-1510591509098-f4fdc6d0ff04?w=400",
      "created_at": "2024-01-01T00:00:00Z"
    },
    {
      "id": 2,
      "name": "Cappuccino",
      "price": 30000.00,
      "stock": 100,
      "image": "https://images.unsplash.com/photo-1572442388796-11668a67e53d?w=400",
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

---

## Sample Data

Semua menu items sudah include gambar dari Unsplash:

| Menu Item | Image URL |
|-----------|-----------|
| Espresso | `https://images.unsplash.com/photo-1510591509098-f4fdc6d0ff04?w=400` |
| Cappuccino | `https://images.unsplash.com/photo-1572442388796-11668a67e53d?w=400` |
| Latte | `https://images.unsplash.com/photo-1561882468-9110e03e0f78?w=400` |
| Americano | `https://images.unsplash.com/photo-1532004491497-ba35c367d634?w=400` |
| Green Tea | `https://images.unsplash.com/photo-1556679343-c7306c1976bc?w=400` |
| Black Tea | `https://images.unsplash.com/photo-1594631252845-29fc4cc8cde9?w=400` |
| Sandwich | `https://images.unsplash.com/photo-1528735602780-2552fd46c7af?w=400` |
| Burger | `https://images.unsplash.com/photo-1568901346375-23c9450c58cd?w=400` |
| French Fries | `https://images.unsplash.com/photo-1573080496219-bb080dd4f877?w=400` |
| Chicken Wings | `https://images.unsplash.com/photo-1527477396000-e27163b481c2?w=400` |
| Orange Juice | `https://images.unsplash.com/photo-1600271886742-f049cd451bba?w=400` |
| Apple Juice | `https://images.unsplash.com/photo-1576673442511-7e39b6545c87?w=400` |
| Mango Smoothie | `https://images.unsplash.com/photo-1623065422902-30a2d299bbe4?w=400` |
| Chocolate Cake | `https://images.unsplash.com/photo-1578985545062-69928b1d9587?w=400` |
| Cheesecake | `https://images.unsplash.com/photo-1533134486753-c833f0ed4866?w=400` |

**Note:** Images optimized to 400px width (`?w=400`) untuk faster loading.

---

## Migration Guide

### Option 1: Fresh Database Setup

Jika setup database baru, gunakan file yang sudah updated:

```bash
mysql -u root -p < database_setup.sql
```

### Option 2: Update Existing Database

Jika sudah ada database, run migration script:

```bash
mysql -u root -p < add_image_column.sql
```

### Option 3: GORM Auto-Migration

GORM akan otomatis menambahkan kolom saat server dijalankan:

```bash
# Restart server
go run ./cmd/server

# GORM akan detect perubahan model dan auto-add column
```

---

## Frontend Implementation

### React/Next.js Example

```jsx
// Fetch menus
const fetchMenus = async () => {
  const response = await fetch('http://localhost:8080/api/menus', {
    headers: {
      'Authorization': `Bearer ${token}`
    }
  });
  const data = await response.json();
  return data.data;
};

// Display menu with image
const MenuCard = ({ menu }) => (
  <div className="menu-card">
    <img
      src={menu.image}
      alt={menu.name}
      className="menu-image"
      loading="lazy"
    />
    <h3>{menu.name}</h3>
    <p>Rp {menu.price.toLocaleString()}</p>
    <p>Stock: {menu.stock}</p>
  </div>
);
```

### Flutter Example

```dart
// Menu model
class Menu {
  final int id;
  final String name;
  final double price;
  final int stock;
  final String? image;  // Nullable
  final DateTime createdAt;

  Menu.fromJson(Map<String, dynamic> json)
      : id = json['id'],
        name = json['name'],
        price = json['price'],
        stock = json['stock'],
        image = json['image'],
        createdAt = DateTime.parse(json['created_at']);
}

// Display menu
Widget buildMenuCard(Menu menu) {
  return Card(
    child: Column(
      children: [
        menu.image != null
            ? Image.network(
                menu.image!,
                height: 150,
                width: double.infinity,
                fit: BoxFit.cover,
                loadingBuilder: (context, child, loadingProgress) {
                  if (loadingProgress == null) return child;
                  return CircularProgressIndicator();
                },
              )
            : Container(
                height: 150,
                color: Colors.grey[300],
                child: Icon(Icons.image, size: 50),
              ),
        Text(menu.name, style: TextStyle(fontSize: 18)),
        Text('Rp ${menu.price.toStringAsFixed(0)}'),
        Text('Stock: ${menu.stock}'),
      ],
    ),
  );
}
```

---

## Image Hosting Options

### Current: Unsplash (External)
- ✅ Free high-quality images
- ✅ Fast CDN delivery
- ⚠️ External dependency
- ⚠️ Need internet connection

### Production Alternatives:

#### 1. AWS S3 + CloudFront
```
Cost: ~$0.023/GB storage + $0.085/GB transfer
Setup:
1. Create S3 bucket
2. Upload images
3. Configure CloudFront CDN
4. Update image URLs in database
```

#### 2. Google Cloud Storage
```
Cost: ~$0.020/GB storage + $0.12/GB egress
Setup:
1. Create GCS bucket
2. Upload images
3. Make bucket public or use signed URLs
4. Update image URLs
```

#### 3. Cloudinary
```
Cost: Free tier 25GB/month, then $89/month
Setup:
1. Sign up at cloudinary.com
2. Upload via API or dashboard
3. Use Cloudinary URLs
4. Auto-optimization included
```

#### 4. Self-Hosted
```
Cost: Server bandwidth only
Setup:
1. Create /static/images folder
2. Serve with Gin static files
3. Upload images to server
4. Use relative URLs
```

**Self-Hosted Example:**

```go
// In router.go
router.Static("/images", "./static/images")

// Image URL example:
// http://localhost:8080/images/espresso.jpg
```

---

## Best Practices

### Image Optimization

1. **Size**: Resize to appropriate dimensions (400x400px for thumbnails)
2. **Format**: Use WebP for modern browsers, JPG fallback
3. **Compression**: Optimize quality (80-85% usually sufficient)
4. **CDN**: Use CDN for faster global delivery
5. **Lazy Loading**: Implement lazy loading in frontend

### Database

1. **Store URLs, not files**: Never store binary images in database
2. **Nullable**: Make image field nullable (optional)
3. **Validation**: Validate URL format if accepting user input
4. **Default**: Consider default image for items without image

### Security

1. **Validate URLs**: If accepting user-uploaded images
2. **Sanitize**: Remove any scripts from image URLs
3. **CORS**: Configure proper CORS headers
4. **Rate Limiting**: Prevent abuse of image uploads

---

## Testing

### Test Menu Endpoint with Images

```bash
# Get all menus
curl -X GET http://localhost:8080/api/menus \
  -H "Authorization: Bearer YOUR_TOKEN" | jq

# Check image field exists
curl -X GET http://localhost:8080/api/menus \
  -H "Authorization: Bearer YOUR_TOKEN" | jq '.data[0].image'
```

### Expected Output

```json
{
  "success": true,
  "message": "Menus retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Espresso",
      "price": 25000,
      "stock": 100,
      "image": "https://images.unsplash.com/photo-1510591509098-f4fdc6d0ff04?w=400",
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

---

## Troubleshooting

### Column Doesn't Exist Error

**Error:**
```
Error 1054: Unknown column 'image' in 'field list'
```

**Solution:**
```bash
# Run migration script
mysql -u root -p < add_image_column.sql

# Or restart server for auto-migration
go run ./cmd/server
```

### Image Not Showing in Frontend

**Checklist:**
1. ✅ Check API response includes `image` field
2. ✅ Check image URL is valid (test in browser)
3. ✅ Check CORS headers if frontend on different domain
4. ✅ Check image URLs not blocked by firewall
5. ✅ Check frontend handles null/empty images

### Images Loading Slowly

**Solutions:**
1. Use CDN for image hosting
2. Implement lazy loading
3. Reduce image dimensions
4. Use WebP format
5. Add image caching headers

---

## Future Enhancements

### Planned Features:

1. **Image Upload API**
   - POST endpoint for image uploads
   - Integration with S3/Cloudinary
   - Image validation and resizing

2. **Multiple Images**
   - Support for image gallery per menu item
   - Different sizes (thumbnail, full size)

3. **Image Management**
   - Admin API for image CRUD
   - Image compression on upload
   - WebP conversion

4. **Caching**
   - Implement Redis caching for image URLs
   - CDN integration
   - Client-side caching

---

## Related Files

- **Model**: `internal/model/menu.go`
- **Database Setup**: `database_setup.sql`
- **Migration**: `add_image_column.sql`
- **API Docs**: `docs/API_TESTING.md`
- **Changelog**: `CHANGELOG.md`

---

**✅ Image feature is now live and ready to use!**

For questions or issues, check the main documentation in `docs/` folder.
