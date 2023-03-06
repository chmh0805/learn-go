package mydict

import "errors"

// Dictionary
type Dictionary map[string]string // type alias

var errNotFound = errors.New("not Found");
var errAlreadyExist = errors.New("That word already Exists");

// Search a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]; // string, boolean
	if (exists) {
		return value, nil;
	}
	return "", errNotFound;
}

// Add a word to dictionary
func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word);
	switch err {
	case errNotFound:
		d[word] = definition;
	case nil:
		return errAlreadyExist;
	}
	return nil;
}