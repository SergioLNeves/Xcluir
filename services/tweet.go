package services

import (
	"encoding/json"
	"fmt"
	"github.com/SergioLNeves/Xcluir/domain"
	"io/ioutil"
	"os"
	"regexp"
	"time"
)

type tweetService struct {
	tweetRepo domain.TweetRepository
}

func NewTweetServices(tweetRepo domain.TweetRepository) domain.TweetService {
	return &tweetService{tweetRepo}
}

func (s *tweetService) DeleteTweetsFromFile(filename string) error {

	tweets, err := loadTweetsFromFile(filename)
	if err != nil {
		return fmt.Errorf("failed to loading tweets from file: %v", err)
	}
	for _, tweet := range tweets {
		err = s.tweetRepo.DeleteTweet(tweet.ID)
		if err != nil {
			return fmt.Errorf("failed to delete tweet %d: %v", tweet.ID, err)
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func loadTweetsFromFile(filename string) ([]domain.Tweet, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	// Remove a parte JavaScript antes do JSON
	re := regexp.MustCompile(`window\.YTD\.tweets\.part0\s+=\s+`)
	jsonData := re.ReplaceAll(byteValue, nil)

	var tweetsFile []map[string]map[string]interface{}
	if err := json.Unmarshal(jsonData, &tweetsFile); err != nil {
		return nil, err
	}

	var tweets []domain.Tweet
	for _, tweetMap := range tweetsFile {
		tweetData := tweetMap["tweet"]
		id := tweetData["id_str"].(string)

		tweet := domain.Tweet{
			ID: id,
		}
		tweets = append(tweets, tweet)
	}

	return tweets, nil
}
