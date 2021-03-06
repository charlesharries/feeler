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
	Length      int      `json:"length"`
	Score       int      `json:"score"`
	Comparative float64  `json:"comparative"`
	Positive    []string `json:"positive_words"`
	Negative    []string `json:"negative_words"`
}

// stripPunctuation removes any punctuation that could mess up our word lookup.
func stripPunctuation(str string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")

	return reg.ReplaceAllString(str, "")
}

// NewAnalyser returns a new analyser for running sentiment analysis against.
func NewAnalyser(a map[string]int) *Analyser {
	return &Analyser{dictionary: a}
}

// valueOf gets the value of a word in the received Analyser's dictionary, or
// 0 if it doesn't exist in the dictionary.
func (a *Analyser) valueOf(word string) int {
	if val, ok := a.dictionary[word]; ok {
		return val
	}

	return 0
}

// isNegator determines if a certain word appears in the list of negators.
func (a *Analyser) isNegator(word string) bool {
	for _, negator := range negators {
		if strings.ToLower(word) == negator {
			return true
		}
	}

	return false
}

// NewSentiment generates a Sentiment result for a given sentence.
func (a *Analyser) NewSentiment(str string) Sentiment {
	words := strings.Fields(str)

	sentimentScore := 0
	positive_words := []string{}
	negative_words := []string{}

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
			positive_words = append(positive_words, word)
		} else if val < 0 {
			negative_words = append(negative_words, word)
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
		Length:      len(words),
		Score:       sentimentScore,
		Comparative: comparative,
		Positive:    positive_words,
		Negative:    negative_words,
	}
}

// AverageSentiment generates a single average Sentiment from an
// array of Sentiments.
func (a *Analyser) AverageSentiment(sentiments ...Sentiment) Sentiment {
	score := 0
	words := 0
	positive := []string{}
	negative := []string{}

	for _, s := range sentiments {
		score += s.Score
		words += s.Length

		positive = append(positive, s.Positive...)
		negative = append(negative, s.Negative...)
	}

	verdict := "NEUTRAL"
	if score > 0 {
		verdict = "POSITIVE"
	} else if score < 0 {
		verdict = "NEGATIVE"
	}

	return Sentiment{
		Verdict:     verdict,
		Length:      words,
		Score:       score,
		Comparative: float64(score) / float64(words),
		Positive:    positive,
		Negative:    negative,
	}
}
