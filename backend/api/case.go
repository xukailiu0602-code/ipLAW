package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ListCases(c *gin.Context) {
	// TODO: 从向量库检索类案
	cases := []map[string]any{
		{"caseId": "2022XYZ", "court": "北京知识产权法院", "year": 2023, "summary": "...", "similarity": 0.85},
	}
	c.JSON(http.StatusOK, gin.H{"cases": cases})
}