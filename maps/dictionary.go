package maps

const (
	ErrNotFound         = DictionaryErr("word doesn't exist")
	ErrWordExist        = DictionaryErr("word exist")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type Dictionary map[string]string
type DictionaryErr string

func (dr DictionaryErr) Error() string {
	return string(dr)
}

func (d Dictionary) Search(key string) (string, error) {
	// ok will be a boolean , true if key is found false if isn't
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, wordDefintion string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
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
	case ErrNotFound:
		return ErrNotFound
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}
