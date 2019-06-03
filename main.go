package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Sentence : A sentence to have the sentiment analysed
type Sentence struct {
	s string
}

// Sentiment : A breakdown of a sentence's sentiment
type Sentiment struct {
	Verdict     string   `json:"verdict"`
	Score       int      `json:"score"`
	Comparative float64  `json:"comparative"`
	Positive    []string `json:"positiveWords"`
	Negative    []string `json:"negativeWords"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/", checkSentiment).Methods("POST")

	http.ListenAndServe("localhost:3001", router)
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "{ \"status\": \"works divine\" }")
}

func checkSentiment(res http.ResponseWriter, req *http.Request) {
	var body map[string]string
	json.NewDecoder(req.Body).Decode(&body)

	sentence := body["s"]
	sentiment := getSentiment(sentence)

	fmt.Println(sentiment)

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(sentiment)
}

func getSentiment(str string) Sentiment {
	words := strings.Fields(str)
	m := getMap("data/AFINN-111.txt")

	sentimentScore := 0
	var positiveWords []string
	var negativeWords []string

	for _, word := range words {
		if val, ok := m[word]; ok {
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

func getMap(filename string) map[string]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	m := make(map[string]int)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		keyValue := strings.Split(line, "\t")
		value, err := strconv.Atoi(strings.TrimSpace(keyValue[1]))
		if err != nil {
			log.Fatal(err)
		}

		m[keyValue[0]] = value
	}

	return m
}
