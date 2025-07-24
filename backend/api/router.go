package api

import (
	"backend/api"
	"backend/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(r *gin.Engine, mongoClient *mongo.Client) {
	apiGroup := r.Group("/api")

	apiGroup.POST("/login", api.Login)
	apiGroup.POST("/register", api.Register)

	apiGroup.Use(middleware.JWTAuth())
	{
		apiGroup.GET("/user/me", api.GetMe)
		apiGroup.POST("/upload", api.UploadDocument)
		apiGroup.POST("/ask", api.Ask)
		apiGroup.GET("/cases", api.ListCases)
		apiGroup.GET("/laws", api.ListLaws)
		apiGroup.GET("/report/:id", api.ExportReport)
		// 管理后台
		apiGroup.POST("/admin/model", api.AdminModelConfig)
		apiGroup.POST("/admin/slice", api.AdminSliceRule)
		apiGroup.POST("/admin/reindex", api.AdminReindex)
		apiGroup.GET("/admin/logs", api.AdminLogs)
		apiGroup.GET("/admin/users", api.AdminListUsers)
		apiGroup.POST("/admin/role", api.AdminSetRole)
	}
}