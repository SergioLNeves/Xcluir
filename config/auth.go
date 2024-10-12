package config

import (
	"fmt"
	"os"
)

func GetTwitterBearerToken() (string, error) {
	token := os.Getenv("TWITTER_BEARER_TOKEN")
	if token == "" {
		return "", fmt.Errorf("TWITTER_BEARER_TOKEN não configurado nas variáveis de ambiente")
	}
	return "Bearer " + token, nil
}
