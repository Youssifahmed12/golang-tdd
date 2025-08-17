package maps

import "errors"

type Dictionary map[string]string

var ErrWordDoesntExist = errors.New("word doesn't exist")

func (d Dictionary) Search(dic map[string]string, key string) (string, error) {
	// ok will be a boolean , true if key is found false if isn't
	definition, ok := dic[key]
	if !ok {
		return "", ErrWordDoesntExist
	}
	return definition, nil
}
