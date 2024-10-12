package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() error {
	return godotenv.Load()
}

func GetTwitterCredentials() (string, string, string, string, error) {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessTokenSecret == "" {
		return "", "", "", "", fmt.Errorf("one or more Twitter credentials are missing")
	}

	return consumerKey, consumerSecret, accessToken, accessTokenSecret, nil
}
