package main

import (
	"fmt"
	"learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{
		"first": "First Word",
	}
	baseWord := "hello";
	dictionary.Add(baseWord, "First");
	
	err := dictionary.Update(baseWord, "Second");
	if (err != nil) {
		fmt.Println(err);
	}

	word, _ := dictionary.Search(baseWord);
	fmt.Println(word);

	err2 := dictionary.Delete(baseWord);
	if (err2 != nil) {
		fmt.Println(err2);
	}
	
	word, err3 := dictionary.Search(baseWord);
	if err3 != nil {
		fmt.Println(err3);
	} else {
		fmt.Println(word);
	}
}