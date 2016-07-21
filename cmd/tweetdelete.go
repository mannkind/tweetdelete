package cmd

import (
	"github.com/mannkind/tweetdelete/handlers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

const version string = "0.1.0"

var cfgFile string

// TweetDeleteCmd - The root command
var TweetDeleteCmd = &cobra.Command{
	Use:   "tweetdelete",
	Short: "Delete Tweets",
	Long:  "Delete Tweets",
	Run: func(cmd *cobra.Command, args []string) {
		consumerKey := viper.GetString("consumer_key")
		consumerSecret := viper.GetString("consumer_secret")
		oauthToken := viper.GetString("oauth_token")
		oauthTokenSecret := viper.GetString("oauth_token_secret")
		client := handlers.NewTwitterClient(consumerKey, consumerSecret, oauthToken, oauthTokenSecret)

		timelineCount := viper.GetInt("timeline_count")
		savedIDStrs := viper.GetStringSlice("save")
		tweets := handlers.GetTweets(client.Timelines, timelineCount, savedIDStrs)

		maxAge := viper.GetFloat64("max_age")
		handlers.DeleteTweets(client.Statuses, tweets, maxAge)
	},
}

// Execute - Adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := TweetDeleteCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.Printf("TweetDelete Version: %s", version)

	cobra.OnInitialize(func() {
		viper.SetConfigFile(cfgFile)
		viper.SetDefault("timeline_count", 25)
		viper.SetDefault("max_age", 72)
		log.Printf("Loading Configuration %s", cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error Loading Configuration: %s ", err)
		}
		log.Printf("Loaded Configuration %s", cfgFile)
	})

	TweetDeleteCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", ".tweetdelete.yaml", "The path to the configuration file")
}
