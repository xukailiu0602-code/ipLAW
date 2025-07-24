package api

import (
	"context"
	"net/http"
	"backend/eino"
	"github.com/gin-gonic/gin"
	"backend/services"
	"backend/config"
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
	cfg := config.LoadConfig()
	qwen := services.NewQwenClient(cfg)
	milvus := services.InitMilvus(cfg)
	graph := eino.NewGraph()
	graph.AddNode("embed_retrieve", eino.RetrieverNode(qwen, milvus))
	graph.AddNode("rerank", eino.RerankerNode(qwen))
	graph.AddNode("llm_answer", eino.LLMAnswerNode(qwen))
	graph.AddNode("post_process", eino.PostProcessNode())
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