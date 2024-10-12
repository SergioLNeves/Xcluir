package domain

type Tweet struct {
	ID string `json:"id_str"`
}

type TweetFile struct {
	Tweets []Tweet `json:"tweets"`
}

type TweetRepository interface {
	DeleteTweet(tweetID string) error
}

type TweetService interface {
	DeleteTweetsFromFile(filename string) error
}
