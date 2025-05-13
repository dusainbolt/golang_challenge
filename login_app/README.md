
# 🛡️ Golang Login App

Ứng dụng Golang mẫu với tính năng xác thực người dùng bằng **JWT**, gồm:

- ✅ Đăng ký người dùng (Sign Up)
- ✅ Đăng nhập (Sign In)
- ✅ Refresh Token

---

## ⚙️ Yêu cầu cài đặt (Prerequisites)

- [Go](https://go.dev/dl/) >= 1.18
- [Git](https://git-scm.com/)
- Hệ điều hành: Linux / macOS / Windows
- (Tùy chọn) cài `curl` để test nhanh API

---

## 📦 Cài đặt & Chạy thử

```bash
# Clone source code
git clone https://github.com/your-username/login-app.git
cd login-app

# Khởi tạo module Go
go mod tidy

# Chạy ứng dụng
go run main.go
```

Ứng dụng sẽ chạy tại: `http://localhost:8080`

---

## 🔐 JWT Secret Key

Secret key được định nghĩa ở `utils/jwt.go`. Bạn có thể đổi thành biến môi trường để bảo mật hơn.

```go
var JwtKey = []byte("g14L7aUJlG1NDK3mAdrxu6ZxywEqSPgLMkdDwK8oRmg=")
```

---

## 🧪 Test nhanh bằng `curl`

### ✅ 1. Đăng ký tài khoản

```bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser", "password":"123456"}'
```

### ✅ 2. Đăng nhập

```bash
curl -X POST http://localhost:8080/signin \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser", "password":"123456"}'
```

🔁 Response ví dụ:
```json
{
  "access_token": "xxxxx.yyyyy.zzzzz",
  "refresh_token": "aaaaa.bbbbb.ccccc"
}
```

### 🔄 3. Refresh Token

```bash
curl -X POST http://localhost:8080/refresh \
  -H "Content-Type: application/json" \
  -d '{"token":"<refresh_token_here>"}'
```

---

## 📂 Cấu trúc thư mục

```
.
├── controllers/      # Xử lý các API auth
├── database/         # Kết nối DB bằng GORM
├── middleware/       # JWT middleware (nếu có)
├── models/           # User struct
├── routes/           # Khai báo các route
├── utils/            # Hàm tiện ích (JWT, hash)
├── main.go           # Entry point
└── go.mod
```

---

## 📌 Ghi chú

- Mặc định sử dụng **SQLite** cho phát triển nhanh
- Có thể đổi sang PostgreSQL/MySQL dễ dàng với GORM driver