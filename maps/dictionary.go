package maps

import "errors"

func Search(dictionary map[string]string, word string) string {
	for k, v := range dictionary {
		if k == word {
			return v
		} else {
			return "Couldn't find the word"
		}
	}

	return dictionary[word]
}

type Dictionary map[string]string

var ErrorNotFound = errors.New("Couldn't find the word")

func (d Dictionary) SearchV2(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, def string) {
	d[word] = def
}
