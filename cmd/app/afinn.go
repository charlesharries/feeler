package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func newAfinn() (map[string]int, error) {
	a := map[string]int{}

	// Open the file.
	file, err := os.Open("data/AFINN-165.txt")
	if err != nil {
		return a, err
	}
	defer file.Close()

	// Create a scanner, which will read the file line by line. For each
	// line, add the word to the map, with its corresponding value.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keyValue := strings.Split(scanner.Text(), "\t")
		value, err := strconv.Atoi(strings.TrimSpace(keyValue[1]))
		if err != nil {
			return a, err
		}

		a[keyValue[0]] = value
	}

	return a, nil
}
