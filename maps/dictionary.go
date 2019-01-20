package maps

//Search function searches the dictionary for a specific key, This was my original implementation of the function
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

// Textbook Implementation of the function
type Dictionary map[string]string

const (
	ErrorNotFound   = DictionaryErr("Couldn't find the word")
	ErrorWordExists = DictionaryErr("Error word exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) SearchV2(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

//Add function adds a word and definition to an existing Dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.SearchV2(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil

}
