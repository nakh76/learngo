package mydict

import "errors"

type Dictionary map[string]string

var errNotFount = errors.New("not found")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}

	return "", errNotFount
}
