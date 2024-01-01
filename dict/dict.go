package dict

type Dictionary map[string]string
type DictError string

func (d DictError) Error() string {
	return string(d)
}

const (
	DictKeyNotFoundErr    = DictError("could not find the word you were looking for")
	DictKeyCollisionError = DictError("could not add word because it already exists")
)

func (d Dictionary) Search(targetKey string) (string, error) {
	value, doesExist := d[targetKey]

	if !doesExist {
		return "", DictKeyNotFoundErr
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case DictKeyNotFoundErr:
		d[key] = value
		return nil
	case nil:
		return DictKeyCollisionError
	default:
		return err
	}
}

func (d Dictionary) Update(key, value string) (string, error) {
	existingVal, err := d.Search(key)

	switch err {
	case DictKeyNotFoundErr:
		return "", DictKeyNotFoundErr
	case nil:
		d[key] = value
		return existingVal, nil
	default:
		return "", nil
	}
}

func (d Dictionary) Delete(key string) string {
	existingVal, _ := d.Search(key)
	delete(d, key)

	return existingVal
}
