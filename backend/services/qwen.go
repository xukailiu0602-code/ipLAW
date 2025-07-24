package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"backend/config"
)

type QwenClient struct {
	ApiKey  string
	BaseUrl string
}

func NewQwenClient(cfg *config.Config) *QwenClient {
	return &QwenClient{
		ApiKey:  cfg.QwenApiKey,
		BaseUrl: cfg.QwenBaseUrl,
	}
}

func (q *QwenClient) Embedding(text string) ([]float32, error) {
	// 伪代码，实际需按 Qwen3-Embedding API 文档实现
	url := q.BaseUrl + "/embedding"
	body := map[string]string{"text": text}
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(b))
	req.Header.Set("Authorization", "Bearer "+q.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil { return nil, err }
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	// 假设返回 {"embedding": [0.1,0.2,...]}
	var res struct{ Embedding []float32 `json:"embedding"` }
	_ = json.Unmarshal(data, &res)
	return res.Embedding, nil
}

func (q *QwenClient) Rerank(query string, docs []string) ([]int, error) {
	// 伪代码，实际需按 Qwen3-Reranker API 文档实现
	return []int{0,1,2}, nil
}

func (q *QwenClient) LLMAnswer(payload map[string]any) (map[string]any, error) {
	// 伪代码，实际需按 Qwen-Plus API 文档实现
	url := q.BaseUrl + "/llm"
	b, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(b))
	req.Header.Set("Authorization", "Bearer "+q.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil { return nil, err }
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	var res map[string]any
	_ = json.Unmarshal(data, &res)
	return res, nil
}