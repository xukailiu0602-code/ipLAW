package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AdminModelConfig(c *gin.Context) {
	// TODO: 模型配置
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func AdminSliceRule(c *gin.Context) {
	// TODO: 切片规则管理
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func AdminReindex(c *gin.Context) {
	// TODO: 索引重建
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func AdminLogs(c *gin.Context) {
	// TODO: 日志监控
	c.JSON(http.StatusOK, gin.H{"logs": []string{"log1", "log2"}})
}

func AdminListUsers(c *gin.Context) {
	// TODO: 用户列表
	c.JSON(http.StatusOK, gin.H{"users": []string{"u1", "u2"}})
}

func AdminSetRole(c *gin.Context) {
	// TODO: 设置用户角色
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}