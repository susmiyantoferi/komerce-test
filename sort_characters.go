package main

import (
	"fmt"
	"strings"
)

func SortChar(s string) {
	sToLower := strings.ToLower(s)
	text := strings.ReplaceAll(sToLower, " ", "")

	var vowels, conso []string

	for _, v := range text {
		val := string(v)
		isVowels := val == "a" || val == "i" || val == "u" || val == "e" || val == "o"
		if isVowels {
			vowels = append(vowels, val)
		} else if val >= "a" && val <= "z" {
			conso = append(conso, val)
		}
	}

	fmt.Println("Vowels:", strings.Join(vowels, ""))
	fmt.Println("Consonants:", strings.Join(conso, ""))
}

func main() {

	case1 := "Sample Case"
	case2 := "Next Case"

	SortChar(case1)
	SortChar(case2)

}
