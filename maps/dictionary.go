package maps

import "errors"

type Dictionary map[string]string

var ErrWordDoesntExist = errors.New("word doesn't exist")

func (d Dictionary) Search(key string) (string, error) {
	// ok will be a boolean , true if key is found false if isn't
	definition, ok := d[key]
	if !ok {
		return "", ErrWordDoesntExist
	}
	return definition, nil
}

func (d Dictionary) Add(word, wordDefintion string) {
	d[word] = wordDefintion
}
