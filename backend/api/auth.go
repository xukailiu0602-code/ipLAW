package api

import (
	"net/http"
	"time"
	"backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("iprisk-secret")

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// TODO: 查询用户并校验密码
	user := models.User{ID: "u1", Username: req.Username, Role: "admin"}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role": user.Role,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString(jwtSecret)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// TODO: 写入用户表
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func GetMe(c *gin.Context) {
	userID := c.GetString("user_id")
	role := c.GetString("role")
	c.JSON(http.StatusOK, gin.H{"user_id": userID, "role": role})
}