package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ListLaws(c *gin.Context) {
	// TODO: 查询法条库
	laws := []map[string]any{
		{"article": "著作权法第10条", "description": "..."},
	}
	c.JSON(http.StatusOK, gin.H{"laws": laws})
}