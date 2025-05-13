package controllers

import (
	"login-app/database"
	"login-app/models"
	"login-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func SignUp(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	hashed, _ := utils.HashPassword(user.Password)
	user.Password = hashed
	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tài khoản đã tồn tại"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Đăng ký thành công"})
}

func SignIn(c *gin.Context) {
	var input models.User
	var user models.User
	c.BindJSON(&input)

	result := database.DB.Where("username = ?", input.Username).First(&user)
	if result.Error != nil || !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Sai tài khoản hoặc mật khẩu"})
		return
	}

	token, _ := utils.GenerateToken(user.Username)
	refreshToken, _ := utils.GenerateRefreshToken(user.Username)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  token,
		"refresh_token": refreshToken,
	})
}

func Refresh(c *gin.Context) {
	type RefreshRequest struct {
		Token string `json:"token"`
	}
	var req RefreshRequest
	c.BindJSON(&req)

	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		return utils.JwtKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		newToken, _ := utils.GenerateToken(username)
		c.JSON(http.StatusOK, gin.H{"access_token": newToken})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ", "err": err})
	}
}
