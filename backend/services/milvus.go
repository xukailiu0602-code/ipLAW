package services

import (
	"context"
	"log"
	"backend/config"
	milvus "github.com/milvus-io/milvus-sdk-go/v2/client"
)

type MilvusClient struct {
	Client milvus.Client
}

func InitMilvus(cfg *config.Config) *MilvusClient {
	cli, err := milvus.NewGrpcClient(context.Background(), cfg.MilvusHost+":"+string(rune(cfg.MilvusPort)))
	if err != nil {
		log.Fatalf("Milvus connect error: %v", err)
	}
	return &MilvusClient{Client: cli}
}

func (m *MilvusClient) Close() {
	if m.Client != nil {
		_ = m.Client.Close()
	}
}