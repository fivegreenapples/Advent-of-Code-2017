package main

import "fmt"
import "strings"
import "sort"

func main() {

	for _, phrase := range []string{
		"aa bb cc dd ee",
		"aa bb cc dd aa",
		"aa bb cc dd aaa",
	} {
		fmt.Println("Test:", phrase, isPhraseValid(phrase))
	}
	validCount := 0
	for _, phrase := range input {
		if isPhraseValid(phrase) {
			validCount++
		}
	}
	fmt.Println("Answer1: ", validCount)

	for _, phrase := range []string{
		"abcde fghij",
		"abcde xyz ecdab",
		"a ab abc abd abf abj",
		"iiii oiii ooii oooi oooo",
		"oiii ioii iioi iiio",
	} {
		fmt.Println("Test2:", phrase, isPhraseValid2(phrase))
	}
	validCount2 := 0
	for _, phrase := range input {
		if isPhraseValid2(phrase) {
			validCount2++
		}
	}
	fmt.Println("Answer2: ", validCount2)
}

func isPhraseValid(passPhrase string) bool {
	words := strings.Split(passPhrase, " ")

	wordMap := map[string]struct{}{}

	for _, w := range words {
		if _, found := wordMap[w]; found {
			return false
		}
		wordMap[w] = struct{}{}
	}
	return true
}
func isPhraseValid2(passPhrase string) bool {
	words := strings.Split(passPhrase, " ")

	wordMap := map[string]struct{}{}

	for _, w := range words {
		canonicalised := []byte(w)
		sort.Slice(canonicalised, func(i, j int) bool {
			return canonicalised[i] < canonicalised[j]
		})
		if _, found := wordMap[string(canonicalised)]; found {
			return false
		}
		wordMap[string(canonicalised)] = struct{}{}
	}
	return true
}
