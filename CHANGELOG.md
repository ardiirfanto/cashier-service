# Changelog

All notable changes to the Cashier API project will be documented in this file.

## [1.1.0] - 2025-10-30

### Added
- **Image field for menu items** 🖼️
  - Added `image` column to `menus` table (VARCHAR 255)
  - Updated `Menu` model with `Image` field
  - Sample menu items now include image URLs from Unsplash
  - Migration script `add_image_column.sql` for existing databases

### Changed
- Updated `database_setup.sql` to include image column
- Updated sample data with high-quality food/beverage images

### API Response Change
Menu endpoints now return image URLs:
```json
{
  "id": 1,
  "name": "Espresso",
  "price": 25000.00,
  "stock": 100,
  "image": "https://images.unsplash.com/photo-1510591509098-f4fdc6d0ff04?w=400",
  "created_at": "2024-01-01T00:00:00Z"
}
```

### Migration Guide

#### For New Setup:
Just run the updated database setup:
```bash
mysql -u root -p < database_setup.sql
```

#### For Existing Database:
Run the migration script:
```bash
mysql -u root -p < add_image_column.sql
```

Or manually with GORM auto-migration (restart server and it will auto-add the column):
```bash
go run ./cmd/server
```

---

## [1.0.0] - 2025-10-30

### Added
- Initial release of Cashier API Backend
- JWT Authentication with bcrypt password hashing
- Concurrent checkout processing with goroutines and channels
- RESTful API endpoints:
  - POST `/api/login` - Login and get JWT token
  - GET `/api/menus` - Get all menu items
  - POST `/api/checkout` - Process checkout
  - GET `/api/transactions` - Get transaction history
  - GET `/health` - Health check
- Clean Architecture pattern (Handler → Service → Repository → Model)
- MySQL database with GORM ORM
- Viper configuration management
- Complete documentation (13 files in `docs/` folder)
- Docker-ready architecture
- Production deployment guides

### Features
- ✅ Secure authentication (JWT + bcrypt)
- ✅ High-performance concurrent processing
- ✅ Clean and maintainable code structure
- ✅ Auto-migration support
- ✅ Comprehensive error handling
- ✅ Production-ready deployment

---

## Upgrade Instructions

### From 1.0.0 to 1.1.0

1. **Update Go code:**
   - Pull latest changes (model updated)
   - Restart server (GORM will auto-migrate)

2. **Update database (if not using auto-migration):**
   ```bash
   mysql -u root -p < add_image_column.sql
   ```

3. **No breaking changes:**
   - Existing endpoints still work
   - Image field is optional (can be NULL)
   - Backward compatible

---

## Image Sources

All sample images are from [Unsplash](https://unsplash.com) - free high-quality images:
- Coffee images: Professional espresso, cappuccino, latte shots
- Tea images: Green tea and black tea in elegant cups
- Food images: Sandwiches, burgers, fries, chicken wings
- Beverage images: Fresh juices and smoothies
- Dessert images: Cakes and cheesecakes

Images are optimized to 400px width for faster loading.

---

## Notes

- GORM auto-migration will automatically add the `image` column when you restart the server
- If you prefer manual migration, use `add_image_column.sql`
- Images are stored as URLs (external links), not uploaded files
- For production, consider using a CDN or S3 bucket for image hosting
