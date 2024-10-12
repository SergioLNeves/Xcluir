package repository

import (
	"errors"
	"fmt"
	"github.com/SergioLNeves/Xcluir/config"
	"github.com/SergioLNeves/Xcluir/domain"
	"net/http"
)

type tweetRepository struct{}

func NewTweetRepository() domain.TweetRepository {
	return &tweetRepository{}
}
func (r *tweetRepository) DeleteTweet(tweetID string) error {
	bearerToken, err := config.GetTwitterBearerToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	return nil

}
