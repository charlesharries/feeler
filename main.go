package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/charlesharries/feeler/sentiment"
)

func main() {
	router := http.NewServeMux()

	// Get the file here
	file, err := os.Open("data/AFINN-111.txt")
	if err != nil {
		// Dont call log.Fatal because it'll crash the whole program
		// don't call it except for in main()
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	afinn := parseAfinn(reader)

	router.HandleFunc("/", home(afinn))

	http.ListenAndServe("localhost:3001", router)
}

func parseAfinn(file *bufio.Reader) map[string]int {
	m := make(map[string]int)

	for {
		line, err := file.ReadString('\n')
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

func home(afinn map[string]int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			homeGet(w, r)
		case "POST":
			homePost(w, r, afinn)
		default:
			io.WriteString(w, "{ \"status\": \"method not allowed\" }")
		}
	}
}

func homeGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonObj := map[string]string{
		"status": "works divine",
	}

	out, err := json.Marshal(jsonObj)
	if err != nil {
		http.Error(w, "Couldn't marshal json", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func homePost(w http.ResponseWriter, r *http.Request, afinn map[string]int) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Not right content type bud", 400)
	}

	req := struct {
		Sentence *string `json:"s"`
	}{}

	json.NewDecoder(r.Body).Decode(&req)

	sent := sentiment.NewSentiment(*req.Sentence, afinn)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sent)
}
