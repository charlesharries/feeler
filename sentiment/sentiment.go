package sentiment

import "strings"

// Sentiment : A breakdown of a sentence's sentiment
type Sentiment struct {
	Verdict     string   `json:"verdict"`
	Score       int      `json:"score"`
	Comparative float64  `json:"comparative"`
	Positive    []string `json:"positiveWords"`
	Negative    []string `json:"negativeWords"`
}

func NewSentiment(str string, afinn map[string]int) Sentiment {
	words := strings.Fields(str)

	sentimentScore := 0
	positiveWords := []string{}
	negativeWords := []string{}

	for _, word := range words {
		if val, ok := afinn[word]; ok {
			sentimentScore += val

			if val > 0 {
				positiveWords = append(positiveWords, word)
			} else if val < 0 {
				negativeWords = append(negativeWords, word)
			}
		}
	}

	comparative := float64(sentimentScore) / float64(len(words))

	verdict := "POSITIVE"
	if sentimentScore < 0 {
		verdict = "NEGATIVE"
	} else if sentimentScore == 0 {
		verdict = "NEUTRAL"
	}

	sentiment := Sentiment{Verdict: verdict, Score: sentimentScore, Comparative: comparative, Positive: positiveWords, Negative: negativeWords}

	return sentiment
}
