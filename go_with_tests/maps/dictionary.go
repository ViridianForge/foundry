package maps

import (
	"errors"
)

var ErrNotFound = errors.New("that search term could not be found")

type Dictionary map[string]string

func (d Dictionary) Add(term string, definition string) {
	d[term] = definition
}

func (d Dictionary) Search(term string) (string, error) {
	if d[term] == "" {
		return "", ErrNotFound
	}
	return d[term], nil
}
