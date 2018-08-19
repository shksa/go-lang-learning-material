package main

import (
	"fmt"
)

func crateCustomTypes() {
	type Count int
	type StringMap map[string]string
	type IntList []int
	type StringForString func(string) string // A custom type that specifies a function signature

	// variable of custom type "Count"
	// ci := Count{5} // invalid syntax
	var i Count = 7
	fmt.Printf("i: Type: %T and Value: %v \n", i, i)

	// variable of custom type "StringMap"
	var ownerOf = StringMap{"rocky": "johncena"}
	fmt.Printf("ownerOf: Type: %T and Value: %#v \n", ownerOf, ownerOf)

	// variable of custom type "IntList"
	var scores = IntList{4, 7, 3, 2}
	fmt.Printf("scores: Type: %T and Value: %v \n", scores, scores)

	// variable of custom type "StringForString"
	var removePunctuation = func(testString string) string {
		newString := ""
		for _, runeCharacter := range testString {
			switch character := string(runeCharacter); character {
			case "a", "e", "i", "o", "u":
			default:
				newString += string(character)
			}
		}
		return newString
	}

	var processStrings = func(removePunctuation StringForString, strings []string) []string {
		result := []string{}
		for _, testString := range strings {
			result = append(result, removePunctuation(testString))
		}
		return result
	}

	var stringsWithoutVowels = processStrings(removePunctuation, []string{
		"apple",
		"peanut",
		"domestic",
	})

	fmt.Printf("%v \n", stringsWithoutVowels)
}

func main() {
	crateCustomTypes()
}
