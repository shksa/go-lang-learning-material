package main

import (
	"fmt"
	"strings"
	"unicode"
)

type StringPair struct {
	first, second string
}

type LowerCaser interface {
	LowerCase()
}

type UpperCaser interface {
	UpperCase()
}

type LowerUpperCaser interface {
	LowerCaser // As if we had written LowerCase()
	UpperCaser // As if we had written UpperCase()
}

type FixCaser interface {
	FixCase()
}

type ChangeCaser interface {
	LowerUpperCaser // As if we had written LowerCase(); UpperCase()
	FixCaser        // As if we had written FixCase()
}

func (pair *StringPair) UpperCase() {
	pair.first = strings.ToUpper(pair.first)
	pair.second = strings.ToUpper(pair.second)
}

func (pair *StringPair) LowerCase() {
	pair.first = strings.ToLower(pair.first)
	pair.second = strings.ToLower(pair.second)
}

func (pair *StringPair) FixCase() {
	pair.first = fixCase(pair.first)
	pair.second = fixCase(pair.second)
}

func fixCase(s string) string {
	var chars []rune
	upper := true
	for _, char := range s {
		if upper {
			char = unicode.ToUpper(char)
		} else {
			char = unicode.ToLower(char)
		}
		chars = append(chars, char)
		upper = unicode.IsSpace(char) || unicode.Is(unicode.Hyphen, char)
	}
	return string(chars)
}

func EmbeddingInterfaces() {
	lobelia := StringPair{"LOBELIA", "SACKVILLE-BAGGINS"}
	// Values of *StringPair belongs to all of these custom interface types
	// 1. LowerCaser
	// 2. UpperCaser
	// 3. LowerUpperCaser
	// 4. FixCaser
	// 5. ChangeCaser
	//
	lobelia.FixCase()
	fmt.Printf("%#v \n", lobelia)
	// main.StringPair{first:"Lobelia", second:"Sackville-Baggins"}
}
