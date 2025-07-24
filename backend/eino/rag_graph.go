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
		// 1. 获取 query embedding
		// 2. Milvus 召回 top-50
		return input, nil
	}
}

func RerankerNode(qwen *services.QwenClient) NodeFunc {
	return func(ctx context.Context, input map[string]any) (map[string]any, error) {
		// 1. Qwen3-Reranker 精排
		return input, nil
	}
}

func LLMAnswerNode(qwen *services.QwenClient) NodeFunc {
	return func(ctx context.Context, input map[string]any) (map[string]any, error) {
		// 1. Qwen-Plus 生成
		return input, nil
	}
}

func PostProcessNode() NodeFunc {
	return func(ctx context.Context, input map[string]any) (map[string]any, error) {
		// 1. 风险评分、类案检索、结构化输出
		return input, nil
	}
}