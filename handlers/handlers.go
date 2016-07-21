package handlers

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"net/http"
	"time"
)

type tdTimelineClient interface {
	UserTimeline(params *twitter.UserTimelineParams) ([]twitter.Tweet, *http.Response, error)
}

type tdStatusesClient interface {
	Destroy(id int64, params *twitter.StatusDestroyParams) (*twitter.Tweet, *http.Response, error)
}

const sleepCount = 2
const rfc2822 = "Mon Jan 02 15:04:05 -0700 2006"

// NewTwitterClient - Creates a new Twitter client given the oauth tokens and such
func NewTwitterClient(consumerKey, consumerSecret, oauthToken, oauthTokenSecret string) *twitter.Client {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(oauthToken, oauthTokenSecret)

	return twitter.NewClient(config.Client(oauth1.NoContext, token))
}

// GetTweets - Gets a count of tweets for the current user
func GetTweets(timeline tdTimelineClient, count int, saveIDStrs []string) []twitter.Tweet {
	timelineParams := &twitter.UserTimelineParams{Count: count}
	tweets, _, _ := timeline.UserTimeline(timelineParams)
	var deleteTweets []twitter.Tweet

NextTweet:
	for _, tweet := range tweets {
		for _, saveIDStr := range saveIDStrs {
			if tweet.IDStr == saveIDStr {
				log.Printf("Tweet ID: %s... ", tweet.IDStr)
				log.Printf(" - Text: %s", tweet.Text)
				log.Printf(" - Status: Saved")
				continue NextTweet
			}
		}

		deleteTweets = append(deleteTweets, tweet)
	}

	return deleteTweets
}

// DeleteTweets - Delete all non-saved tweets older than a specified max age
func DeleteTweets(statuses tdStatusesClient, tweets []twitter.Tweet, maxAge float64) {
	for _, tweet := range tweets {
		log.Printf("Checking %s...", tweet.IDStr)
		log.Printf(" - Text: %s", tweet.Text)

		created, _ := time.Parse(rfc2822, tweet.CreatedAt)
		age := time.Since(created).Hours()
		if age >= maxAge {
			log.Printf(" - Status: Deleted")
			statuses.Destroy(tweet.ID, nil)
		} else {
			log.Printf(" - Status: Kept for %f more hours", maxAge-age)
		}
		log.Println("")
		time.Sleep(sleepCount * time.Second)
	}
}
