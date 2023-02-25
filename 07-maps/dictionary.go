// @Title Dictionary
// @Description learn maps for go
// @Author Belaoura Yacoub
// @Update v0.0.1
package maps

type Dictionary map[string]string

type ErrDictionary string

func (e ErrDictionary) Error() string {
	return string(e)
}

const (
	ErrNotFound          = ErrDictionary("could not find the word you were looking for")
	ErrWordExiste        = ErrDictionary("word already existe")
	ErrWordDoseNotExiste = ErrDictionary("Can't update word dose not existe")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExiste
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = newDefinition
	case ErrNotFound:
		return ErrWordDoseNotExiste
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
