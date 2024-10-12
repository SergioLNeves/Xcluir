package repository

import (
	"errors"
	"fmt"
	"github.com/SergioLNeves/Xcluir/config"
	"github.com/SergioLNeves/Xcluir/domain"
	"github.com/dghubble/oauth1"
	"net/http"
)

type tweetRepository struct{}

func NewTweetRepository() domain.TweetRepository {
	return &tweetRepository{}
}
func (r *tweetRepository) DeleteTweet(tweetID string) error {
	consumerKey, consumerSecret, accessToken, accessSecret, err := config.GetTwitterCredentials()
	if err != nil {
		return err
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	url := fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}
