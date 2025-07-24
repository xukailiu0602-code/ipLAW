package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ExportReport(c *gin.Context) {
	id := c.Param("id")
	// TODO: 查询问答记录，生成 PDF/Word
	c.Header("Content-Disposition", "attachment; filename=report.pdf")
	c.Data(http.StatusOK, "application/pdf", []byte("PDF DATA"))
}