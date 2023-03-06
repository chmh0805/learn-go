package mydict

import "errors"

// Dictionary
type Dictionary map[string]string // type alias

var errNotFound = errors.New("not Found");

// Search a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]; // string, boolean
	if (exists) {
		return value, nil;
	}
	return "", errNotFound;
}