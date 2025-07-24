package api

import (
	"context"
	"net/http"
	"backend/eino"
	"github.com/gin-gonic/gin"
)

func Ask(c *gin.Context) {
	userID := c.GetString("user_id")
	var req struct {
		Query   string   `json:"query"`
		DocIDs  []string `json:"doc_ids"`
		History []any    `json:"history"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// TODO: 构造 Eino RAG 编排输入
	graph := eino.NewGraph()
	// TODO: 注册节点
	// TODO: 传入 Qwen/Milvus/Mongo 依赖
	graph.AddNode("embed_retrieve", nil)
	graph.AddNode("rerank", nil)
	graph.AddNode("llm_answer", nil)
	graph.AddNode("post_process", nil)
	graph.AddEdge("embed_retrieve", "rerank")
	graph.AddEdge("rerank", "llm_answer")
	graph.AddEdge("llm_answer", "post_process")
	graph.AddEdge("post_process", "END")
	input := map[string]any{"query": req.Query, "user_id": userID, "history": req.History}
	res, err := graph.Run(context.Background(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}