package api

import (
	"net/http"
	"time"
	"backend/models"
	"github.com/gin-gonic/gin"
)

func UploadDocument(c *gin.Context) {
	userID := c.GetString("user_id")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file"})
		return
	}
	// TODO: 保存文件、抽取文本、切片、embedding 入 Milvus
	doc := models.Document{
		ID:        "d1",
		UserID:    userID,
		Title:     file.Filename,
		Type:      "pdf",
		CreatedAt: time.Now().Unix(),
	}
	// TODO: 存入 MongoDB
	c.JSON(http.StatusOK, gin.H{"doc_id": doc.ID})
}