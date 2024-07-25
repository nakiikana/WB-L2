package main

import (
	"sort"
	"strings"
)

func sortString(s string) string {
    chars := strings.Split(s, "")
    sort.Strings(chars)
    return strings.Join(chars, "")
}

func findAnagramSets(words []string) map[string][]string {
    anagramSets := make(map[string][]string)

    for _, word := range words {
        sortedWord := sortString(strings.ToLower(word))

        anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
    }

    result := make(map[string][]string)
    for _, value := range anagramSets {
        if len(value) > 1 {
            sort.Strings(value)
            result[value[0]] = value
        }
    }

    return result
}

func main() {
    words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
    anagramSets := findAnagramSets(words)
 
    for key, value := range anagramSets {
        println(key, ":", strings.Join(value, ", "))
    }
}
