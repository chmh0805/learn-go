package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string // type alias

var (
	errNotFound		= errors.New("not Found");
	errAlreadyExist = errors.New("that word already Exists");
	errCantUpdate	= errors.New("can't update non-existing word");
	errCantDelete	= errors.New("can't delete non-existing word");
);

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

// Update a word in dictionary
func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word);
	switch err {
	case errNotFound:
		return errCantUpdate;
	case nil:
		d[word] = definition;
	}
	return nil;
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word);
	switch err {
	case errNotFound:
		return errCantDelete;
	case nil:
		delete(d, word);
	}
	return nil;
}