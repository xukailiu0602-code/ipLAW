package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MongoURI    string
	MongoDB     string
	MilvusHost  string
	MilvusPort  int
	MilvusUser  string
	MilvusPass  string
	QwenApiKey  string
	QwenBaseUrl string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
	return &Config{
		MongoURI:    viper.GetString("MONGO_URI"),
		MongoDB:     viper.GetString("MONGO_DB"),
		MilvusHost:  viper.GetString("MILVUS_HOST"),
		MilvusPort:  viper.GetInt("MILVUS_PORT"),
		MilvusUser:  viper.GetString("MILVUS_USER"),
		MilvusPass:  viper.GetString("MILVUS_PASS"),
		QwenApiKey:  viper.GetString("QWEN_API_KEY"),
		QwenBaseUrl: viper.GetString("QWEN_BASE_URL"),
	}
}