package maps

const (
	ErrNotFound         = DictionaryError("that search term could not be found")
	ErrTermExists       = DictionaryError("term already defined")
	ErrWordDoesNotExist = DictionaryError("terms not in dictionary cannot be updated")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Add(term string, definition string) error {
	_, err := d.Search(term)
	switch err {
	case nil:
		return ErrTermExists
	case ErrNotFound:
		d[term] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Search(term string) (string, error) {
	if d[term] == "" {
		return "", ErrNotFound
	}
	return d[term], nil
}

func (d Dictionary) Update(term string, definition string) error {
	_, err := d.Search(term)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[term] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(term string) {
	delete(d, term)
}
