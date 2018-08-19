package main

import "fmt"

// Count is custom integer type
type Count int

// Person is a custom structure type
type Person struct {
	Name      string   "name of the person" //tags
	ForeNames []string "fore-father names"
	SurName   string   "sur-name"
}

// Author is custom structure type that embed's the Person structure type
type Author struct {
	Person
	BookTitles []string
	YearBorn   int
}

// EmbedStructs embeds structs
func EmbedStructs() {
	author := Author{
		Person{"robert", []string{"louis", "bob", "max"}, "martin"},
		[]string{"GOT", "GOT2", "GOT3", "GOT4"},
		1950,
	}
	fmt.Println("author.Person \n", author.Person)
	// author.Person
	// {robert [louis bob max] martin}
	fmt.Println("author.Name, author.ForeNames, author.SurName \n", author.Name, author.ForeNames, author.SurName)
	// author.Name, author.ForeNames, author.SurName
	// robert [louis bob max] martin
	fmt.Println("author.BookTitles, author.YearBorn \n", author.BookTitles, author.YearBorn)
	// author.BookTitles, author.YearBorn
	//[GOT GOT2 GOT3 GOT4] 1950
}
