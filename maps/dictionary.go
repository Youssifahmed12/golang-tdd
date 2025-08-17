package maps

type Dictionary map[string]string

func (d Dictionary) Search(dic map[string]string, key string) string {
	return dic[key]
}
