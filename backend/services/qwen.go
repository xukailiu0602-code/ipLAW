package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"backend/config"
	"os"
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

func (q *QwenClient) Embedding(texts interface{}, dimensions int) ([]float32, error) {
	// texts: string 或 []string
	url := q.BaseUrl + "/embeddings"
	body := map[string]interface{}{
		"model": "text-embedding-v4",
		"input": texts,
		"dimensions": dimensions,
		"encoding_format": "float",
	}
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(b))
	req.Header.Set("Authorization", "Bearer "+q.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil { return nil, err }
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	var res struct {
		Data []struct {
			Embedding []float32 `json:"embedding"`
			Index     int       `json:"index"`
			Object    string    `json:"object"`
		} `json:"data"`
	}
	err = json.Unmarshal(data, &res)
	if err != nil || len(res.Data) == 0 {
		return nil, fmt.Errorf("embedding failed: %v, resp: %s", err, string(data))
	}
	return res.Data[0].Embedding, nil
}

func (q *QwenClient) Rerank(query string, docs []string, topN int) ([]int, error) {
	url := "https://dashscope.aliyuncs.com/api/v1/services/rerank/text-rerank/text-rerank"
	body := map[string]interface{}{
		"model": "gte-rerank-v2",
		"input": map[string]interface{}{
			"query": query,
			"documents": docs,
		},
		"parameters": map[string]interface{}{
			"return_documents": true,
			"top_n": topN,
		},
	}
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(b))
	req.Header.Set("Authorization", "Bearer "+q.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil { return nil, err }
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	var res struct {
		Output struct {
			Results []struct {
				Index int `json:"index"`
			} `json:"results"`
		} `json:"output"`
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("rerank failed: %v, resp: %s", err, string(data))
	}
	indices := make([]int, 0, len(res.Output.Results))
	for _, r := range res.Output.Results {
		indices = append(indices, r.Index)
	}
	return indices, nil
}

func (q *QwenClient) LLMAnswer(messages []map[string]string, extra map[string]interface{}) (string, error) {
	url := q.BaseUrl + "/chat/completions"
	body := map[string]interface{}{
		"model": "qwen-plus",
		"messages": messages,
	}
	for k, v := range extra {
		body[k] = v
	}
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(b))
	req.Header.Set("Authorization", "Bearer "+q.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil { return "", err }
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	var res struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	err = json.Unmarshal(data, &res)
	if err != nil || len(res.Choices) == 0 {
		return "", fmt.Errorf("llm answer failed: %v, resp: %s", err, string(data))
	}
	return res.Choices[0].Message.Content, nil
}