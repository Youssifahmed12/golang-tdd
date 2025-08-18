package maps

const (
	ErrWordDoesntExist = DictionaryError("word doesn't exist")
	ErrWordExist       = DictionaryError("word exist")
)

type Dictionary map[string]string
type DictionaryError string

func (dr DictionaryError) Error() string {
	return string(dr)
}

func (d Dictionary) Search(key string) (string, error) {
	// ok will be a boolean , true if key is found false if isn't
	definition, ok := d[key]
	if !ok {
		return "", ErrWordDoesntExist
	}
	return definition, nil
}

func (d Dictionary) Add(word, wordDefintion string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordDoesntExist:
		d[word] = wordDefintion
	case nil:
		return ErrWordExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordDoesntExist:
		return ErrWordDoesntExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}
