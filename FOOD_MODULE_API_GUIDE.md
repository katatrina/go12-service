# Food Module API Testing Guide

Hướng dẫn test API của Food Module cho các thành viên trong dự án.

## Prerequisite

### 1. Khởi động ứng dụng
```bash
# Đảm bảo MySQL và Jaeger containers đang chạy
go run main.go
```

### 2. Tạo user và lấy token
```bash
# Tạo user
curl -X POST http://localhost:8080/v1/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"Password123","first_name":"Test","last_name":"User"}'

# Lấy token
curl -X POST http://localhost:8080/v1/authenticate \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"Password123"}'

# Lưu token vào biến (thay YOUR_TOKEN bằng token thực)
export TOKEN="YOUR_TOKEN_HERE"
```

### 3. Tạo category và restaurant (cần cho food)
```bash
# Tạo category
curl -X POST http://localhost:8080/v1/categories \
  -H "Content-Type: application/json" \
  -d '{"name":"Vietnamese Food","description":"Traditional Vietnamese cuisine"}'
# → Lưu category_id từ response

# Tạo restaurant
curl -X POST http://localhost:8080/v1/restaurants \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name":"Pho Saigon","addr":"123 Nguyen Hue St","category_id":"YOUR_CATEGORY_ID"}'
# → Lưu restaurant_id từ response
```

## Food API Endpoints

### 1. List Foods (Public)
```bash
curl -X GET "http://localhost:8080/v1/foods?page=1&limit=20"

# Expected Response:
{
  "data": {
    "data": [],
    "page": 1,
    "limit": 20,
    "total": 0
  }
}
```

### 2. Create Food (Authenticated) - **WITH VALIDATION**
```bash
curl -X POST http://localhost:8080/v1/foods \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "name": "Pho Bo",
    "description": "Traditional beef noodle soup",
    "price": 25.5,
    "category_id": "YOUR_CATEGORY_ID",
    "restaurant_id": "YOUR_RESTAURANT_ID"
  }'

# Expected Response (201):
{
  "data": {
    "id": "01988d94-...",
    "restaurant_id": "01988d94-...",
    "category_id": "01988d67-...",
    "name": "Pho Bo",
    "description": "Traditional beef noodle soup",
    "price": 25.5,
    "status": "active",
    "created_at": "2025-08-09T06:50:00Z",
    "updated_at": "2025-08-09T06:50:00Z"
  }
}
```

#### **Create Food Flow & Validation**

The create food API now includes **comprehensive validation** to ensure data integrity:

**Validation Flow:**
1. **Basic DTO Validation**
   - `name` is required and not empty
   - `restaurant_id` is required and not empty  
   - `price` must be greater than 0

2. **UUID Format Validation**
   - `restaurant_id` must be valid UUID format
   - `category_id` must be valid UUID format (if provided)

3. **🆕 Restaurant Existence Validation (via gRPC)**
   - **gRPC Call**: `Restaurant.GetRestaurantsByIDs` (port 6003)
   - **Tracing**: OpenTelemetry span `restaurant-grpc.validate-exists`
   - **Purpose**: Ensures restaurant_id exists and is accessible

4. **🆕 Category Existence Validation (via gRPC)**
   - **gRPC Call**: `Category.GetCategoriesByIDs` (port 6000) 
   - **Tracing**: OpenTelemetry span `category-grpc.validate-exists`
   - **Purpose**: Ensures category_id exists and is accessible (if provided)

5. **Food Entity Creation & Database Insert**
   - **Tracing**: OpenTelemetry span `food-repo.insert`
   - **Database**: MySQL insert with GORM

**Observability:**
- Main span: `food-service.create`
- Restaurant validation: `restaurant-grpc.validate-exists`  
- Category validation: `category-grpc.validate-exists`
- Database insert: `food-repo.insert`

### 3. Get Food by ID (Public, Enhanced Response)
```bash
curl -X GET http://localhost:8080/v1/foods/YOUR_FOOD_ID

# Expected Response với thông tin category và restaurant:
{
  "data": {
    "id": "01988d94-...",
    "restaurant_id": "01988d94-...",
    "category_id": "01988d67-...",
    "name": "Pho Bo",
    "description": "Traditional beef noodle soup",
    "price": 25.5,
    "status": "active",
    "created_at": "2025-08-09T06:50:00Z",
    "updated_at": "2025-08-09T06:50:00Z",
    "category": {
      "id": "01988d67-...",
      "name": "Vietnamese Food",
      "status": "active"
    },
    "restaurant": {
      "id": "01988d94-...",
      "name": "Pho Saigon",
      "address": "123 Nguyen Hue St",
      "category_id": "01988d67-...",
      "status": "active"
    }
  }
}
```

### 4. Update Food (Authenticated)
```bash
curl -X PATCH http://localhost:8080/v1/foods/YOUR_FOOD_ID \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "name": "Pho Bo Special",
    "price": 30.0
  }'

# Expected Response (200): Updated food object
```

### 5. Delete Food (Authenticated) - **SOFT DELETE**
```bash
curl -X DELETE http://localhost:8080/v1/foods/YOUR_FOOD_ID \
  -H "Authorization: Bearer $TOKEN"

# Expected Response (200):
{
  "data": {
    "message": "Food deleted successfully"
  }
}
```

**Implementation Details:**
- **Type**: Soft Delete (status-based)
- **Action**: Updates `status` from `"active"` → `"deleted"`
- **Data Preservation**: Food record remains in database
- **Validation**: Prevents double deletion (returns 400 if already deleted)
- **Recovery**: Possible by updating status back to `"active"`

### 6. Filter Foods by Category/Restaurant
```bash
# Filter by category
curl -X GET "http://localhost:8080/v1/foods?category_id=YOUR_CATEGORY_ID"

# Filter by restaurant  
curl -X GET "http://localhost:8080/v1/foods?restaurant_id=YOUR_RESTAURANT_ID"

# Filter by price range
curl -X GET "http://localhost:8080/v1/foods?min_price=20&max_price=50"
```

## Testing Error Cases

### 1. Validation Errors (400)
```bash
# Missing required fields
curl -X POST http://localhost:8080/v1/foods \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name": ""}'
# → 400 Bad Request

# Invalid price
curl -X POST http://localhost:8080/v1/foods \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name": "Test", "price": -5}'
# → 400 Bad Request

# Invalid UUID format
curl -X POST http://localhost:8080/v1/foods \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name": "Test", "price": 25.0, "restaurant_id": "invalid-uuid"}'
# → 400 Bad Request: "invalid restaurant_id format"

# 🆕 Non-existent restaurant_id
curl -X POST http://localhost:8080/v1/foods \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name": "Test", "price": 25.0, "restaurant_id": "01988d94-1234-1234-1234-123456789012"}'
# → 400 Bad Request: "restaurant_id does not exist or is not accessible"

# 🆕 Non-existent category_id  
curl -X POST http://localhost:8080/v1/foods \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "name": "Test", 
    "price": 25.0, 
    "restaurant_id": "YOUR_VALID_RESTAURANT_ID",
    "category_id": "01988d94-1234-1234-1234-123456789012"
  }'
# → 400 Bad Request: "category_id does not exist or is not accessible"
```

### 2. Authentication Errors (401)
```bash
# No token
curl -X POST http://localhost:8080/v1/foods \
  -H "Content-Type: application/json" \
  -d '{"name": "Test"}'
# → 401 Unauthorized

# Invalid token
curl -X POST http://localhost:8080/v1/foods \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer invalid-token" \
  -d '{"name": "Test"}'
# → 401 Unauthorized
```

### 3. Not Found Errors (404)
```bash
# Get non-existent food
curl -X GET http://localhost:8080/v1/foods/non-existent-id
# → 404 Not Found
```

## gRPC Services (Port 6001)

Food module cũng cung cấp gRPC services cho inter-service communication:

- `GetFoodsByIDs` - Batch lookup foods
- `GetFoodsByRestaurantID` - Get foods by restaurant
- `GetFoodsByCategoryID` - Get foods by category

## Expected Status Codes

- `200` - OK (GET, PATCH, DELETE successful)
- `201` - Created (POST successful) 
- `400` - Bad Request (validation errors, malformed JSON)
- `401` - Unauthorized (missing/invalid token)
- `404` - Not Found (resource doesn't exist)
- `500` - Internal Server Error (unexpected server errors)

## Architecture Features

- ✅ **Enhanced Response**: Food API tự động load thông tin category + restaurant via gRPC
- ✅ **Pure gRPC Communication**: Zero HTTP calls giữa services  
- ✅ **JWT Authentication**: Token validation qua User gRPC service
- ✅ **🆕 Data Integrity Validation**: Restaurant & Category existence validation via gRPC
- ✅ **OpenTelemetry Tracing**: Distributed tracing support với validation spans
- ✅ **Hexagonal Architecture**: Clean separation of concerns

## Notes

1. **Authentication**: POST, PATCH, DELETE operations yêu cầu JWT token
2. **Enhanced Response**: GET /v1/foods/:id trả về thông tin đầy đủ với category + restaurant
3. **🆕 Data Validation**: POST /v1/foods validates restaurant_id & category_id existence via gRPC
4. **gRPC Communication**: Food service gọi Category và Restaurant services qua gRPC
5. **Error Handling**: Status codes chính xác, phân biệt client vs server errors
6. **Tracing**: Tất cả requests được trace qua OpenTelemetry (bao gồm validation spans)
7. **🆕 Referential Integrity**: Đảm bảo food chỉ được tạo với restaurant & category tồn tại

Happy testing! 🚀