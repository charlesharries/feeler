package sentiment

import (
	"fmt"
	"regexp"
	"strings"
)

// negators are words that reverse the feeling of the following word.
var negators = []string{
	"cant",
	"can't",
	"dont",
	"don't",
	"doesnt",
	"doesn't",
	"not",
	"non",
	"wont",
	"won't",
	"isnt",
	"isn't",
}

// Analyser is a struct for running analyses. It depends on a map of
// words to integer values but otherwise should be pretty adaptable.
type Analyser struct {
	dictionary map[string]int
}

// Sentiment is a data breakdown of a string's sentiment.
type Sentiment struct {
	Verdict     string   `json:"verdict"`
	Score       int      `json:"score"`
	Comparative float64  `json:"comparative"`
	Positive    []string `json:"positiveWords"`
	Negative    []string `json:"negativeWords"`
}

// stripPunctuation removes any punctuation that could mess up our word lookup.
func stripPunctuation(str string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")

	return reg.ReplaceAllString(str, "")
}

// NewAnalyser returns a new analyser for running sentiment analysis against.
func NewAnalyser(a map[string]int) Analyser {
	return Analyser{dictionary: a}
}

// valueOf gets the value of a word in the received Analyser's dictionary, or
// 0 if it doesn't exist in the dictionary.
func (a Analyser) valueOf(word string) int {
	if val, ok := a.dictionary[word]; ok {
		return val
	}

	return 0
}

// isNegator determines if a certain word appears in the list of negators.
func (a Analyser) isNegator(word string) bool {
	for _, negator := range negators {
		if strings.ToLower(word) == negator {
			return true
		}
	}

	return false
}

// NewSentiment generates a Sentiment result for a given sentence.
func (a Analyser) NewSentiment(str string) Sentiment {
	words := strings.Fields(str)

	sentimentScore := 0
	positiveWords := []string{}
	negativeWords := []string{}

	// Iterate over each word in the sentence
	for i, word := range words {
		word = stripPunctuation(word)
		val := a.valueOf(word)

		if i > 0 && a.isNegator(words[i-1]) {
			val = val * -1
			word = fmt.Sprintf("%s (negated)", word)
		}

		sentimentScore += val

		if val > 0 {
			positiveWords = append(positiveWords, word)
		} else if val < 0 {
			negativeWords = append(negativeWords, word)
		}
	}

	comparative := float64(sentimentScore) / float64(len(words))
	if len(words) == 0 {
		comparative = 0
	}

	verdict := "POSITIVE"
	if sentimentScore < 0 {
		verdict = "NEGATIVE"
	} else if sentimentScore == 0 {
		verdict = "NEUTRAL"
	}

	return Sentiment{
		Verdict:     verdict,
		Score:       sentimentScore,
		Comparative: comparative,
		Positive:    positiveWords,
		Negative:    negativeWords,
	}
}
