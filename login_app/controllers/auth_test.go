package controllers

import (
	"bytes"
	"encoding/json"
	"login-app/database"
	"login-app/models"
	"login-app/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Khởi tạo database trước khi chạy test
func setup() {
	database.Init()
	// Xoá dữ liệu test nếu cần thiết
	database.DB.Exec("DELETE FROM users WHERE username LIKE 'testuser%'")
}

func TestSignUp_Success(t *testing.T) {
	setup()
	gin.SetMode(gin.TestMode)

	user := models.User{
		Username: "testuser1",
		Password: "123456",
	}

	body, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	SignUp(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Đăng ký thành công", resp["message"])
}

func TestSignUp_Fail_DuplicateUser(t *testing.T) {
	setup()
	gin.SetMode(gin.TestMode)

	// Đăng ký user mới
	user := models.User{
		Username: "testuser2",
		Password: "123456",
	}
	database.DB.Create(&user)

	body, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	SignUp(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Contains(t, resp["error"], "Tài khoản đã tồn tại")
}

func TestSignIn_Success(t *testing.T) {
	setup()
	gin.SetMode(gin.TestMode)

	// Tạo user test trong db với mật khẩu đã hash
	user := models.User{
		Username: "testuser3",
		Password: "123456",
	}
	hashed, _ := utils.HashPassword(user.Password)
	user.Password = hashed
	database.DB.Create(&user)

	creds := map[string]string{
		"username": user.Username,
		"password": "123456",
	}
	body, _ := json.Marshal(creds)
	req, _ := http.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	SignIn(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NotEmpty(t, resp["access_token"])
	assert.NotEmpty(t, resp["refresh_token"])
}

func TestSignIn_Fail_InvalidCredentials(t *testing.T) {
	setup()
	gin.SetMode(gin.TestMode)

	creds := map[string]string{
		"username": "nonexistent",
		"password": "wrongpass",
	}
	body, _ := json.Marshal(creds)
	req, _ := http.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	SignIn(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Sai tài khoản hoặc mật khẩu", resp["error"])
}

func TestSignIn_Fail_BadRequest(t *testing.T) {
	setup()
	gin.SetMode(gin.TestMode)

	body := []byte(`invalid-json`)
	req, _ := http.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	SignIn(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Dữ liệu đầu vào không hợp lệ", resp["error"])
}

func TestRefresh_Success(t *testing.T) {
	setup()
	gin.SetMode(gin.TestMode)

	// Tạo token hợp lệ
	token, _ := utils.GenerateRefreshToken("testuser3")

	reqBody := map[string]string{"token": token}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/refresh", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	Refresh(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NotEmpty(t, resp["access_token"])
}

func TestRefresh_Fail_InvalidToken(t *testing.T) {
	setup()
	gin.SetMode(gin.TestMode)

	reqBody := map[string]string{"token": "invalid.token.here"}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/refresh", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	Refresh(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Token không hợp lệ", resp["error"])
}
