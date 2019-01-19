package maps

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
