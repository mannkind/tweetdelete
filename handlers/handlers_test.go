package handlers

import (
	"bytes"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/viper"
	"net/http"
	"testing"
	"time"
)

func readTestConfig(t *testing.T) {
	testConfig := []byte(`
consumer_key: "YOUR CONSUMER KEY"
consumer_secret: "YOUR CONSUMER SECRET"
oauth_token: "YOUR OAUTH TOKEN"
oauth_token_secret: "YOUR OAUTH TOKEN SECRET"
timeline_count: 50
max_age: 72
save:
  - "123456789012345678"
`)
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(testConfig)); err != nil {
		t.Error(err)
	}
}

func TestNewTwitterClient(t *testing.T) {
	readTestConfig(t)

	consumerKey := viper.GetString("consumer_key")
	consumerSecret := viper.GetString("consumer_secret")
	oauthToken := viper.GetString("oauth_token")
	oauthTokenSecret := viper.GetString("oauth_token_secret")

	client := NewTwitterClient(consumerKey, consumerSecret, oauthToken, oauthTokenSecret)
	if client == nil {
		t.Error("Unable to create Twitter client")
	}
}

func TestGetTweets(t *testing.T) {
	readTestConfig(t)

	timeline := tdTimelineClientTest{}
	timelineCount := viper.GetInt("timeline_count")
	savedIDStrs := viper.GetStringSlice("save")
	tweets := GetTweets(timeline, timelineCount, savedIDStrs)

	for _, tweet := range tweets {
		if tweet.IDStr == "123456789012345678" {
			t.Error("Tweet ID should have been removed.")
		}
	}
}

func TestDeleteTweets(t *testing.T) {
	maxAge := float64(24)
	statuses := tdStatusesClientTest{}
	tweets := []twitter.Tweet{
		{IDStr: "123456789012345677", Text: "Plz Delete this Tweet", CreatedAt: time.Now().Add(-25 * time.Hour).Format(rfc2822)},
		{IDStr: "123456789012345679", Text: "Plz Keep this Tweet", CreatedAt: time.Now().Add(-1 * time.Hour).Format(rfc2822)},
	}

	DeleteTweets(statuses, tweets, maxAge)
}

type tdTimelineClientTest struct {
}

func (t tdTimelineClientTest) UserTimeline(params *twitter.UserTimelineParams) ([]twitter.Tweet, *http.Response, error) {
	return []twitter.Tweet{
		{IDStr: "123456789012345677", Text: "Plz Ignore this Tweet"},
		{IDStr: "123456789012345678", Text: "Plz Save this Tweet"},
	}, nil, nil
}

type tdStatusesClientTest struct {
}

func (t tdStatusesClientTest) Destroy(id int64, params *twitter.StatusDestroyParams) (*twitter.Tweet, *http.Response, error) {
	return nil, nil, nil
}
