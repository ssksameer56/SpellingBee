package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

//https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func CleanWords() {
	x, _ := os.Getwd()
	fmt.Println(x)
	path := "../data/raw_words.txt"
	array, err := readLines(path)
	if err != nil {
		log.Err(err).Msg("dwdsd")
	}
	newArray := []string{}
	for _, word := range array {
		if len(word) >= 4 {
			newArray = append(newArray, word)
		}
	}
	writeLines(newArray, "../data/words.txt")
}
