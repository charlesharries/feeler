package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/charlesharries/feeler/pkg/afinn"
	"github.com/charlesharries/feeler/pkg/sentiment"
	"github.com/charlesharries/feeler/pkg/twitter"
	"github.com/joho/godotenv"
)

func main() {
	limit := flag.Int("limit", 100, "Number of tweets to scan")
	user := flag.String("user", "", "User whose tweets you want to scan")

	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Printf("Scanning %d of %s's tweets...\n", *limit, *user)

	client := twitter.NewClient(os.Getenv("TWITTER_BEARER_TOKEN"))

	tweets, err := client.GetTweetsForUsername(*user, *limit)
	if err != nil {
		log.Fatal(err)
	}

	afinn, err := afinn.NewAfinn()
	if err != nil {
		log.Fatal(err)
	}

	analyser := sentiment.NewAnalyser(afinn)
	var sentiments []sentiment.Sentiment
	for _, tw := range tweets {
		sentiments = append(sentiments, analyser.NewSentiment(tw.Text))
	}

	average := analyser.AverageSentiment(sentiments...)
	fmt.Printf("%#v", average)
}
