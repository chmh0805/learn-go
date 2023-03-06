package main

import (
	"fmt"
	"learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{
		"first": "First Word",
	}

	word := "hello";
	definition := "Greeting";
	err := dictionary.Add(word, definition);
	if err != nil {
		fmt.Println(err);
	}

	def, err2 := dictionary.Search(word);
	if err2 != nil {
		fmt.Println(err2);
	} else {
		fmt.Println(def);
	}

	err3 := dictionary.Add(word, definition);
	if err3 != nil {
		fmt.Println(err3);
	}
}