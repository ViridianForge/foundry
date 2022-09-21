package maps

type Dictionary map[string]string

func (d Dictionary) Search(search string) string {
	return d[search]
}
