package eino

import (
	"context"
	"backend/services"
)

type EvalResult map[string]any

type NodeFunc func(ctx context.Context, input map[string]any) (map[string]any, error)

type Graph struct {
	Nodes map[string]NodeFunc
	Edges map[string]string
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]NodeFunc),
		Edges: make(map[string]string),
	}
}

func (g *Graph) AddNode(name string, fn NodeFunc) {
	g.Nodes[name] = fn
}

func (g *Graph) AddEdge(from, to string) {
	g.Edges[from] = to
}

func (g *Graph) Run(ctx context.Context, input map[string]any) (map[string]any, error) {
	cur := "embed_retrieve"
	data := input
	var err error
	for cur != "END" {
		fn := g.Nodes[cur]
		data, err = fn(ctx, data)
		if err != nil { return nil, err }
		cur = g.Edges[cur]
	}
	return data, nil
}

// 节点实现示例
func RetrieverNode(qwen *services.QwenClient, milvus *services.MilvusClient) NodeFunc {
	return func(ctx context.Context, input map[string]any) (map[string]any, error) {
		query, _ := input["query"].(string)
		// 1. 获取 query embedding
		emb, err := qwen.Embedding(query, 1024)
		if err != nil {
			return nil, err
		}
		// 2. Milvus 召回 top-50（伪代码，假设有Search接口）
		// docs, docIDs := milvus.Search(emb, 50)
		// 这里用占位
		docs := []string{"文档1内容", "文档2内容", "文档3内容"}
		input["docs"] = docs
		return input, nil
	}
}

func RerankerNode(qwen *services.QwenClient) NodeFunc {
	return func(ctx context.Context, input map[string]any) (map[string]any, error) {
		query, _ := input["query"].(string)
		docs, _ := input["docs"].([]string)
		if len(docs) == 0 {
			return input, nil
		}
		indices, err := qwen.Rerank(query, docs, 5)
		if err != nil {
			return nil, err
		}
		// 取前5个文档
		var topDocs []string
		for _, idx := range indices {
			if idx < len(docs) {
				topDocs = append(topDocs, docs[idx])
			}
		}
		input["top_docs"] = topDocs
		return input, nil
	}
}

func LLMAnswerNode(qwen *services.QwenClient) NodeFunc {
	return func(ctx context.Context, input map[string]any) (map[string]any, error) {
		query, _ := input["query"].(string)
		topDocs, _ := input["top_docs"].([]string)
		history, _ := input["history"].([]any)
		// 构造messages
		messages := []map[string]string{
			{"role": "system", "content": "你是一个专业的法律助手。"},
			{"role": "user", "content": query},
		}
		if len(topDocs) > 0 {
			contextStr := "检索到的相关文档如下：\n" + topDocs[0]
			messages = append(messages, map[string]string{"role": "system", "content": contextStr})
		}
		answer, err := qwen.LLMAnswer(messages, map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		input["answer"] = answer
		return input, nil
	}
}

func PostProcessNode() NodeFunc {
	return func(ctx context.Context, input map[string]any) (map[string]any, error) {
		// 1. 风险评分、类案检索、结构化输出
		return input, nil
	}
}